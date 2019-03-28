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
   * Notice that both of the above functions take a handler*
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
    Form
    Form contains the parsed form data, including both the ***URL*** field's query parameters and the POST or PUT form data.
    This field is only available after **ParseForm** is called.
    The HTTP client ignores Form and uses Body instead.
    Form url.Values
   
    PostForm contains the parsed form data from POST, PATCH or PUT body parameters.
    This field is only available after **ParseForm** is called.
    The HTTP client ignores PostForm and uses Body instead.
    PostForm url.Values
    ```
    *ParseForm*
    ```go 
    func (r *Request) ParseForm() error 
    ```

     we see that this is a method attached to a *http.Request.

 * Http.Method
      The http.Request type is a struct which has a Method field.

* URL Values
      The http.Request type is a struct which has a URL field. Notice that the type is a [*url.URL](https://godoc.org/net/url#URL)
```go
type URL struct {
    Scheme     string
    Opaque     string    // encoded opaque data
    User       *Userinfo // username and password information
    Host       string    // host or host:port
    Path       string
    RawPath    string // encoded path hint (Go 1.5 and later only; see EscapedPath method)
    ForceQuery bool   // append a query ('?') even if RawQuery is empty
    RawQuery   string // encoded query values, without '?'
    Fragment   string // fragment for references, without '#'
}
```
The general form represented is:
```go
[scheme:][//[userinfo@]host][/]path[?query][#fragment]
```
* Work with the HTTP header
The http.Request type is a struct which has a Header field.
From the documentation about the http.Header struct, we see that:
  ```go 
  type Header map[string][]string
   ```