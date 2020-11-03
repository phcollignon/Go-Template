# db.properties java file
# without whitespace control
{{ with .Db}}
    db.username={{ .Username }}
    {{ if eq $.Env "prod" }}
        # db.password= # loaded from os env
    {{ else }}
        db.password={{ .Password }}
    {{ end }}
    db.host={{ .Host }}
    db.port={{ .Port }}
{{ end}}


# db.properties java file
# without whitespace control
{{- with .Db }}
db.username={{ .Username }}
    {{- if eq $.Env "prod" }}
# db.password= # loaded from os env
    {{- else }}
db.password={{ .Password }}
    {{- end }}
db.host={{ .Host }}
db.port={{ .Port }}
{{- end}}



