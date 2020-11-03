public class User implements java.io.Serializable {
    {{range .User.Property }}
        private {{.Type}} {{ .Name  }};
    {{end}}


    {{range .User.Property }}
        public {{.Type}} {{ .Name | ToGetterName }}(){
            return this.{{.Name}};
        }
    {{end}}

    {{range .User.Property }}
        public void {{ .Name   | ToSetterName   }}({{.Type}} {{.Name}}){
            this.{{.Name}} = {{.Name}};
        }
    {{end}}
}