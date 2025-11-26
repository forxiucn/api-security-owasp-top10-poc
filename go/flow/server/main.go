package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type TransferState struct {
	ID               string  `json:"id"`
	Owner            string  `json:"owner"`
	To               string  `json:"to"`
	Amount           float64 `json:"amount"`
	Stage            int     `json:"stage"` // 0: initiated, 1: withdrawPinOK, 2: smsOK, 3: submitted
	QueryPinVerified bool    `json:"queryPinVerified"`
	WithdrawPinOK    bool    `json:"withdrawPinOk"`
	SmsVerified      bool    `json:"smsVerified"`
}

type LoginSession struct {
	Username string
	State    int // 0: credentials received, 1: sms sent, 2: verified
}

var sessMu sync.Mutex
var sessions = map[string]string{}             // token -> username (verified sessions)
var loginSessions = map[string]*LoginSession{} // loginSessionId -> login state

var balMu sync.Mutex
var balances = map[string]float64{
	"alice": 1000.0,
}

var trMu sync.Mutex
var transfers = map[string]*TransferState{}
var trCounter = 0

var loginCounter = 0

// RSA keys
var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

// Initialize RSA keys
func init() {
	var err error
	privateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	publicKey = &privateKey.PublicKey
}

// Helper: extract Authorization header token
func getToken(c *gin.Context) string {
	auth := c.GetHeader("Authorization")
	if len(auth) > 7 && auth[:7] == "Bearer " {
		return auth[7:]
	}
	return ""
}

// Helper: verify session token and get username
func getUser(c *gin.Context) (string, bool) {
	token := getToken(c)
	if token == "" {
		return "", false
	}
	sessMu.Lock()
	user := sessions[token]
	sessMu.Unlock()
	return user, user != ""
}

func main() {
	r := gin.Default()

	// Endpoint to get public key
	r.GET("/flow/public-key", func(c *gin.Context) {
		pubKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to marshal public key"})
			return
		}
		
		pubKeyPEM := pem.EncodeToMemory(&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: pubKeyBytes,
		})
		
		c.JSON(http.StatusOK, gin.H{
			"publicKey": string(pubKeyPEM),
		})
	})

	// Step 1: Initiate login with encrypted credentials (returns loginSessionId)
	r.POST("/flow/login-step1", func(c *gin.Context) {
		var body struct {
			Username    string `json:"username"`
			EncPassword string `json:"encPassword"` // Encrypted password
		}
		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
			return
		}
		
		// Base64 decode the encrypted password
		encData, err := base64.StdEncoding.DecodeString(body.EncPassword)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "decryption failed"})
			return
		}
		
		// Decrypt password
		passwordBytes, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, []byte(encData))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "decryption failed"})
			return
		}
		password := string(passwordBytes)
		
		// Simple auth: alice / secret
		if body.Username != "alice" || password != "secret" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}
		// Create a login session
		loginCounter++
		loginSessId := fmt.Sprintf("login_%d", loginCounter)
		sessMu.Lock()
		loginSessions[loginSessId] = &LoginSession{Username: body.Username, State: 0}
		sessMu.Unlock()
		c.JSON(http.StatusOK, gin.H{"loginSessionId": loginSessId})
	})

	// Step 2: Request SMS code (updates state to 1)
	r.POST("/flow/login-step2", func(c *gin.Context) {
		var body struct {
			LoginSessionId string `json:"loginSessionId"`
		}
		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
			return
		}
		sessMu.Lock()
		sess := loginSessions[body.LoginSessionId]
		sessMu.Unlock()
		if sess == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "login session not found"})
			return
		}
		if sess.State != 0 {
			c.JSON(http.StatusConflict, gin.H{"error": "invalid state"})
			return
		}
		// Mark as SMS sent
		sessMu.Lock()
		sess.State = 1
		sessMu.Unlock()
		// In real world, send SMS; here we just return success
		c.JSON(http.StatusOK, gin.H{"status": "sms_sent"})
	})

	// Step 3: Submit login with SMS code (returns token)
	r.POST("/flow/login-step3", func(c *gin.Context) {
		var body struct {
			LoginSessionId string `json:"loginSessionId"`
			SmsCode        string `json:"smsCode"`
		}
		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
			return
		}
		sessMu.Lock()
		sess := loginSessions[body.LoginSessionId]
		sessMu.Unlock()
		if sess == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "login session not found"})
			return
		}
		if sess.State != 1 {
			c.JSON(http.StatusConflict, gin.H{"error": "invalid state"})
			return
		}
		// Verify SMS code: 000000
		if body.SmsCode != "000000" {
			c.JSON(http.StatusForbidden, gin.H{"error": "wrong sms code"})
			return
		}
		// Create session token
		token := fmt.Sprintf("token_%s_%d", sess.Username, loginCounter)
		sessMu.Lock()
		sessions[token] = sess.Username
		delete(loginSessions, body.LoginSessionId) // Clean up login session
		sessMu.Unlock()
		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	// User info (use Authorization header)
	r.GET("/flow/userinfo", func(c *gin.Context) {
		user, ok := getUser(c)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"username": user, "accountId": "acc_" + user})
	})

	// Balance (use Authorization header)
	r.GET("/flow/balance", func(c *gin.Context) {
		user, ok := getUser(c)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		balMu.Lock()
		b := balances[user]
		balMu.Unlock()
		c.JSON(http.StatusOK, gin.H{"balance": b})
	})

	// Query PIN verification (simulates entering a PIN to view details, use Authorization header)
	r.POST("/flow/query-pin", func(c *gin.Context) {
		_, ok := getUser(c)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		var body struct {
			Pin string `json:"pin"`
		}
		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
			return
		}
		// simple PIN check: 1234
		if body.Pin != "1234" {
			c.JSON(http.StatusForbidden, gin.H{"error": "wrong pin"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Initiate transfer: create transfer record and return transferId (use Authorization header)
	r.POST("/flow/initiate-transfer", func(c *gin.Context) {
		user, ok := getUser(c)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		var body struct {
			To     string  `json:"to"`
			Amount float64 `json:"amount"`
		}
		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
			return
		}
		trMu.Lock()
		trCounter++
		id := fmt.Sprintf("tr_%d", trCounter)
		t := &TransferState{ID: id, Owner: user, To: body.To, Amount: body.Amount, Stage: 0}
		transfers[id] = t
		trMu.Unlock()
		c.JSON(http.StatusOK, gin.H{"transferId": id})
	})

	// Verify withdrawal PIN for a transfer (must be in stage 0, use Authorization header)
	r.POST("/flow/withdraw-pin", func(c *gin.Context) {
		user, ok := getUser(c)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		var body struct {
			TransferId string `json:"transferId"`
			Pin        string `json:"pin"`
		}
		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
			return
		}
		trMu.Lock()
		t := transfers[body.TransferId]
		if t == nil {
			trMu.Unlock()
			c.JSON(http.StatusNotFound, gin.H{"error": "transfer not found"})
			return
		}
		if t.Owner != user {
			trMu.Unlock()
			c.JSON(http.StatusForbidden, gin.H{"error": "not owner"})
			return
		}
		if t.Stage != 0 {
			trMu.Unlock()
			c.JSON(http.StatusConflict, gin.H{"error": "invalid stage"})
			return
		}
		// check pin
		if body.Pin != "2345" {
			trMu.Unlock()
			c.JSON(http.StatusForbidden, gin.H{"error": "wrong withdraw pin"})
			return
		}
		t.WithdrawPinOK = true
		t.Stage = 1
		trMu.Unlock()
		c.JSON(http.StatusOK, gin.H{"status": "withdraw_pin_ok"})
	})

	// Verify SMS code for a transfer (must be stage 1, use Authorization header)
	r.POST("/flow/sms-code", func(c *gin.Context) {
		user, ok := getUser(c)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		var body struct {
			TransferId string `json:"transferId"`
			Code       string `json:"code"`
		}
		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
			return
		}
		trMu.Lock()
		t := transfers[body.TransferId]
		if t == nil {
			trMu.Unlock()
			c.JSON(http.StatusNotFound, gin.H{"error": "transfer not found"})
			return
		}
		if t.Owner != user {
			trMu.Unlock()
			c.JSON(http.StatusForbidden, gin.H{"error": "not owner"})
			return
		}
		if t.Stage != 1 {
			trMu.Unlock()
			c.JSON(http.StatusConflict, gin.H{"error": "invalid stage"})
			return
		}
		// code check: 000000
		if body.Code != "000000" {
			trMu.Unlock()
			c.JSON(http.StatusForbidden, gin.H{"error": "wrong sms code"})
			return
		}
		t.SmsVerified = true
		t.Stage = 2
		trMu.Unlock()
		c.JSON(http.StatusOK, gin.H{"status": "sms_ok"})
	})

	// Submit transfer (finalize) - must be stage 2 (use Authorization header)
	r.POST("/flow/submit-transfer", func(c *gin.Context) {
		user, ok := getUser(c)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		var body struct {
			TransferId string `json:"transferId"`
		}
		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
			return
		}
		trMu.Lock()
		t := transfers[body.TransferId]
		if t == nil {
			trMu.Unlock()
			c.JSON(http.StatusNotFound, gin.H{"error": "transfer not found"})
			return
		}
		if t.Owner != user {
			trMu.Unlock()
			c.JSON(http.StatusForbidden, gin.H{"error": "not owner"})
			return
		}
		if t.Stage != 2 {
			trMu.Unlock()
			c.JSON(http.StatusConflict, gin.H{"error": "invalid stage"})
			return
		}
		// perform transfer (deduct)
		balMu.Lock()
		if balances[user] < t.Amount {
			balMu.Unlock()
			trMu.Unlock()
			c.JSON(http.StatusPaymentRequired, gin.H{"error": "insufficient funds"})
			return
		}
		balances[user] = balances[user] - t.Amount
		balMu.Unlock()
		t.Stage = 3
		trMu.Unlock()
		c.JSON(http.StatusOK, gin.H{"status": "submitted", "transferId": t.ID})
	})

	// helper: get transfer state (for testing)
	r.GET("/flow/transfer/:id", func(c *gin.Context) {
		id := c.Param("id")
		trMu.Lock()
		t := transfers[id]
		trMu.Unlock()
		if t == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		c.JSON(http.StatusOK, t)
	})

	r.GET("/health", func(c *gin.Context) { c.String(http.StatusOK, "ok") })

	r.Run(":5060")
}
