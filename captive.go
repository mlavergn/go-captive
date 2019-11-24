package main

import (
	"captive"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Enum of known SSIDs
const (
	pteclrWiFi string = "pteclrWIFI"
	rougeWifi  string = "RougeWiFi"
	unitedWifi string = "United_Wi-F"
)

/**
 * Handle PC WiFi Request
 */
func handlerRoot(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("handlerRoot")
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.WriteHeader(http.StatusOK)

	resp.Write([]byte(captive.PortalHTML))
}

/**
 * Handle captive page with access code
 */
func handlerCaptive(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("handlerCaptive")
	defer req.Body.Close()

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}

	var captiveMessage captive.Request
	err = json.Unmarshal(data, &captiveMessage)
	if err != nil {
		fmt.Println(err)
	}

	urlStr := captiveMessage.Data[captive.URL]
	code := captiveMessage.Data[captive.Code]
	fmt.Println(urlStr, code)

	captiveResponse := captive.PCWIFI(urlStr, code, false)

	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)

	data, err = json.Marshal(captiveResponse)
	resp.Write(data)
}

/**
 * Parse CLI args and POST to captive page
 */
func main() {
	http.Handle("/", http.HandlerFunc(handlerRoot))
	http.Handle("/captive", http.HandlerFunc(handlerCaptive))

	http.ListenAndServe(":8888", nil)
}
