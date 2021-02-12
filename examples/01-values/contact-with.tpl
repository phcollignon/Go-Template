
Contact Card:
  - Name : {{ .Name }}, {{ .Firstname }}
  - Address : 
      {{- with .Address}}
        {{ .Street }}
        {{ .Postcode }}, {{ .City}}
      {{- end}}
  - Contact
        - Phones : 
            {{- with .Phones}}
            - {{ .Private }} (Private)
            - {{ .Work }} (Work)
            - {{ .Mobile }} (Mobile)
            {{- end}}


