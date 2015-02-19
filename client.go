package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

type ServerInfo struct {
        ServerLoad []string
        ServerTime string
	ServerName string
}



func getHTTPContent(url string, website string, c chan ServerInfo) {
	var sInfo ServerInfo
	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &sInfo)
	sInfo.ServerName = website
	c <- sInfo
}




func printResult(sInfo ServerInfo) {
	fmt.Printf("name: %v\ntime: %v\nLoad: %v \n\n", sInfo.ServerName, sInfo.ServerTime, sInfo.ServerLoad)

}

func main() {

	site := make(map[string]string)

	site["localhost"] = "http://localhost:8080/"

	c := make(chan ServerInfo)

	for website, url := range site {
		go getHTTPContent(url, website, c)
	}

	for i := 0; i < (len(site)); i++ {
		sInfo := <-c
		printResult(sInfo)
	}

}
