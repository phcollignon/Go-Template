
Contact Card:
  - Name : {{ .Name }}, {{ .Firstname }}
  - Address : 
        {{ .Address.Street }}
        {{ .Address.Postcode }}, {{ .Address.City}}
  - Contact
        - Phones : 
            - {{ .Phones.Private }} (Private)
            - {{ .Phones.Work }} (Work)
            - {{ .Phones.Mobile }} (Mobile)
