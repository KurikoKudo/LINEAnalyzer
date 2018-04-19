package lineanalyzer

import "time"

type LineAllTalk struct {
	Title string
	SaveDate time.Time
	Massagaes []Massagae
}

type Massagae struct {
	SendTime time.Time
	Sender string
	Comment string
}