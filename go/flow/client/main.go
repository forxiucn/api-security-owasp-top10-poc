package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func must(v interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return v
}

func postJSON(client *http.Client, url string, body interface{}, token string) (*http.Response, []byte, error) {
	b, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", url, bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	data, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return resp, data, nil
}

func getJSON(client *http.Client, url string, token string) (*http.Response, []byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	if token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	data, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return resp, data, nil
}

func main() {
	addr := flag.String("addr", "http://127.0.0.1:5060", "flow server base URL including protocol")
	username := flag.String("username", "alice", "username")
	password := flag.String("password", "secret", "password")
	to := flag.String("to", "bob", "transfer destination")
	amount := flag.Float64("amount", 100.0, "transfer amount")
	queryPin := flag.String("query-pin", "1234", "PIN for query-pin step (leave empty to skip)")
	withdrawPin := flag.String("withdraw-pin", "2345", "PIN for withdraw-pin step")
	smsCode := flag.String("sms-code", "000000", "SMS code for sms-code step")
	flag.Parse()

	client := &http.Client{}
	var token string

	// 1. Three-step login
	// Step 1: Send credentials
	step1URL := fmt.Sprintf("%s/flow/login-step1", *addr)
	fmt.Println("POST", step1URL)
	resp, data, err := postJSON(client, step1URL, map[string]string{"username": *username, "password": *password}, "")
	if err != nil {
		panic(err)
	}
	fmt.Println("->", resp.Status, string(data))
	if resp.StatusCode != 200 {
		return
	}
	var d1 map[string]string
	json.Unmarshal(data, &d1)
	loginSessionId := d1["loginSessionId"]

	// Step 2: Request SMS
	step2URL := fmt.Sprintf("%s/flow/login-step2", *addr)
	fmt.Println("POST", step2URL)
	resp, data, err = postJSON(client, step2URL, map[string]string{"loginSessionId": loginSessionId}, "")
	if err != nil {
		panic(err)
	}
	fmt.Println("->", resp.Status, string(data))
	if resp.StatusCode != 200 {
		return
	}

	// Step 3: Submit with SMS code
	step3URL := fmt.Sprintf("%s/flow/login-step3", *addr)
	fmt.Println("POST", step3URL)
	resp, data, err = postJSON(client, step3URL, map[string]string{"loginSessionId": loginSessionId, "smsCode": *smsCode}, "")
	if err != nil {
		panic(err)
	}
	fmt.Println("->", resp.Status, string(data))
	if resp.StatusCode != 200 {
		return
	}
	var d3 map[string]string
	json.Unmarshal(data, &d3)
	token = d3["token"]
	fmt.Println("Got token:", token)

	// 2. optional query-pin (for demonstration)
	if *queryPin != "" {
		qpURL := fmt.Sprintf("%s/flow/query-pin", *addr)
		fmt.Println("POST", qpURL)
		resp, data, err = postJSON(client, qpURL, map[string]interface{}{"pin": *queryPin}, token)
		if err != nil {
			panic(err)
		}
		fmt.Println("->", resp.Status, string(data))
	}

	// 3. userinfo
	uiURL := fmt.Sprintf("%s/flow/userinfo", *addr)
	fmt.Println("GET", uiURL)
	resp, data, err = getJSON(client, uiURL, token)
	if err != nil {
		panic(err)
	}
	fmt.Println("->", resp.Status, string(data))

	// 4. balance
	balURL := fmt.Sprintf("%s/flow/balance", *addr)
	fmt.Println("GET", balURL)
	resp, data, err = getJSON(client, balURL, token)
	if err != nil {
		panic(err)
	}
	fmt.Println("->", resp.Status, string(data))

	// 5. initiate transfer
	initURL := fmt.Sprintf("%s/flow/initiate-transfer", *addr)
	fmt.Println("POST", initURL)
	resp, data, err = postJSON(client, initURL, map[string]interface{}{"to": *to, "amount": *amount}, token)
	if err != nil {
		panic(err)
	}
	fmt.Println("->", resp.Status, string(data))
	if resp.StatusCode != 200 {
		return
	}
	var d2 map[string]string
	json.Unmarshal(data, &d2)
	tr := d2["transferId"]

	// 6. withdraw PIN
	wpURL := fmt.Sprintf("%s/flow/withdraw-pin", *addr)
	fmt.Println("POST", wpURL)
	resp, data, err = postJSON(client, wpURL, map[string]interface{}{"transferId": tr, "pin": *withdrawPin}, token)
	if err != nil {
		panic(err)
	}
	fmt.Println("->", resp.Status, string(data))
	if resp.StatusCode != 200 {
		return
	}

	// 7. sms code
	smsURL := fmt.Sprintf("%s/flow/sms-code", *addr)
	fmt.Println("POST", smsURL)
	resp, data, err = postJSON(client, smsURL, map[string]interface{}{"transferId": tr, "code": *smsCode}, token)
	if err != nil {
		panic(err)
	}
	fmt.Println("->", resp.Status, string(data))
	if resp.StatusCode != 200 {
		return
	}

	// 8. submit transfer
	subURL := fmt.Sprintf("%s/flow/submit-transfer", *addr)
	fmt.Println("POST", subURL)
	resp, data, err = postJSON(client, subURL, map[string]interface{}{"transferId": tr}, token)
	if err != nil {
		panic(err)
	}
	fmt.Println("->", resp.Status, string(data))

	// 9. balance after
	fmt.Println("GET", balURL)
	resp, data, err = getJSON(client, balURL, token)
	if err != nil {
		panic(err)
	}
	fmt.Println("->", resp.Status, string(data))
}
