package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	port := flag.String("port", "5019", "port to run the server on")
	contentPath := flag.String("contentPath", "", "content path prefix for API endpoints (e.g., /v1/api)")
	flag.Parse()

	fmt.Printf("Starting server on port %s, contentPath=%s\n", *port, *contentPath)

	r := gin.Default()

	// Helper function to add contentPath prefix to routes
	addRoute := func(method, path string, handler gin.HandlerFunc) {
		fullPath := *contentPath + path
		if method == "GET" {
			r.GET(fullPath, handler)
		} else if method == "POST" {
			r.POST(fullPath, handler)
		}
	}

	// 1. Broken Object Level Authorization
	addRoute("GET", "/api1/items/:itemId", func(c *gin.Context) {
		c.JSON(200, gin.H{"item_id": c.Param("itemId"), "detail": "Object info (no auth check)"})
	})

	// 2. Broken User Authentication
	addRoute("POST", "/api2/login", func(c *gin.Context) {
		var json map[string]interface{}
		c.BindJSON(&json)
		if json["username"] == "admin" && json["password"] == "123456" {
			c.JSON(200, gin.H{"msg": "Login success", "token": "fake-jwt-token"})
		} else {
			c.JSON(401, gin.H{"msg": "Login failed"})
		}
	})

	// 3. Excessive Data Exposure
	addRoute("GET", "/api3/userinfo", func(c *gin.Context) {
		c.JSON(200, gin.H{"username": "alice", "email": "alice@example.com", "password": "plaintextpassword"})
	})

	// 4. Lack of Resources & Rate Limiting
	addRoute("GET", "/api4/nolimit", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "No rate limit here"})
	})

	// 5. Broken Function Level Authorization
	addRoute("GET", "/api5/admin", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "Admin function accessed"})
	})

	// 6. Mass Assignment
	addRoute("POST", "/api6/profile", func(c *gin.Context) {
		var profile map[string]interface{}
		c.BindJSON(&profile)
		c.JSON(200, gin.H{"msg": "Profile updated", "profile": profile})
	})

	// 7. Security Misconfiguration
	addRoute("GET", "/api7/debug", func(c *gin.Context) {
		c.JSON(200, gin.H{"debug": true, "config": "Sensitive config here"})
	})

	// 8. Injection
	addRoute("POST", "/api8/search", func(c *gin.Context) {
		var json map[string]interface{}
		c.BindJSON(&json)
		c.JSON(200, gin.H{"result": "You searched for: " + json["q"].(string)})
	})

	// 9. Improper Assets Management
	addRoute("GET", "/api9/old-api", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "Deprecated API still accessible"})
	})

	// 10. Insufficient Logging & Monitoring
	addRoute("POST", "/api10/transfer", func(c *gin.Context) {
		var data map[string]interface{}
		c.BindJSON(&data)
		c.JSON(200, gin.H{"msg": "Transfer completed", "data": data})
	})

	r.Run(":" + *port)
}
