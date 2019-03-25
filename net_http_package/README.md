## [net/http](https://godoc.org/net/http) package in golang

1. type Handler <br/> 
[http.Handler](https://godoc.org/net/http#Handler) <br/>
Handler is a interface if any struct have ServerHTTP method.
```GO
type Handler interface{
   ServerHTTP(ResposeWriter, *Request) 
}

```

