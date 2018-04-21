package main

import (
	"os"
	"github.com/KurikoKudo/go-lineanalyzer"
	"fmt"
	"time"
)

func main() {



	lineAllTalk , err := lineanalyzer.FileLoader(`/Users/kudoukuriko/go/src/github.com/KurikoKudo/go-lineanalyzer/sysprotalk.txt`)

	if err != nil{
		fmt.Println("file open err!!")
		os.Exit(1)
	}

	fmt.Println("------------トーク概要-----------")
	fmt.Println("タイトル：",lineAllTalk.Title)
	fmt.Println("保存日時：",lineAllTalk.SaveDate)
	fmt.Println("トーク総数：",len(lineAllTalk.Massages))

	developMassages := 0
	designMassages := 0

	for _,v := range lineAllTalk.Massages {
		if !time.Date(2017, time.December, 11, 0, 0, 0, 0, time.UTC).After(v.SendTime) {
			developMassages++
		} else {
			designMassages++
		}
	}

	fmt.Println("------------トーク数内訳------------")
	fmt.Println("設計ターム（2017/12/10以前）：",designMassages)
	fmt.Println("実装ターム（2017/12/11以降）：",developMassages)

}