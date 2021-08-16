package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	EndPoint       = "https://webapi.kfca.com.sa:8081/api/Crossingtime/Current/"
	BackupEndPoint = "https://zahmaola-api.azurewebsites.net/ADS"
)

func getFromEndPoint(endpoint string) error {
	res, err := http.Get(endpoint)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var rjson []map[string]string
	if ok := json.Unmarshal(body, &rjson); ok != nil {
		return err
	}
	for _, i := range rjson {
		delete(i, "statusColor")
		fmt.Println()
		for key, value := range i {
			fmt.Println(key, "	", value)
		}
	}
	return nil
}

func main() {

	fmt.Println("Time in minutes")
	if err := getFromEndPoint(EndPoint); err != nil {
		// if the first end point returned err, backup endpoint will be used
		// if the backup also returned err, then it will be displayed
		if err := getFromEndPoint(BackupEndPoint); err != nil {
			panic(err)
		}
	}
}
