
{{- template "mydomain.subchart" }}
{{- template "mydomain.subchart" . }}
{{- template "mydomain.subchart" .Address }}



{{- define "mydomain.subchart" }}
 - labels:
    app: nginx
    {{- if .Name}}
    name: .Name
    {{- end}}
    street: {{ .Street }} {{ .City }}
{{- end }}