package hessian2

import "strings"

type TypeAnnotation struct {
	anno  string
	types []string
}

func NewTypeAnnotation(anno string) *TypeAnnotation {
	return &TypeAnnotation{anno: anno}
}

func (ta *TypeAnnotation) GetType(i int) string {
	if len(ta.getTypes()) > i {
		return ta.getTypes()[i]
	}
	return ""
}

func (ta *TypeAnnotation) getTypes() []string {
	if ta.types == nil {
		if ta.anno[:5] == "type:" {
			ta.anno = ta.anno[5:]
		}
		ta.types = strings.Split(ta.anno, ",")
	}
	return ta.types
}

type MethodAnnotation struct {
	reqAnnos  *TypeAnnotation
	respAnnos *TypeAnnotation
}

func NewMethodAnnotation(annos map[string][]string) *MethodAnnotation {
	ma := new(MethodAnnotation)
	if v, ok := annos[HESSIAN_REQUEST_TAG]; ok && len(v) > 0 {
		ma.reqAnnos = NewTypeAnnotation(v[0])
	}
	if v, ok := annos[HESSIAN_RESPONSE_TAG]; ok && len(v) > 0 {
		ma.respAnnos = NewTypeAnnotation(v[0])
	}
	return ma
}

func (ma *MethodAnnotation) GetRequestTypeAnnos() *TypeAnnotation {
	return ma.reqAnnos
}

func (ma *MethodAnnotation) GetResponseTypeAnnos() *TypeAnnotation {
	return ma.respAnnos
}
