package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const path = "./cache"

type ProxySwicher struct {
}

func (*ProxySwicher) Init() {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.Mkdir(path, 0700)
	}
	var client = &http.Client{}
	url := "https://proxylist.geonode.com/api/proxy-list?limit=500&page=1&sort_by=lastChecked&sort_type=desc"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	// fmt.Print(string(responseData))

	var responseObject Proxylist

	json.Unmarshal(responseData, &responseObject)
	// s := responseObject.Data
	fmt.Printf("cap %v, len %v, %p\n", cap(responseData), len(responseData), responseData)

}
