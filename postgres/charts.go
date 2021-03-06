package postgres

import "github.com/raedahgroup/dcrextdata/cache"

func (pg *PgDb) RegisterCharts(charts *cache.ChartData, syncSources []string, syncSourceDbProvider func(source string) (*PgDb, error)) {
	pg.syncSourceDbProvider = syncSourceDbProvider
	pg.syncSources = syncSources

	charts.AddUpdater(cache.ChartUpdater{
		Tag:      "mempool chart",
		Fetcher:  pg.retrieveChartMempool,
		Appender: appendChartMempool,
	})

	charts.AddUpdater(cache.ChartUpdater{
		Tag:      "block propagation chart",
		Fetcher:  pg.fetchBlockPropagationChart,
		Appender: appendBlockPropagationChart,
	})

	charts.AddUpdater(cache.ChartUpdater{
		Tag:      "PoW chart",
		Fetcher:  pg.fetchPowChart,
		Appender: appendPowChart,
	})

	charts.AddUpdater(cache.ChartUpdater{
		Tag:      "VSP chart",
		Fetcher:  pg.fetchVspChart,
		Appender: appendVspChart,
	})

	charts.AddUpdater(cache.ChartUpdater{
		Tag:      "Exchange chart",
		Fetcher:  pg.fetchExchangeChart,
		Appender: appendExchangeChart,
	})
}
