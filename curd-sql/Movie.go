package main

type Movies struct {
	MovieId int64 `json: "movieid"`

	MovieName string `json: "moviename"`

	MovieCost float64 `json: "moviecost"`
}
