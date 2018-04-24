package lineanalyzer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"log"
)

//FileLoader はファイルを読み込む関数
func FileLoader(path string) (LineAllTalk, error) {

	var err error
	var lineAllTalk LineAllTalk
	var date string


	kuriko := 0
	madoka := 0
	atsushi := 0
	masaya := 0
	system := 0


	fp, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return lineAllTalk,err
	}


	scanner := bufio.NewScanner(fp)



	for i:=0;scanner.Scan();i++ {

		if i == 0 {
			lineAllTalk.Title = scanner.Text()
		} else if i == 1 {
			const saveDateFmt= "2006/01/02 15:04"
			buf := scanner.Text()
			line := strings.Split(buf, "保存日時：")
			str := strings.Join(line, "")
			lineAllTalk.SaveDate, _ = time.Parse(saveDateFmt, str)
		} else {

			buf := scanner.Text()
			//fmt.Println(buf)
			str := strings.Split(buf, "")

			//fmt.Println("str",str)

			if len(str) != 0 {
				if len(str)>2&&str[2] == ":" {
					// 一般的なメッセージ
					massage := MassageFormatter(buf, date)
					lineAllTalk.Massages = append(lineAllTalk.Massages, massage)

					if massage.Sender == "くどうくりこ＊" {
						kuriko++
					}
					if massage.Sender == "あつし" {
						atsushi++
					}
					if massage.Sender == "まどか" {
						madoka++
					}
					if massage.Sender == "田中 将哉" {
						masaya++
					}
					if massage.Sender == "システムによるメッセージ" {
						system++
					}

				} else if len(str)>4&&str[4] == "/" {
					// 日付
					line := strings.Split(buf, "(")
					date = line[0]

				} else {
					// 二行以上のメッセージ
					lineAllTalk.Massages[len(lineAllTalk.Massages)-1].Comment += buf
				}

			}

		}



	}
	if err := scanner.Err(); err != nil {
		return lineAllTalk, err
	}


	/*for _,v := range lineAllTalk.Massages {
		fmt.Print("date : ",v.SendTime)
		fmt.Print(", sender : ",v.Sender)
		fmt.Print(", massage : ",v.Comment,"\n")
	}*/


	fmt.Println("---------------送信者内訳----------------")
	fmt.Println("kuriko:",kuriko)
	fmt.Println("atsushi:",atsushi)
	fmt.Println("madoka:",madoka)
	fmt.Println("masaya:",masaya)
	fmt.Println("システム:",system)

	return lineAllTalk, err

}
