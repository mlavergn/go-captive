package captive

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

/**
 * Unclear which Cisco AP is backing this, might be applicable
 * to more APs
 */

// PCWIFI method
func PCWIFI(urlStr string, code string, debug bool) Response {
	params := url.Values{}
	params.Set("redirurl", "http://www.apple.com/library/test/success.html")
	params.Set("zone", "lan")
	params.Set("auth_voucher", code)
	params.Set("accept", "Continuer / Continue")
	body := bytes.NewBufferString(params.Encode())

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("POST", urlStr, body)

	// Headers
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Debug
	if debug {
		fmt.Println("response Status:\n", resp.Status, "\nHeader:\n", resp.Header, "\nBody\n", string(respBody))
	}

	var response Response

	if resp.StatusCode == http.StatusOK {
		response.Message = "success"
		fmt.Println("Connected")
	} else {
		response.Message = "failed " + resp.Status
		fmt.Println("Failed:", resp.Status)
	}

	return response
}
