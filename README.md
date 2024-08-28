# jwt_auth_fiber

## Kurulum (Visual Studio Code terminalinde)

1. Bağımlılıkları yükleyin:
   ```terminal
   go mod tidy
   ```

2. Uygulamayı başlatın:
   ```terminal
   go run main.go
   ```

Postman kullanarak deneyebilirsiniz. POST kısmında http://localhost:8080/api/v1/Login yazıp deneyin. Ortaya çıkan JWT token ile de GET kısmında http://localhost:8080/api/v1/protected yazıp eklediğiniz kişnin bilgilerini görebilirsiniz.

(Bearer Token kullanımı Linki) https://rebeccaminx922.medium.com/bearer-token-ile-postman-kullan%C4%B1m%C4%B1-temel-rehberiniz-3ae6f62e2bf0
