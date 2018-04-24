package lineanalyzer

import (
	"strings"
	"time"
)

//MassageFormatter は読み込まれたメッセージをフォーマットする関数
func MassageFormatter (buf string,date string) Massage {

	var massage Massage

	splitStr := strings.Split(buf,"	")
	const layout = "2006/01/02 15:04"

	date = date+" "+splitStr[0]
	//fmt.Println("date",date)
	massage.SendTime,_ = time.Parse(layout,date)

	if len(splitStr) == 3 {
		massage.Sender = splitStr[1]
		massage.Comment = splitStr[2]
	} else {
		massage.Sender = "システムによるメッセージ"
		massage.Comment = splitStr[1]
	}

	return massage

}