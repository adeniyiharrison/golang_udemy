# Introduction to Mircoservices
* https://www.youtube.com/watch?v=VzBGi_n65iU&list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_

### Creating an HTTP Server
* `net/http.ListenAndServe`
    * https://pkg.go.dev/net/http#ListenAndServe
    * `func ListenAndServe(addr string, handler Handler) error`
        * `addr` is the bind address // "bind this to every IP address on my machine to port 9090" 
            * could bind to explicit IP address like `127.0.0.1:9090` etc
    ```
    func main() {
        http.ListenAndServe(":9090", nil)
    }    
    ```
    * run this and then `curl -v localhost:9090`
        * it shows we have connected but then finishes with 404
        * Basically because theres nothing in out program
* How do we Handle requests?
    ```
    func main() {
        http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
            log.Println("This is a test")
        })
        http.ListenAndServe(":9090", nil)
    }
    ```
    * https://pkg.go.dev/net/http#HandleFunc
    * `func HandleFunc(pattern string, handler func(ResponseWriter, *Request))`
    * HandleFunc registers the handler function for the given pattern in the DefaultServeMux.
        * ServeMux is responsible to redirecting paths
        * When a request comes in that matches this path it runs the function
    * ServeMux is similar to Handler and a Handler is an interface needing the ServeHTTP method
        * https://pkg.go.dev/net/http#Handler
        ```
        type Handler interface {
            ServeHTTP(ResponseWriter, *Request)
        }
        ```
    * When we try the curl command again we get a 200 and the terminal page running the service logged out "this is a test"
    * __TLDR the `http.HandleFunc` takes the annoymous function, creates a handler type from it and it adds it to the DefaultServeMux__

* HTTP ResponseWriter
    * https://pkg.go.dev/net/http#ResponseWriter
    * Allows you to respond back to the client with status codes, errors, etc
    
* HTTP Requests
    * Pull data in from the curl commands

```
func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("This is a test")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Oops", http.StatusBadRequest)
		}
		log.Printf("Data: %s\n", d)
		fmt.Fprintf(rw, "Sent from the HandleFunc: %s\n", d)
		rw.WriteHeader(http.StatusAccepted)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("This test is over")
	})
	http.ListenAndServe(":9090", nil)
}
```

# Part Two
### Server Type
* https://pkg.go.dev/net/http#Server


