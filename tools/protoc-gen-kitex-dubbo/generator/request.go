package generator

import "google.golang.org/protobuf/compiler/protogen"

type Field struct {
	Name string
}

type ExtStruct struct {
	Name        string
	Fields      []*Field
	JavaPackage string
	GoPackage   string
}

type ExtService struct {
	Name        string
	PkgName     string
	Functions   []*ExtFunction
	Annotations map[string]string
	Version     string
}

type ExtFunction struct {
	Method     string
	InputType  string
	OutputType string
}

type ExtFile struct {
	JavaPackage string
	GoPackage   string
	PkgName     protogen.GoPackageName
	IDLName     string
	StructLikes []*ExtStruct
	Services    []*ExtService
	Version     string
}
