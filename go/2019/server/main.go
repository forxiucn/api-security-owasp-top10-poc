package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 1. Broken Object Level Authorization
	r.GET("/api1/items/:itemId", func(c *gin.Context) {
		c.JSON(200, gin.H{"item_id": c.Param("itemId"), "detail": "Object info (no auth check)"})
	})

	// 2. Broken User Authentication
	r.POST("/api2/login", func(c *gin.Context) {
		var json map[string]interface{}
		c.BindJSON(&json)
		if json["username"] == "admin" && json["password"] == "123456" {
			c.JSON(200, gin.H{"msg": "Login success", "token": "fake-jwt-token"})
		} else {
			c.JSON(401, gin.H{"msg": "Login failed"})
		}
	})

	// 3. Excessive Data Exposure
	r.GET("/api3/userinfo", func(c *gin.Context) {
		c.JSON(200, gin.H{"username": "alice", "email": "alice@example.com", "password": "plaintextpassword"})
	})

	// 4. Lack of Resources & Rate Limiting
	r.GET("/api4/nolimit", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "No rate limit here"})
	})

	// 5. Broken Function Level Authorization
	r.GET("/api5/admin", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "Admin function accessed"})
	})

	// 6. Mass Assignment
	r.POST("/api6/profile", func(c *gin.Context) {
		var profile map[string]interface{}
		c.BindJSON(&profile)
		c.JSON(200, gin.H{"msg": "Profile updated", "profile": profile})
	})

	// 7. Security Misconfiguration
	r.GET("/api7/debug", func(c *gin.Context) {
		c.JSON(200, gin.H{"debug": true, "config": "Sensitive config here"})
	})

	// 8. Injection
	r.POST("/api8/search", func(c *gin.Context) {
		var json map[string]interface{}
		c.BindJSON(&json)
		c.JSON(200, gin.H{"result": "You searched for: " + json["q"].(string)})
	})

	// 9. Improper Assets Management
	r.GET("/api9/old-api", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "Deprecated API still accessible"})
	})

	// 10. Insufficient Logging & Monitoring
	r.POST("/api10/transfer", func(c *gin.Context) {
		var data map[string]interface{}
		c.BindJSON(&data)
		c.JSON(200, gin.H{"msg": "Transfer completed", "data": data})
	})

	r.Run(":5019")
}
