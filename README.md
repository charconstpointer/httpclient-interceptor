# Example usage of `http.Transport` to intercept requests
```go
var AuthMiddleware = func(r *http.Request) {
    log.Println("auth middleware, adding bearer token")
    r.Header.Add("Authorization", "Bearer 12345")
}

var SleepMiddleware = func(_ *http.Request) {
    log.Println("sleep middleware 3s")
    time.Sleep(time.Second * 3)
}

var HeadersLoggerMiddleware = func(r *http.Request) {
    log.Println(strings.Repeat("-", 20))
    defer log.Println(strings.Repeat("-", 20))
    log.Println(r.Header)
}

c := fleet.NewHTTPClient(AuthMiddleware, SleepMiddleware, HeadersLoggerMiddleware)
// c is of type *http.Client
```