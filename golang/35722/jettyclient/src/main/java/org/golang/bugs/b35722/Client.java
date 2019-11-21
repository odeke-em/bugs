package org.golang.bugs.b35722;

import java.nio.charset.StandardCharsets;
import org.eclipse.jetty.client.HttpClient;
import org.eclipse.jetty.client.HttpRequest;
import org.eclipse.jetty.util.ssl.SslContextFactory;
import org.eclipse.jetty.client.api.ContentResponse;

public class Client {
    public static void main(String[] args) {
        SslContextFactory insecureSkipVerifyFactory = new SslContextFactory.Client(true);
        HttpClient httpClient = new HttpClient(insecureSkipVerifyFactory);

        try {
            httpClient.start();
        } catch (Exception e) {
            System.err.println("Failed to start the HTTP client " + e);
            return;
        }

        String url = "https://localhost:8888/hello";
        if (args.length > 0 && !args[0].isEmpty())
            url = args[0];

        HttpRequest req = (HttpRequest) httpClient.newRequest(url);
        try {
            ContentResponse res = req.send();
            String payload = new String(res.getContent(), StandardCharsets.UTF_8);
            System.out.println("Response:\n" + payload);
        } catch (Exception e) {
            System.err.println("Exception instead of response " + e);
        }
    }
}
