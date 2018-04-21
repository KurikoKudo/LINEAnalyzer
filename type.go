package lineanalyzer

import "time"

type LineAllTalk struct {
	Title string
	SaveDate time.Time
	Massages []Massage
}

type Massage struct {
	SendTime time.Time
	Sender string
	Comment string
}