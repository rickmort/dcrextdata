// Copyright (c) 2018-2019 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"sync"
	// "context"

	"github.com/raedahgroup/dcrextdata/exchanges"
	"github.com/raedahgroup/dcrextdata/postgres"
	"github.com/raedahgroup/dcrextdata/version"
	// "github.com/raedahgroup/dcrextdata/vsp"
	"github.com/raedahgroup/dcrextdata/web"
)

// const dcrlaunchtime int64 = 1454889600
// var opError error
// var beginShutdown = make(chan bool)

func _main(ctx context.Context) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}

	defer func() {
		if logRotator != nil {
			logRotator.Close()
		}
	}()

	// ctx, _ := context.WithCancel(context.Background())
	enterHttpMode(cfg.HTTPHost, cfg.HTTPPort)

	// Display app version.
	log.Infof("%s version %v (Go version %s)", version.AppName,
		version.Version(), runtime.Version())

	db, err := postgres.NewPgDb(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName)
	defer func(db *postgres.PgDb) {
		err := db.Close()
		if err != nil {
			log.Error("Could not close database connection: %v", err)
		}
	}(db)

	if err != nil {
		return err
	}

	wg := new(sync.WaitGroup)

	if !cfg.DisableVSP {
		vspCollector, err := vsp.NewVspCollector(cfg.VSPInterval, db)
		if err == nil {
			wg.Add(1)
			go vspCollector.Run(ctx, wg)
		} else {
			log.Error(err)
		}
	}

	if !cfg.DisableExchangeTicks {
		ticksHub, err := exchanges.NewTickHub(ctx, cfg.DisabledExchanges, db)
		if err == nil {
			wg.Add(1)
			go ticksHub.Run(ctx, wg)
		} else {
			log.Error(err)
		}
	}

	wg.Wait()

	log.Info("Goodbye")
	return nil
}

func main() {
	// Create a context that is cancelled when a shutdown request is received
	// via requestShutdown.
	ctx := withShutdownCancel(context.Background())
	// Listen for both interrupt signals and shutdown requests.
	go shutdownListener()

	if err := _main(ctx); err != nil {
		if logRotator != nil {
			log.Error(err)
		} else {
			fmt.Println(err)
		}
		os.Exit(1)
	}
	os.Exit(0)
}

func enterHttpMode(host, port string) {
	web.StartHttpServer(host, port)
	// only trigger shutdown if some error occurred, ctx.Err cases would already have triggered shutdown, so ignore
	
}