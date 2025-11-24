import java.io.*;
import java.net.HttpURLConnection;
import java.net.URL;

public class OwaspApi2023Client {
    private static final String BASE_URL = "http://127.0.0.1:8080";

    public static void main(String[] args) throws Exception {
        testApi1();
        testApi2();
        testApi3();
        testApi4();
        testApi5();
        testApi6();
        testApi7();
        testApi8();
        testApi9();
        testApi10();
    }

    private static void testApi1() throws Exception {
        System.out.println("API1: " + get("/api1/items/123"));
    }
    private static void testApi2() throws Exception {
        System.out.println("API2: " + post("/api2/login", "{\"username\":\"admin\",\"password\":\"123456\"}"));
        System.out.println("API2 (fail): " + post("/api2/login", "{\"username\":\"admin\",\"password\":\"wrong\"}"));
    }
    private static void testApi3() throws Exception {
        System.out.println("API3: " + get("/api3/userinfo"));
    }
    private static void testApi4() throws Exception {
        for (int i = 0; i < 3; i++) {
            System.out.println("API4 [" + (i+1) + "]: " + get("/api4/nolimit"));
        }
    }
    private static void testApi5() throws Exception {
        System.out.println("API5: " + get("/api5/admin"));
    }
    private static void testApi6() throws Exception {
        System.out.println("API6: " + post("/api6/transfer", "{\"from\":\"alice\",\"to\":\"bob\",\"amount\":100}"));
    }
    private static void testApi7() throws Exception {
        System.out.println("API7: " + post("/api7/ssrf", "{\"url\":\"http://example.com\"}"));
    }
    private static void testApi8() throws Exception {
        System.out.println("API8: " + get("/api8/debug"));
    }
    private static void testApi9() throws Exception {
        System.out.println("API9: " + get("/api9/old-api"));
    }
    private static void testApi10() throws Exception {
        System.out.println("API10: " + post("/api10/unsafe", "{\"external\":\"data\"}"));
    }

    private static String get(String path) throws Exception {
        URL url = new URL(BASE_URL + path);
        HttpURLConnection conn = (HttpURLConnection) url.openConnection();
        conn.setRequestMethod("GET");
        return readResponse(conn);
    }

    private static String post(String path, String json) throws Exception {
        URL url = new URL(BASE_URL + path);
        HttpURLConnection conn = (HttpURLConnection) url.openConnection();
        conn.setRequestMethod("POST");
        conn.setRequestProperty("Content-Type", "application/json");
        conn.setDoOutput(true);
        try (OutputStream os = conn.getOutputStream()) {
            os.write(json.getBytes("UTF-8"));
        }
        return readResponse(conn);
    }

    private static String readResponse(HttpURLConnection conn) throws Exception {
        int code = conn.getResponseCode();
        InputStream is = (code >= 200 && code < 400) ? conn.getInputStream() : conn.getErrorStream();
        BufferedReader in = new BufferedReader(new InputStreamReader(is));
        StringBuilder sb = new StringBuilder();
        String line;
        while ((line = in.readLine()) != null) {
            sb.append(line);
        }
        in.close();
        return code + ": " + sb.toString();
    }
}
