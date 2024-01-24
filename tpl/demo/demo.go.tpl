package {{.packageName}}

import "github.com/dan-kuroto/gin-stronger/gs"

func init() {
	gs.UseController(&Controller)
}
