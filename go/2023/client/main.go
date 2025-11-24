package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL = "http://127.0.0.1:5023"

func main() {
	testApi1()
	testApi2()
	testApi3()
	testApi4()
	testApi5()
	testApi6()
	testApi7()
	testApi8()
	testApi9()
	testApi10()
}

func testApi1() { getPrint("/api1/items/123") }
func testApi2() {
	postPrint("/api2/login", map[string]interface{}{"username": "admin", "password": "123456"})
	postPrint("/api2/login", map[string]interface{}{"username": "admin", "password": "wrong"})
}
func testApi3() { getPrint("/api3/userinfo") }
func testApi4() {
	for i := 0; i < 3; i++ {
		getPrint("/api4/nolimit")
	}
}
func testApi5() { getPrint("/api5/admin") }
func testApi6() { postPrint("/api6/transfer", map[string]interface{}{"from": "alice", "to": "bob", "amount": 100}) }
func testApi7() { postPrint("/api7/ssrf", map[string]interface{}{"url": "http://example.com"}) }
func testApi8() { getPrint("/api8/debug") }
func testApi9() { getPrint("/api9/old-api") }
func testApi10() { postPrint("/api10/unsafe", map[string]interface{}{"external": "data"}) }

func getPrint(path string) {
	resp, err := http.Get(baseURL + path)
	if err != nil {
		fmt.Println(path, "error:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s: %d %s\n", path, resp.StatusCode, string(body))
}

func postPrint(path string, data map[string]interface{}) {
	b, _ := json.Marshal(data)
	resp, err := http.Post(baseURL+path, "application/json", bytes.NewBuffer(b))
	if err != nil {
		fmt.Println(path, "error:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s: %d %s\n", path, resp.StatusCode, string(body))
}
