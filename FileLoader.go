package lineanalyzer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func FileLoader(fp *os.File) (LineAllTalk, err) {

	var err error
	var lineAllTalk LineAllTalk

	if len(os.Args) < 2 {
		fp = os.Stdin
	} else {
		fp, err = os.Open(os.Args[1])
		if err != nil {
			return lineAllTalk, err
		}
		defer fp.Close()
	}

	scanner := bufio.NewScanner(fp)

	lineAllTalk.Title = scanner.Text()

	const saveDateFmt = "2018/04/18 09:47"
	buf := scanner.Text()
	str := buf[5:20]
	lineAllTalk.SaveDate,_ = time.Parse(saveDateFmt,str)
	_ = scanner.Text()

	buf = scanner.Text()
	date := buf[0:9]

	fmt.Println("first date",date)


	for scanner.Scan() {

		buf = scanner.Text()
		fmt.Println(buf)
		str := strings.Split(buf, "")

		if str[0] != "¥n" {
			if str[2] == ":" {
				// 一般的なメッセージ
				massage := MassageFormatter(buf,date)
				lineAllTalk.Massagaes = append(lineAllTalk.Massagaes, massage)
			} else if str[4] == "/"{
				// 日付
				date = buf[0:9]

			} else {
				// 二行以上のメッセージ
				lineAllTalk.Massagaes[len(lineAllTalk.Massagaes)-1].Comment += buf
			}

		}

	}
	if err := scanner.Err(); err != nil {
		return lineAllTalk, err
	}

	for _,v := range lineAllTalk.Massagaes {
		fmt.Println(v)
	}

	return lineAllTalk, err

}
