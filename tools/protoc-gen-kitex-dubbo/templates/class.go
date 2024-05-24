package templates

var Class = `
import java.io.Serializable;

public class {{.ClassName}} implements Serializable {
{{- range .Fields}}
	{{- InsertionPoint $.Category $.Name .Name}}
	{{- if and Features.ReserveComments .ReservedComments}}
	{{.ReservedComments}}
	{{- end}}
	{{(.GoName)}} {{.GoTypeName}} {{GenFieldTags . (InsertionPoint $.Category $.Name .Name "tag")}} 
{{- end}}{{/* range .Fields */}}

	{{- if Features.KeepUnknownFields}}
	{{- UseStdLibrary "unknown"}}
	_unknownFields unknown.Fields
	{{- end}} {{/*- if Features.KeepUnknownFields*/}}


    String req;

    public GreetRequest(String req) {
        this.req = req;
    }

    public String getReq() {
        return req;
    }
}
`
