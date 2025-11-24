package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 1. Broken Object Level Authorization (BOLA)
	r.GET("/api1/items/:itemId", func(c *gin.Context) {
		c.JSON(200, gin.H{"item_id": c.Param("itemId"), "detail": "Object info (no auth check)"})
	})

	// 2. Broken Authentication
	r.POST("/api2/login", func(c *gin.Context) {
		var json map[string]interface{}
		c.BindJSON(&json)
		if json["username"] == "admin" && json["password"] == "123456" {
			c.JSON(200, gin.H{"msg": "Login success", "token": "fake-jwt-token"})
		} else {
			c.JSON(401, gin.H{"msg": "Login failed"})
		}
	})

	// 3. Broken Object Property Level Authorization (BOPLA)
	r.GET("/api3/userinfo", func(c *gin.Context) {
		c.JSON(200, gin.H{"username": "alice", "email": "alice@example.com", "role": "admin", "salary": 10000})
	})

	// 4. Unrestricted Resource Consumption
	r.GET("/api4/nolimit", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "No resource limit"})
	})

	// 5. Broken Function Level Authorization
	r.GET("/api5/admin", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "Admin function accessed"})
	})

	// 6. Unrestricted Access to Sensitive Business Flows
	r.POST("/api6/transfer", func(c *gin.Context) {
		var data map[string]interface{}
		c.BindJSON(&data)
		c.JSON(200, gin.H{"msg": "Business flow executed", "data": data})
	})

	// 7. Server Side Request Forgery (SSRF)
	r.POST("/api7/ssrf", func(c *gin.Context) {
		var json map[string]interface{}
		c.BindJSON(&json)
		c.JSON(200, gin.H{"msg": "Requested URL: " + json["url"].(string)})
	})

	// 8. Security Misconfiguration
	r.GET("/api8/debug", func(c *gin.Context) {
		c.JSON(200, gin.H{"debug": true, "config": "Sensitive config here"})
	})

	// 9. Improper Inventory Management
	r.GET("/api9/old-api", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "Deprecated API still accessible"})
	})

	// 10. Unsafe Consumption of APIs
	r.POST("/api10/unsafe", func(c *gin.Context) {
		var data map[string]interface{}
		c.BindJSON(&data)
		c.JSON(200, gin.H{"msg": "Unsafe API consumed", "data": data})
	})

	r.Run(":5023")
}
