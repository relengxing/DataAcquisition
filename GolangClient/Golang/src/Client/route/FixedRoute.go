package route

import "net/http"

type fixedRoute struct {
    fixedRouteList map[string]func(http.ResponseWriter, *http.Request)
}
var FixedRoute fixedRoute
func init() {
    FixedRoute = fixedRoute{}
    FixedRoute.fixedRouteList = make(map[string]func(http.ResponseWriter, *http.Request))
}

func (this *fixedRoute)SetFixedRouteDir(url string,handle func(http.ResponseWriter, *http.Request))  {
    this.fixedRouteList[url] = handle
}

func (this *fixedRoute)RegisteFixedRoute()  {
    for k, v := range this.fixedRouteList {
        http.HandleFunc(k, v)
    }
}
