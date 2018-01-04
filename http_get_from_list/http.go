package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// 1ホストあたりの最大コネクション数を設定
func setMaxConnection(n int) {
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = n
}

func get(url string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	//req.Header.Add("Connection", "Keep-Alive")	// 要らないみたい
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	// 1. レスポンスを最後まで読み込む
	// 2. レスポンスをクローズする
	// をしないとKeep-Aliveが有効にならない
	defer res.Body.Close()
	//io.Copy(ioutil.Discard, res.Body)

	// ステータスコード
	fmt.Printf("status = %d\n", res.StatusCode)
	// レスポンスボディ
	bodyBytes, err := ioutil.ReadAll(res.Body)
	fmt.Print(string(bodyBytes))
	// レスポンスヘッダ
	for k, values := range res.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", k, value)
		}
	}
}
