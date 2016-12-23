package route

import "net/http"

type rESTfulRoute struct {
    rESTfulRouteList map[string]RESTfulInterface
}

var RESTfulRoute rESTfulRoute
func init() {
    RESTfulRoute = rESTfulRoute{}
    RESTfulRoute.rESTfulRouteList = make(map[string]RESTfulInterface)
}

func (this *rESTfulRoute)SetRESTfulRouteDir(url string,restful RESTfulInterface)  {
    this.rESTfulRouteList[url] = restful
}

func (this *rESTfulRoute)RegisteRESTfulRoute()  {
    for k, v := range this.rESTfulRouteList {
        http.HandleFunc(k,v.Transfer)
    }
}


