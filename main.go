package main


import (
	"fmt"
	"flag"
	"log"
	"net/http"
	"github.com/777777miSSU7777777/github-aggregator-web/app/view/index"
	"github.com/777777miSSU7777777/github-aggregator-web/app/api"
)

var host string
var port string

func init(){
	flag.StringVar(&host, "host", "127.0.0.1", "")
	flag.StringVar(&host, "h", "127.0.0.1", "")
	flag.StringVar(&port, "port", "8080", "")
	flag.StringVar(&port, "p", "8080", "")
	flag.Parse()
}

func main(){
	log.Printf("Server started on %s:%s",host,port)
	http.HandleFunc("/",index.Render)
	http.HandleFunc("/auth",api.Auth)
	http.HandleFunc("/logout",api.Logout)
	http.ListenAndServe(fmt.Sprintf("%s:%s",host,port), nil)
}