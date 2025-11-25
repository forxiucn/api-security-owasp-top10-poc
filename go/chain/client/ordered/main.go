package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", "http://127.0.0.1:5050", "server base URL including protocol (e.g. http://host:port)")
	chainId := flag.String("chainId", "test-chain", "chain identifier for ordered calls")
	steps := flag.Int("steps", 5, "number of steps in the chain")
	flag.Parse()

	client := &http.Client{}

	fmt.Printf("Running ordered client against %s (chainId=%s) steps=%d\n", *addr, *chainId, *steps)
	for i := 1; i <= *steps; i++ {
		url := fmt.Sprintf("%s/chain/ordered/step/%d?chainId=%s", *addr, i, *chainId)
		resp, err := client.Get(url)
		if err != nil {
			fmt.Println("error calling step", i, err)
			os.Exit(1)
		}
		body, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Printf("step %d -> %d %s\n", i, resp.StatusCode, string(body))
		if resp.StatusCode != 200 {
			fmt.Println("ordered call failed at step", i)
			os.Exit(2)
		}
	}
	fmt.Println("ordered client finished")
}
