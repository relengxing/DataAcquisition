package route

import "net/http"

type staticRoute struct {
    staticRouteList map[string]string
}
var StaticRoute staticRoute
func init() {
    StaticRoute = staticRoute{}
    StaticRoute.staticRouteList = make(map[string]string)
}

func (this *staticRoute)SetStaticRouteDir(url string,path string)  {
    this.staticRouteList[url] = path
}

func (this *staticRoute)RegisteStaticRoute()  {
    for k, v := range this.staticRouteList {
        http.Handle(k, http.FileServer(http.Dir(v)))
    }
}
