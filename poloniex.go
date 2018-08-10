package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

const (
	poloniexBaseURL = viper.Get("ExchangeData[0]")
)

//Poloniex Structure containing Poloniex client data
type Poloniex struct {
	client *http.Client
}

//Structure containing Poloniex Historic Data

type poloniexData struct {
	Result []struct {
		GlobalTradeID int64   `json:"globalTradeID"`
		TradeID       string  `json:"tradeID"`
		Date          string  `json:"date"`
		Types         string  `json:"type"`
		Rate          float64 `json:"rate"`
		Amount        float64 `json:"amount"`
		Total         float64 `json:"total"`
	}
}

// Structure containing Poloniex Chart Data

type chartData struct {
	Result []struct {
		Date            int64   `json:"date"`
		High            float64 `json:"high"`
		Low             float64 `json:"low"`
		Open            float64 `json:"open"`
		Close           float64 `json:"close"`
		Volume          float64 `json:"volume"`
		QuoteVolume     float64 `json:"quoteVolume"`
		WeightedAverage float64 `json:"weightedAverage"`
	}
}

// Get Poloniex Historic Data
// parameters : Currency pair, Start time , End time

func (p *Poloniex) getPoloniexData(currencyPair string, start string, end string) string {

	dbInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		viper.Get("Database.pghost"), 5432, viper.Get("Database.pguser"), viper.Get("Database.pgpass"),
		viper.Get("Database.pgdbname"))

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)

	}

	// boil.SetDB(db)

	//Get Url of Poloniex API

	url := poloniexBaseURL

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err.Error())
	}

	//Append user provided parameters in the URL

	q := req.URL.Query()
	q.Add("command", "returnTradeHistory")
	q.Add("currencyPair", currencyPair)
	q.Add("start", start)
	q.Add("end", end)
	req.URL.RawQuery = q.Encode()

	//Get Historic Data from the API

	request, err := http.NewRequest("GET", req.URL.String(), nil)

	res, _ := p.client.Do(request)

	//Get response of the request as []byte

	fmt.Println(res.StatusCode)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	//Map the data to poloniexData struct and unmarshal the contents

	var data poloniexData
	json.Unmarshal(body, &data)

	fmt.Printf("Results: %v\n", data)

	for i := range data.Result {

		err := db.QueryRow("Insert into historic_data Values $1,$2,$3,$4,$5,$6,$7,$8,$9",
			0, data.Result[i].GlobalTradeID, data.Result[i].TradeID, data.Result[i].Date, data.Result[i].Amount, data.Result[i].Rate,
			data.Result[i].Total, "nil", data.Result[i].Types)

	}

	return "Saved poloneix historic data!"
}

//Returns data from Poloniex exchange
//Parameters : currency pair , start time , end time

func (p *Poloniex) fetchPoloniexData(date string) {

	dbInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		viper.Get("Database.pghost"), 5432, viper.Get("Database.pguser"), viper.Get("Database.pgpass"),
		viper.Get("Database.pgdbname"))

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)

	}

	rows, err := db.Query("SELECT * FROM historic_data where timestamp=$1", date)

}

//Returns Poloniex Chart Data
//Parameters : Currency pair, Start time , End time

func (p *Poloniex) getChartData(currencyPair string, start string, end string) {

	//Get the base URL

	url := poloniexBaseURL

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err.Error())
	}

	//Append the user defined parameters to the url

	q := req.URL.Query()
	q.Add("command", "returnChartData")
	q.Add("currencyPair", currencyPair)
	q.Add("start", start)
	q.Add("end", end)
	req.URL.RawQuery = q.Encode()

	request, err := http.NewRequest("GET", req.URL.String(), nil)

	//Get the data from API and convert the data to byte array

	res, _ := p.client.Do(request)

	fmt.Println(res.StatusCode)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	dbInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		viper.Get("Database.pghost"), 5432, viper.Get("Database.pguser"), viper.Get("Database.pgpass"),
		viper.Get("Database.pgdbname"))

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)

	}
	//Store the data to charData struct

	var data chartData
	json.Unmarshal(body, &data)
	fmt.Printf("Results: %v\n", data)

	//Loop over the entire data and store it in the table
	for i := range data.Result {
		sqlStatement := `INSERT INTO chart_data(exhcangeID,date,high,low,open,close,volume,quoteVolume,baseVolume,weightedAverage) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`
		_, err = db.Exec(sqlStatement, "0", data.Result[i].Date, data.Result[i].High, data.Result[i].Low, data.Result[i].Open, data.Result[i].Close, data.Result[i].Volume, data.Result[i].QuoteVolume, "nil", data.Result[i].WeightedAverage)

	}

}
