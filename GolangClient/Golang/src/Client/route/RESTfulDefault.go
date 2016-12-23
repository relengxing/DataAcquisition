package route

import (
    "net/http"
)

type RESTfulDefault struct {
    Child RESTfulInterface
    W     http.ResponseWriter
    R     *http.Request
}
type RESTfulInterface interface {
    Prepare()
    Transfer(w http.ResponseWriter,r *http.Request)
    Get()
    Post()
    Put()
    Delete()
    Patch()
    Head()
    Options()
    Finish()
    Render()
}

func (this *RESTfulDefault)Transfer(w http.ResponseWriter,r *http.Request){
    this.W = w
    this.R = r
    this.Child.Prepare()
    switch r.Method {
    case "GET":this.Child.Get()
    case "POST":this.Child.Post()
    case "PUT":this.Child.Put()
    case "DELETE":this.Child.Delete()
    case "HEAD":this.Child.Head()
    case "PATCH":this.Child.Patch()
    case "OPTIONS":this.Child.Options()
    }
    this.Child.Finish()
    this.Child.Render()
}

func (this *RESTfulDefault)Prepare(){
}
func (this *RESTfulDefault)Get(){
    http.Error(this.W, "Method Not Allowed", 405)
}
func (this *RESTfulDefault)Post(){
    http.Error(this.W, "Method Not Allowed", 405)
}
func (this *RESTfulDefault)Put(){
    http.Error(this.W, "Method Not Allowed", 405)
}
func (this *RESTfulDefault)Delete(){
    http.Error(this.W, "Method Not Allowed", 405)
}
func (this *RESTfulDefault)Patch(){
    http.Error(this.W, "Method Not Allowed", 405)
}
func (this *RESTfulDefault)Head(){
    http.Error(this.W, "Method Not Allowed", 405)
}
func (this *RESTfulDefault)Options(){
    http.Error(this.W, "Method Not Allowed", 405)
}
func (this *RESTfulDefault)Finish(){

}
func (this *RESTfulDefault)Render(){

}
