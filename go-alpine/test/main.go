package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	fmt.Printf("运行时间: %s\n", time.Now())
	fmt.Printf("运行时区: %s\n", time.Local.String())

	resp, err := http.Get("https://go.dev/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	if len(all) < 100 {
		fmt.Println("resp body < 100")
		return
	}
	fmt.Println(string(all)[0:500])
}
