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

// ACWIFI method
func ACWIFI(code string, debug bool) {
	params := url.Values{}
	params.Set("redirurl", "http://www.apple.com/library/test/success.html")
	params.Set("zone", "lan")
	params.Set("auth_voucher", code)
	params.Set("accept", "Continuer / Continue")
	body := bytes.NewBufferString(params.Encode())

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("POST", "http://10.21.1.1:8002/", body)

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

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Connected")
	} else {
		fmt.Println("Failed:", resp.Status)
	}
}
