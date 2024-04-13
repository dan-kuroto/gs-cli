package main

import (
	_ "{{.projectName}}/demo"{{if .customConfig}}
	"{{.projectName}}/utils"{{end}}
	"net/http"
	"github.com/dan-kuroto/gin-stronger/gs"
	"github.com/gin-gonic/gin"
)

func panicStringHandler(c *gin.Context, err string) {
	c.JSON(http.StatusInternalServerError, gin.H{"err": err})
}

func main() {
	gs.SetGlobalPreffix("/api")
	gs.AddGlobalMiddleware(gs.PackagePanicHandler(panicStringHandler))
    {{if .customConfig}}gs.RunApp(&utils.Config){{else}}gs.RunAppDefault(){{end}}
}
