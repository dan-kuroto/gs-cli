# env:
#   active: dev

gin:
  release: false
  host: 127.0.0.1
  port: 5480

snow-flake:
  start-stmp: 1626779686000{{if .customConfig}}

message: hello world
{{else}}
{{end}}