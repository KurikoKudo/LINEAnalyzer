package lineanalyzer

import (
	"testing"
	"time"
)

func TestMassageFormatter (t *testing.T) {


	date := "2017/10/30"
	buf := "06:42	あつし	そんなに重くないので5限終わりまでになんとかお願いします"

	const layout = "2006/01/02 15:04"
	collectSendTime,_ := time.Parse(layout,"2017/10/30 06:42")


	massage := MassageFormatter(buf, date)

	if massage.Sender != "あつし" {
		t.Fatalf("Senderが違う:%s",massage.Sender)
	}
	if massage.SendTime != collectSendTime{
		t.Fatalf("SendTimeが違う:%s",massage.SendTime)
	}
	if massage.Comment != "そんなに重くないので5限終わりまでになんとかお願いします" {
		t.Fatalf("Commentが違う:%s",massage.Comment)
	}

}