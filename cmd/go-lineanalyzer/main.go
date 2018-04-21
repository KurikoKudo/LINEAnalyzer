package main

import (
	"os"
	"github.com/KurikoKudo/go-lineanalyzer"
	"fmt"
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


}