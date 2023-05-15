package hessian2

import (
	"fmt"
	"strings"
	"time"
)

type JavaBean interface {
	JavaClassPackage() string

	JavaClassName() string
}

type VirtualClass struct {
	className    string
	classPackage string
	fields       map[string]interface{}
}

func NewVirtualClass(typ string, fieldNames []string) *VirtualClass {
	idx := strings.LastIndexByte(typ, '.')
	classPackage := ""
	className := typ
	if idx > 0 {
		classPackage = typ[:idx]
		className = typ[idx:]
	}
	fields := make(map[string]interface{})
	for _, key := range fieldNames {
		fields[key] = nil
	}
	return &VirtualClass{
		classPackage: classPackage,
		className:    className,
		fields:       fields,
	}
}

func (vc *VirtualClass) JavaClassPackage() string {
	return vc.classPackage
}

func (vc *VirtualClass) JavaClassName() string {
	return vc.className
}

func (vc *VirtualClass) JavaFields() map[string]interface{} {
	return vc.fields
}

func convJavaType(v interface{}) (t string, r interface{}) {
	t = "Object"
	r = fmt.Sprintf("%v", v)
	switch v.(type) {
	case string:
		t = "String"
		r = fmt.Sprintf(`"%s"`, v)
	case int, int32, uint, uint32:
		t = "Integer"
	case int8, uint8:
		t = "Byte"
	case int16, uint16:
		t = "Short"
	case int64, uint64:
		t = "Long"
		r = fmt.Sprintf("%vL", v)
	case float32:
		t = "Float"
		r = fmt.Sprintf("%vF", v)
	case float64:
		t = "Double"
	case bool:
		t = "Bool"
	case time.Time:
		t = "Date"
		r = fmt.Sprintf(`new Date("%s")`, v)
	}
	return
}

func (vc *VirtualClass) String() string {
	classMsg := "\nclass " + vc.className + " {"
	if vc.classPackage != "" {
		classMsg = "\npackage " + vc.classPackage + "" + classMsg
	}
	for key, value := range vc.JavaFields() {
		t, r := convJavaType(value)
		classMsg += fmt.Sprintf("\n\tprivate %s %s = %s;", t, key, r)
	}
	classMsg += "\n}"
	return classMsg
}
