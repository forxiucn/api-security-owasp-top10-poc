package org.example;

import org.springframework.web.bind.annotation.*;
import java.util.*;

@RestController
public class OwaspApiController {
    // 1. Broken Object Level Authorization
    @GetMapping("/api1/items/{itemId}")
    public Map<String, Object> api1(@PathVariable String itemId) {
        Map<String, Object> res = new HashMap<>();
        res.put("item_id", itemId);
        res.put("detail", "Object info (no auth check)");
        return res;
    }

    // 2. Broken User Authentication
    @PostMapping("/api2/login")
    public Map<String, Object> api2(@RequestBody Map<String, Object> body) {
        Map<String, Object> res = new HashMap<>();
        if ("admin".equals(body.get("username")) && "123456".equals(body.get("password"))) {
            res.put("msg", "Login success");
            res.put("token", "fake-jwt-token");
        } else {
            res.put("msg", "Login failed");
        }
        return res;
    }

    // 3. Excessive Data Exposure
    @GetMapping("/api3/userinfo")
    public Map<String, Object> api3() {
        Map<String, Object> res = new HashMap<>();
        res.put("username", "alice");
        res.put("email", "alice@example.com");
        res.put("password", "plaintextpassword");
        return res;
    }

    // 4. Lack of Resources & Rate Limiting
    @GetMapping("/api4/nolimit")
    public Map<String, Object> api4() {
        return Collections.singletonMap("msg", "No rate limit here");
    }

    // 5. Broken Function Level Authorization
    @GetMapping("/api5/admin")
    public Map<String, Object> api5() {
        return Collections.singletonMap("msg", "Admin function accessed");
    }

    // 6. Mass Assignment
    @PostMapping("/api6/profile")
    public Map<String, Object> api6(@RequestBody Map<String, Object> profile) {
        Map<String, Object> res = new HashMap<>();
        res.put("msg", "Profile updated");
        res.put("profile", profile);
        return res;
    }

    // 7. Security Misconfiguration
    @GetMapping("/api7/debug")
    public Map<String, Object> api7() {
        Map<String, Object> res = new HashMap<>();
        res.put("debug", true);
        res.put("config", "Sensitive config here");
        return res;
    }

    // 8. Injection
    @PostMapping("/api8/search")
    public Map<String, Object> api8(@RequestBody Map<String, Object> body) {
        String query = (String) body.getOrDefault("q", "");
        return Collections.singletonMap("result", "You searched for: " + query);
    }

    // 9. Improper Assets Management
    @GetMapping("/api9/old-api")
    public Map<String, Object> api9() {
        return Collections.singletonMap("msg", "Deprecated API still accessible");
    }

    // 10. Insufficient Logging & Monitoring
    @PostMapping("/api10/transfer")
    public Map<String, Object> api10(@RequestBody Map<String, Object> data) {
        Map<String, Object> res = new HashMap<>();
        res.put("msg", "Transfer completed");
        res.put("data", data);
        return res;
    }
}
