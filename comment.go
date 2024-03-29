package main

import "time"

type Comment struct {
	Id	int		`json:"id"`
	Lat float64	`json:"lat"`
	Lon float64	`json:"lon"`
	Inside bool `json:"inside"`
	Time time.Time 	`json:"time"`
	Nick string `json:"nick"`
	Text string `json:"text"`
}

func (c *Comment) GetCoordinate() *Coordinate {
	return NewCoordinate(c.Lat, c.Lon)
}

type Comments struct {
	Comments 	*[]Comment `json:"comments"`
}
