package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type configArgs struct {
	hostinf hostInf
	dsn     dsn
}
type hostInf struct {
	HostIP   string `json:"host_ip"`
	HostPort string `json:"host_port"`
}
type dsn struct {
	User     string `json:"db_user"`
	Pwd      string `json:"db_pwd"`
	Hostname string `json:"db_hostname"`
	Port     string `json:"db_port"`
	Schema   string `json:"db_schema"`
}

func main() {
	jsonFile, _ := os.Open("test.json")
	defer jsonFile.Close()
	jsonData, _ := ioutil.ReadAll(jsonFile)
	var data hostInf
	json.Unmarshal(jsonData, &data)
	fmt.Println(data)

}
