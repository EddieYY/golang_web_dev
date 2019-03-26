## [net/http](https://godoc.org/net/http) package in golang

1. type Handler <br/> 
[http.Handler](https://godoc.org/net/http#Handler) <br/>
Handler is a interface if any struct have ServerHTTP method.
   ```GO
   type Handler interface{
      ServerHTTP(ResposeWriter, *Request) 
   }
   ```
***
2. Server
[http.ListenAndServe](https://godoc.org/net/http#ListenAndServe)
   ``` Go
   func ListenAndServe(addr string, handler Handler) error
   ```

[http.ListenAndServeTLS](https://godoc.org/net/http#ListenAndServeTLS)
``` Go
func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error
```
*Notice that both of the above functions take a handler*
***
3. Request
To [http.Request](https://godoc.org/net/http#Request) see the detail </br>
There are a summary for http.Request.
   ```GO
   type Request struct {
    Method string
    URL *url.URL

    // Header = map[string][]string{
    // "Accept-Encoding": {"gzip, deflate"},
    //"Accept-Language": {"en-us"},
    // "Foo": {"Bar", "two"},
    //	}
    Header Header

    Body io.ReadCloser
    ContentLength int64
    Host string

    // This field is only available after ParseForm is called.
    Form url.Values

    // This field is only available after ParseForm is called.
    PostForm url.Values

    MultipartForm *multipart.Form
    // RemoteAddr allows HTTP servers and other software to record
    // the network address that sent the request, usually for logging. 
    RemoteAddr string
   }
   ```
* PostForm and Form

   ```http.Request``` is a struct. It has the fields ```Form``` & ```PostForm```. If we read the documentation on these, we see:

   ```