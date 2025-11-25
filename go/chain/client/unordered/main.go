package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	addr := flag.String("addr", "http://127.0.0.1:5050", "server base URL including protocol (e.g. http://host:port)")
	chainId := flag.String("chainId", "test-chain", "chain identifier for unordered calls")
	steps := flag.Int("steps", 5, "number of steps in the chain")
	flag.Parse()

	client := &http.Client{}
	order := make([]int, *steps)
	for i := 0; i < *steps; i++ {
		order[i] = i + 1
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(order), func(i, j int) { order[i], order[j] = order[j], order[i] })

	fmt.Printf("Running unordered client against %s (chainId=%s) steps=%d\n", *addr, *chainId, *steps)
	for _, step := range order {
		url := fmt.Sprintf("%s/chain/unordered/step/%d?chainId=%s", *addr, step, *chainId)
		resp, err := client.Get(url)
		if err != nil {
			fmt.Println("error calling step", step, err)
			os.Exit(1)
		}
		body, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Printf("step %d -> %d %s\n", step, resp.StatusCode, string(body))
		if resp.StatusCode != 200 {
			fmt.Println("unordered call returned non-200 at step", step)
		}
	}
	fmt.Println("unordered client finished")
}
