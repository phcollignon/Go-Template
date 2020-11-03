
Contact Card:
  - Name : {{ .Name }}, {{ .Firstname }}
  - Address : 
      {{- with .Address}}
        {{ .Street }}
        {{ .Postcode }}, {{ .City}}
      {{- end}}
  - Contact
        - Phones : 
            {{- range $key,$value := .Phones}}
            - {{ $value }} ({{ $key }})
            {{- end}}