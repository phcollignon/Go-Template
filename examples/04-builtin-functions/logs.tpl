First user :
{{ (index .Logs 0   ).last_name }}
{{ (0 | index .Logs   ).last_name }}

Number of users:
{{ .Logs | len }}

