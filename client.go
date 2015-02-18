package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
//	"strings"
	"encoding/json"
//	"time"
)

/*
type ServerInfo struct {
	ServerName string
	ServerLoad string
	ServerTime string
	ServerResponseTime int64
}
*/

type ServerInfo struct {
        ServerLoad []string
        ServerTime string
}



func getHTTPContent(url string, website string, c chan ServerInfo) {
	var sInfo ServerInfo
	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &sInfo)
	c <- sInfo
}




func printResult(sInfo ServerInfo) {
	fmt.Printf("time: %v\nLoad: %v \n\n", sInfo.ServerTime, sInfo.ServerLoad)
//	fmt.Println("Json marshalized server object: ", string(si) ," \n")

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
