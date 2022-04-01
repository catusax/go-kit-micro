package template

// Module is the go.mod template used for new projects.
var Module = `module {{.Vendor}}{{.Service}}{{if .Client}}-client{{end}}

go 1.16

require (
)

replace {{.Vendor}}{{lower .Service}} => ../{{lower .Service}}
`
