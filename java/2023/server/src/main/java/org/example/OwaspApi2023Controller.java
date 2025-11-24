package main.java.org.example;

import org.springframework.web.bind.annotation.*;
import java.util.*;

@RestController
public class OwaspApi2023Controller {
    // 1. Broken Object Level Authorization (BOLA)
    @GetMapping("/api1/items/{itemId}")
    public Map<String, Object> api1(@PathVariable String itemId) {
        Map<String, Object> res = new HashMap<>();
        res.put("item_id", itemId);
        res.put("detail", "Object info (no auth check)");
        return res;
    }

    // 2. Broken Authentication
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

    // 3. Broken Object Property Level Authorization (BOPLA)
    @GetMapping("/api3/userinfo")
    public Map<String, Object> api3() {
        Map<String, Object> res = new HashMap<>();
        res.put("username", "alice");
        res.put("email", "alice@example.com");
        res.put("role", "admin");
        res.put("salary", 10000);
        return res;
    }

    // 4. Unrestricted Resource Consumption
    @GetMapping("/api4/nolimit")
    public Map<String, Object> api4() {
        return Collections.singletonMap("msg", "No resource limit");
    }

    // 5. Broken Function Level Authorization
    @GetMapping("/api5/admin")
    public Map<String, Object> api5() {
        return Collections.singletonMap("msg", "Admin function accessed");
    }

    // 6. Unrestricted Access to Sensitive Business Flows
    @PostMapping("/api6/transfer")
    public Map<String, Object> api6(@RequestBody Map<String, Object> data) {
        Map<String, Object> res = new HashMap<>();
        res.put("msg", "Business flow executed");
        res.put("data", data);
        return res;
    }

    // 7. Server Side Request Forgery (SSRF)
    @PostMapping("/api7/ssrf")
    public Map<String, Object> api7(@RequestBody Map<String, Object> body) {
        String url = (String) body.getOrDefault("url", "");
        return Collections.singletonMap("msg", "Requested URL: " + url);
    }

    // 8. Security Misconfiguration
    @GetMapping("/api8/debug")
    public Map<String, Object> api8() {
        Map<String, Object> res = new HashMap<>();
        res.put("debug", true);
        res.put("config", "Sensitive config here");
        return res;
    }

    // 9. Improper Inventory Management
    @GetMapping("/api9/old-api")
    public Map<String, Object> api9() {
        return Collections.singletonMap("msg", "Deprecated API still accessible");
    }

    // 10. Unsafe Consumption of APIs
    @PostMapping("/api10/unsafe")
    public Map<String, Object> api10(@RequestBody Map<String, Object> data) {
        Map<String, Object> res = new HashMap<>();
        res.put("msg", "Unsafe API consumed");
        res.put("data", data);
        return res;
    }
}
