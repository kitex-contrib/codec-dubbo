package generator

import (
	"html/template"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"

	"github.com/kitex-contrib/codec-dubbo/tools/protoc-gen-kitex-dubbo/templates"
)

const (
	filePrefix = "kitex_gen/"
	fileExt    = "hessian2_ext.go"
	version    = "v0.0.1"
)

type Generator interface {
	Generate(gen *protogen.Plugin) error
}

type generator struct {
	extLang string
	tpl     *template.Template
}

func (g *generator) Generate(gen *protogen.Plugin) error {
	for _, file := range gen.Files {
		err := g.generateAPIFile(gen, file)
		if err != nil {
			return err
		}

		for _, svr := range file.Services {
			err := g.generateServiceExt(gen, svr, file)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// generateAPIFile implements dubbo extension interface every message.
// ref: https://github.com/kitex-contrib/codec-dubbo/blob/dcc4d83669b139d5daa75823d2eda5584d92ebc7/pkg/iface/protocol.go#L22
func (g *generator) generateAPIFile(gen *protogen.Plugin, file *protogen.File) error {
	if len(file.Messages) == 0 {
		return nil
	}
	fileName := filePrefix + file.GeneratedFilenamePrefix + "_" + fileExt
	genFile := gen.NewGeneratedFile(fileName, "")
	request := &ExtFile{
		Version: version,
		PkgName: file.GoPackageName,
	}
	if file.Proto != nil && file.Proto.Options != nil {
		if file.Proto.Options.GoPackage != nil {
			request.GoPackage = *file.Proto.Options.GoPackage
		}
		if file.Proto.Options.JavaPackage != nil {
			request.JavaPackage = *file.Proto.Options.JavaPackage
		}
	}
	messages := file.Messages
	for len(messages) != 0 {
		var nestedMessages []*protogen.Message
		for _, message := range messages {
			s := &StructLike{
				Name:        message.GoIdent.GoName,
				GoPackage:   request.GoPackage,
				JavaPackage: request.JavaPackage,
			}
			for _, field := range message.Fields {
				filed := &Field{
					Name: field.GoName,
				}
				s.Fields = append(s.Fields, filed)
			}
			request.StructLikes = append(request.StructLikes, s)
			nestedMessages = append(nestedMessages, message.Messages...)
		}
		messages = nestedMessages
	}

	for _, svr := range file.Services {
		s := &ExtService{
			Name: svr.GoName,
		}
		request.Services = append(request.Services, s)
	}

	var buf strings.Builder
	err := g.tpl.ExecuteTemplate(&buf, "structlikes", request)
	if err != nil {
		return err
	}
	genFile.P(buf.String())
	return nil
}

// generateServiceExt: implements dubbo extension interface for service method args and response
// ref: https://github.com/cloudwego/kitex/blob/8526b3af30fcd321db268cae59a3545a9c6f237f/tool/internal_pkg/generator/generator.go#L392
func (g *generator) generateServiceExt(gen *protogen.Plugin, svr *protogen.Service, file *protogen.File) error {
	fileName := filePrefix + string(file.GoImportPath) + "/" + strings.ToLower(svr.GoName) + "/" + fileExt
	genFile := gen.NewGeneratedFile(fileName, "")
	rSvr := &ExtService{
		Version: version,
		PkgName: strings.ToLower(svr.GoName),
		Name:    svr.GoName,
	}
	for _, method := range svr.Methods {
		s := &Function{
			Method:     method.GoName,
			InputType:  method.Input.GoIdent.GoName,
			OutputType: method.Output.GoIdent.GoName,
		}
		rSvr.Functions = append(rSvr.Functions, s)
	}
	var buf strings.Builder
	err := g.tpl.ExecuteTemplate(&buf, "service", rSvr)
	if err != nil {
		return err
	}
	genFile.P(buf.String())
	return nil
}

// New implements Generator
func New(req *protogen.Plugin, lang string) (Generator, error) {
	tpl := template.New("kitex-dubbo")
	allTemplates := []string{
		templates.StructLikes,
		templates.Header,
		templates.StructLike,
		templates.Service,
		templates.JavaClassName,
	}
	var err error
	for _, temp := range allTemplates {
		tpl, err = tpl.Parse(temp)
		if err != nil {
			return nil, err
		}
	}
	return &generator{
		tpl:     tpl,
		extLang: lang,
	}, nil
}
