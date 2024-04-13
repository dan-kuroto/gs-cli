package {{.packageName}}

import "github.com/dan-kuroto/gin-stronger/gs"

type controller struct{}

var Controller controller

func (*controller) GetRouter() gs.Router {
	return gs.Router{
		Path: "/{{.packageName}}",
		Children: []gs.Router{
			{
				Path:     "/hello",
				Method:   gs.GET | gs.POST,
				Handlers: gs.PackageHandlers(Controller.Hello),
			},
		},
	}
}

func (*controller) Hello(demo *DemoRequst) DemoResponse {
	return DemoResponse{Message: "Hello~" + demo.Name}
}
