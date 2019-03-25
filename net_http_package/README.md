## [net/http](https://godoc.org/net/http) package in golang

1. type Handler 
[http.Handler](https://godoc.org/net/http#Handler)
Handler is a interface if any struct have ServerHTTP method.
```GO```
type Handler interface{
   ServerHTTP(ResposeWriter, *Request) 
}

```



