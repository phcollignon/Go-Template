
{{range  .Schema.Table}}
	CREATE TABLE {{ .Name}} (
		{{range $idx, $column := .Column}}
			{{if  $idx }} , {{end}}{{ $column.Name}} {{$column.Type}} {{if $column.Size }} {{$column.Size}} {{end}}
		{{end}}
		)
{{end}}
