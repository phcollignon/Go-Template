
Dear {{.Name}},

Hello, 

We would like to assign some tasks for {{.Project}} project :
{{range .Topics}}
    - {{. }}
{{end}}
Important topic is : {{ index .Topics 1  }}

Can we plan a meeting on {{.Date}} ?

Regards,
