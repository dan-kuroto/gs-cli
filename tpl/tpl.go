package tpl

import _ "embed"

//go:embed banner.txt.tpl
var BannerTxt string

//go:embed go.mod.tpl
var GoMod string

//go:embed application.yml.tpl
var ApplicationYml string

//go:embed .gitignore.tpl
var GitIgnore string

//go:embed main.go.tpl
var MainGo string

//go:embed utils/config.go.tpl
var UtilsConfigGo string

//go:embed demo/demo.go.tpl
var DemoDemoGo string

//go:embed demo/controller.go.tpl
var DemoControllerGo string

//go:embed demo/model.go.tpl
var DemoModelGo string
