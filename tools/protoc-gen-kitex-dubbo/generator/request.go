package generator

import "google.golang.org/protobuf/compiler/protogen"

type Field struct {
	Name         string
	Type         string
	FunctionName string
}

type StructLike struct {
	Name        string
	Fields      []*Field
	JavaPackage string
	GoPackage   string
}

type ExtService struct {
	Name        string
	PkgName     string
	Functions   []*Function
	Annotations map[string]string
	Version     string
}

type Function struct {
	Method     string
	InputType  string
	OutputType string
}

type ExtFile struct {
	JavaPackage string
	GoPackage   string
	PkgName     protogen.GoPackageName
	IDLName     string
	StructLikes []*StructLike
	Services    []*ExtService
	Version     string
}
