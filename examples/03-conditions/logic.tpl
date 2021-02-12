{{ range .Cases}}
    {{ .A }} and {{ .B }} is {{ and .A .B }}
    {{ .A }} or {{ .B }} is {{ or .A .B }}
    {{ .A }} and not {{ .B }} is {{ and .A ( not .B ) }}
{{end}}