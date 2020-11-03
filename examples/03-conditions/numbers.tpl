{{ range .Number}}
    {{- if eq  .n1 .n2 }}
        {{- .n1}} = {{.n2}}
    {{- else}}
        {{- if lt .n1 .n2 }}
            {{- .n1}} < {{.n2}}
        {{- else}}
            {{- .n1}} > {{.n2}}
        {{- end}}
    {{- end}}
{{end}}