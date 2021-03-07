package market

import "github.com/emicklei/go-restful"

type Route struct {
	Name    string
	Path    string
	Tag     string
	Reads   interface{}
	Writes  interface{}
	Role_id map[int]interface{}
	Handle  func(request *restful.Request, response *restful.Response)
}
