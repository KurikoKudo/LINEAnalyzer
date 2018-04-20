package lineanalyzer

import (
	"strings"
	"time"
)

func MassageFormatter (buf string,date string) Massagae {

	var massage Massagae

	splitStr := strings.Split(buf,"	")
	const layout = "2018/01/25 11:38"

	date = date+" "+splitStr[0]
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