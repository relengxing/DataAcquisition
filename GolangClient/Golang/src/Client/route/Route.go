package route

import (
    "net/http"
    "html/template"
    "fmt"
)
var TAG string = "Route==>"
type route struct {
    
}
var Route route

func init() {
    Route = route{}
}
func (this *route)register()  {
    StaticRoute.RegisteStaticRoute()
    FixedRoute.RegisteFixedRoute()
    RESTfulRoute.RegisteRESTfulRoute()
}

func (this *route)Start(port string)  {
    this.register()
    http.ListenAndServe(port,nil)
}

func NotfoundHandler(w http.ResponseWriter, r *http.Request) {
    t,err := template.ParseFiles("static/html/404.html")
    if err != nil {
        fmt.Println(TAG,err)
    }
    t.Execute(w,nil)
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
    t,err := template.ParseFiles("static/html/welcome.html")
    if err != nil {
        fmt.Println(TAG,err)
    }
    t.Execute(w,nil)
}