package decl

import (
	"fmt"

	"github.com/datumbrain/gofp/owlfunctional/meta"
)

type Declaration struct {
	IRI string
}

// func (s *Declaration) PrefixedName() string {
// 	return parser.FmtPrefixedName(s.Prefix, s.Name)
// }

type AnnotationPropertyDecl struct {
	Declaration
}

var _ meta.AnnotationProperty = (*AnnotationPropertyDecl)(nil)

type ClassDecl struct {
	Declaration
}

var _ meta.ClassExpression = (*ClassDecl)(nil)

func (s *ClassDecl) IsNamedClass() bool {
	return true
}

func (s *ClassDecl) String() string {
	return fmt.Sprintf("C{%v}", s.IRI)
}

type DataPropertyDecl struct {
	Declaration
	IsFunctional bool
}

var _ meta.DataProperty = (*DataPropertyDecl)(nil)

func (s *DataPropertyDecl) IsNamedDataProperty() bool {
	return true
}

type DatatypeDecl struct {
	Declaration
}

var _ meta.DataRange = (*DatatypeDecl)(nil)

func (s *DatatypeDecl) IsNamedDatatype() bool {
	return true
}

type NamedIndividualDecl struct {
	Declaration
}

func (s *NamedIndividualDecl) String() string {
	return fmt.Sprintf("NI{%v}", s.IRI)
}

type ObjectPropertyDecl struct {
	Declaration
	IsAsymmetric        bool
	IsFunctional        bool
	IsInverse           bool
	IsInverseFunctional bool
	IsIrreflexive       bool
	IsReflexive         bool
	IsSymmetric         bool
	IsTransitive        bool
}

var _ meta.ObjectPropertyExpression = (*ObjectPropertyDecl)(nil)

func (s *ObjectPropertyDecl) IsNamedObjectProperty() bool {
	return true
}

func (s *ObjectPropertyDecl) String() string {
	return fmt.Sprintf("OP{%v}", s.IRI)
}
