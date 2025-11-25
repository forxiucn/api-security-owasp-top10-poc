package main

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

// orderedChains stores the last successfully executed step for each chainId
var orderedMu sync.Mutex
var orderedChains = map[string]int{}

func main() {
	r := gin.Default()

	// Ordered chain: enforces sequence
	r.GET("/chain/ordered/step/:n", func(c *gin.Context) {
		nStr := c.Param("n")
		n, err := strconv.Atoi(nStr)
		if err != nil || n < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid step"})
			return
		}
		chainId := c.Query("chainId")
		if chainId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing chainId query param"})
			return
		}

		orderedMu.Lock()
		last := orderedChains[chainId]
		// expect next step == last + 1
		if n != last+1 {
			orderedMu.Unlock()
			c.JSON(http.StatusConflict, gin.H{"error": "wrong order", "expected_next": last + 1, "received": n})
			return
		}
		orderedChains[chainId] = n
		orderedMu.Unlock()

		c.JSON(http.StatusOK, gin.H{"status": "ok", "step": n})
	})

	// Unordered chain: accepts any step in any order
	r.GET("/chain/unordered/step/:n", func(c *gin.Context) {
		nStr := c.Param("n")
		n, err := strconv.Atoi(nStr)
		if err != nil || n < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid step"})
			return
		}
		chainId := c.Query("chainId")
		if chainId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing chainId query param"})
			return
		}
		// For unordered we simply accept and return success
		c.JSON(http.StatusOK, gin.H{"status": "ok", "step": n})
	})

	// Reset endpoint to clear a chain state (for testing)
	r.POST("/chain/ordered/reset", func(c *gin.Context) {
		chainId := c.Query("chainId")
		if chainId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing chainId"})
			return
		}
		orderedMu.Lock()
		delete(orderedChains, chainId)
		orderedMu.Unlock()
		c.JSON(http.StatusOK, gin.H{"status": "reset", "chainId": chainId})
	})

	// simple health
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	// default port 5050 (configurable via environment or reverse proxy)
	r.Run(":5050")
}
