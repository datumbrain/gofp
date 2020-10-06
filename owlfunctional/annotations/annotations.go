package annotations

import (
	"github.com/datumbrain/gofp/owlfunctional/meta"
)

type Annotation struct {
	a meta.AnnotationProperty
	t string
}

var _ meta.Annotation = (*Annotation)(nil)

func NewAnnotation(a meta.AnnotationProperty, t string) *Annotation {
	return &Annotation{
		a: a,
		t: t,
	}
}

func (s *Annotation) A() meta.AnnotationProperty {
	return s.a
}

func (s *Annotation) T() string {
	return s.t
}

type AnnotationAssertion struct {
	A meta.AnnotationProperty

	// Literaltype can be one of the builtin datatypes like xsd:integer, or a custom IRI.
	S string

	// langTag can be set on strings. The literal "foo"@en results in langTag="en".
	// If not given, langTag is empty. For non-string types, it is also empty.
	T string
}

type AnnotationPropertyDomain struct {
	A meta.AnnotationProperty
	U string
}

type AnnotationPropertyRange struct {
	A meta.AnnotationProperty
	U string
}

type SubAnnotationPropertyOf struct {
	A1 meta.AnnotationProperty
	A2 meta.AnnotationProperty
}
