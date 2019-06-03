| directory | template               | values       | description                                           |
| --------- | ---------------------- | ------------ | ----------------------------------------------------- |
{{- range .Category }}
    {{- $category := .Name}}
    {{- range .Examples }}
|  {{$category}} | [{{ .Template }}](./src/examples/{{ $category }}/{{ .Template}}) | [{{.Values }}](./src/examples/{{ $category }}/{{.Values }}) | {{.Description }}  |
    {{- end }}     
{{- end }}
                           