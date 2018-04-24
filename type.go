package lineanalyzer

import "time"

// LineAllTalk は全てのメッセージが格納されるストラクトです
type LineAllTalk struct {
	Title string
	SaveDate time.Time
	Massages []Massage
}

// Massage は送信日付,送信者, メッセージ内容を含むストラクトです
type Massage struct {
	SendTime time.Time
	Sender string
	Comment string
}