package main

import (
	"fmt"
	"os/exec"
	"strings"
	"encoding/json"
	"net/http"
)




type ServerInfo struct {
	ServerLoad []string
        ServerTime string
}

func handler(w http.ResponseWriter, r *http.Request) {
    si, _ := json.Marshal(getServerInfo())
    fmt.Fprintf(w, "%s", si)
}





func getLoad() []string {
        loadcmd , _ := exec.Command("cat","/proc/loadavg").Output()
        load := strings.Split(string(loadcmd), " ")
        load = load[:3]
	return load
}

func getTime() string {
        timecmd , _ := exec.Command("/usr/bin/date").Output()
        return string(timecmd)
}


func getServerInfo() ServerInfo {
	sLoad := getLoad()
	sTime := getTime()
	return ServerInfo{ServerLoad: sLoad, ServerTime: sTime}
}


func main() {
//	fmt.Println("server info: ",getServerInfo())

//	si, _ := json.Marshal(getServerInfo())
//        fmt.Println("server info: ",si)

	 http.HandleFunc("/", handler)
	 http.ListenAndServe(":8080", nil)

}

