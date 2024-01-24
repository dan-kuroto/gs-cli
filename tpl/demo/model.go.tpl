package {{.packageName}}

type DemoRequst struct {
	Name string `form:"name"`
}

type DemoResponse struct {
	Message string `json:"message"`
}
