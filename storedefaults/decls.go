package storedefaults

import (
	"fmt"

	"github.com/datumbrain/gofp/owlfunctional/decl"
	"github.com/datumbrain/gofp/owlfunctional/meta"
	"github.com/datumbrain/gofp/store"
)

type DeclStore struct {
	// sets of declarations that were explicitly given:
	// Although declarations are not always required, they can be used to catch obvious errors in ontologies.(https://www.w3.org/2007/OWL/wiki/Syntax#Declaration_Consistency)
	annotationPropertyDecls map[string]*decl.AnnotationPropertyDecl
	classDecls              map[string]*decl.ClassDecl
	dataPropertyDecls       map[string]*decl.DataPropertyDecl
	datatypeDecls           map[string]*decl.DatatypeDecl
	namedIndividualDecls    map[string]*decl.NamedIndividualDecl
	objectPropertyDecls     map[string]*decl.ObjectPropertyDecl

	// the imp* maps are sets of declarations that were implicitly given (by use)
	impAnnotationPropertyDecls map[string]*decl.AnnotationPropertyDecl
	impClassDecls              map[string]*decl.ClassDecl
	impDataPropertyDecls       map[string]*decl.DataPropertyDecl
	impDatatypeDecls           map[string]*decl.DatatypeDecl
	impNamedIndividualDecls    map[string]*decl.NamedIndividualDecl
	impObjectPropertyDecls     map[string]*decl.ObjectPropertyDecl

	// Note: any Declaration can only be in the explicit, or in the implicit sets above, not in both.
	// If an explicit declaration is found by the parser after implicit use, the declaration must me moved
	// into the explicit set.

	// ExplicitDecls = true means, any declaration must be stored explicitly before it can be requested with a Get method.
	// ExplicitDecls = false is the standard OWL behaviour where a declaration is created implicitly when we request it.
	ExplicitDecls bool
}

var _ store.Decls = (*DeclStore)(nil)
var _ store.DeclStore = (*DeclStore)(nil)

func NewDeclStore() *DeclStore {
	return &DeclStore{
		annotationPropertyDecls:    map[string]*decl.AnnotationPropertyDecl{},
		classDecls:                 map[string]*decl.ClassDecl{},
		dataPropertyDecls:          map[string]*decl.DataPropertyDecl{},
		datatypeDecls:              map[string]*decl.DatatypeDecl{},
		namedIndividualDecls:       map[string]*decl.NamedIndividualDecl{},
		objectPropertyDecls:        map[string]*decl.ObjectPropertyDecl{},
		impAnnotationPropertyDecls: map[string]*decl.AnnotationPropertyDecl{},
		impClassDecls:              map[string]*decl.ClassDecl{},
		impDataPropertyDecls:       map[string]*decl.DataPropertyDecl{},
		impDatatypeDecls:           map[string]*decl.DatatypeDecl{},
		impNamedIndividualDecls:    map[string]*decl.NamedIndividualDecl{},
		impObjectPropertyDecls:     map[string]*decl.ObjectPropertyDecl{},
		ExplicitDecls:              true,
	}
}

// === Get - methods that return a single decl by key ========

func (s *DeclStore) AnnotationPropertyDecl(iri string) (decl meta.AnnotationProperty, ok bool) {
	decl, ok = s.annotationPropertyDecls[iri]
	if !ok && !s.ExplicitDecls {
		decl, ok = s.impAnnotationPropertyDecls[iri]
		if !ok {
			s.impAnnotationPropertyDecls[iri] = newAnnotationPropertyDecl(iri)
			decl, ok = s.impAnnotationPropertyDecls[iri]
		}
	}
	return
}

func (s *DeclStore) ClassDecl(iri string) (decl meta.ClassExpression, ok bool) {
	decl, ok = s.classDecls[iri]
	if !ok && !s.ExplicitDecls {
		decl, ok = s.impClassDecls[iri]
		if !ok {
			s.impClassDecls[iri] = newClassDecl(iri)
			decl, ok = s.impClassDecls[iri]
		}
	}

	return
}

func (s *DeclStore) DataPropertyDecl(iri string) (decl meta.DataProperty, ok bool) {
	decl, ok = s.dataPropertyDecls[iri]
	if !ok && !s.ExplicitDecls {
		decl, ok = s.impDataPropertyDecls[iri]
		if !ok {
			s.impDataPropertyDecls[iri] = newDataPropertyDecl(iri)
			decl, ok = s.impDataPropertyDecls[iri]
		}
	}

	return
}

func (s *DeclStore) DatatypeDecl(iri string) (decl meta.DataRange, ok bool) {
	decl, ok = s.datatypeDecls[iri]
	if !ok && !s.ExplicitDecls {
		decl, ok = s.impDatatypeDecls[iri]
		if !ok {
			s.impDatatypeDecls[iri] = newDatatypeDecl(iri)
			decl, ok = s.impDatatypeDecls[iri]
		}
	}

	return
}

func (s *DeclStore) NamedIndividualDecl(iri string) (decl interface{}, ok bool) {
	decl, ok = s.namedIndividualDecls[iri]
	if !ok && !s.ExplicitDecls {
		decl, ok = s.impNamedIndividualDecls[iri]
		if !ok {
			s.impNamedIndividualDecls[iri] = newNamedIndividualDecl(iri)
			decl, ok = s.impNamedIndividualDecls[iri]
		}
	}

	return
}

func (s *DeclStore) ObjectPropertyDecl(iri string) (decl meta.ObjectPropertyExpression, ok bool) {
	decl, ok = s.objectPropertyDecls[iri]
	if !ok && !s.ExplicitDecls {
		decl, ok = s.impObjectPropertyDecls[iri]
		if !ok {
			s.impObjectPropertyDecls[iri] = newObjectPropertyDecl(iri)
			decl, ok = s.impObjectPropertyDecls[iri]
		}
	}

	return
}

// === End Get - methods ========

// === All* - methods that return slices ========
func (s *DeclStore) AllAnnotationPropertyDecls() []*decl.AnnotationPropertyDecl {
	res := make([]*decl.AnnotationPropertyDecl, 0, len(s.annotationPropertyDecls))
	for _, v := range s.annotationPropertyDecls {
		res = append(res, v)
	}
	for _, v := range s.impAnnotationPropertyDecls {
		res = append(res, v)
	}
	return res
}
func (s *DeclStore) AllClassDecls() []*decl.ClassDecl {
	res := make([]*decl.ClassDecl, 0, len(s.classDecls))
	for _, v := range s.classDecls {
		res = append(res, v)
	}
	for _, v := range s.impClassDecls {
		res = append(res, v)
	}
	return res
}

func (s *DeclStore) AllDataPropertyDecls() []*decl.DataPropertyDecl {
	res := make([]*decl.DataPropertyDecl, 0, len(s.dataPropertyDecls))
	for _, v := range s.dataPropertyDecls {
		res = append(res, v)
	}
	for _, v := range s.impDataPropertyDecls {
		res = append(res, v)
	}
	return res
}

func (s *DeclStore) AllDatatypeDecls() []*decl.DatatypeDecl {
	res := make([]*decl.DatatypeDecl, 0, len(s.datatypeDecls))
	for _, v := range s.datatypeDecls {
		res = append(res, v)
	}
	for _, v := range s.impDatatypeDecls {
		res = append(res, v)
	}
	return res
}

func (s *DeclStore) AllNamedIndividualDecls() []*decl.NamedIndividualDecl {
	res := make([]*decl.NamedIndividualDecl, 0, len(s.namedIndividualDecls))
	for _, v := range s.namedIndividualDecls {
		res = append(res, v)
	}
	for _, v := range s.impNamedIndividualDecls {
		res = append(res, v)
	}
	return res
}

func (s *DeclStore) AllObjectPropertyDecls() []*decl.ObjectPropertyDecl {
	res := make([]*decl.ObjectPropertyDecl, 0, len(s.objectPropertyDecls))
	for _, v := range s.objectPropertyDecls {
		res = append(res, v)
	}
	for _, v := range s.impObjectPropertyDecls {
		res = append(res, v)
	}
	return res
}

// === end All - methods =======

// === Store - methods to store explicit declarations =======

func errDoubleExplicitDecl(iri string) error {
	return fmt.Errorf("Repeated explicit declaration of %s", iri)
}

func (s *DeclStore) StoreAnnotationPropertyDecl(iri string) (err error) {
	if _, ok := s.annotationPropertyDecls[iri]; ok {
		err = errDoubleExplicitDecl(iri)
	} else {
		s.annotationPropertyDecls[iri] = newAnnotationPropertyDecl(iri)
	}
	return
}

func (s *DeclStore) StoreClassDecl(iri string) (err error) {
	if _, ok := s.classDecls[iri]; ok {
		err = errDoubleExplicitDecl(iri)
	} else {
		s.classDecls[iri] = newClassDecl(iri)
	}
	return
}

func (s *DeclStore) StoreDataPropertyDecl(iri string) (err error) {
	if _, ok := s.dataPropertyDecls[iri]; ok {
		err = errDoubleExplicitDecl(iri)
	} else {
		s.dataPropertyDecls[iri] = newDataPropertyDecl(iri)
	}
	return
}

func (s *DeclStore) StoreDatatypeDecl(iri string) (err error) {
	if _, ok := s.datatypeDecls[iri]; ok {
		err = errDoubleExplicitDecl(iri)
	} else {
		s.datatypeDecls[iri] = newDatatypeDecl(iri)
	}
	return
}

func (s *DeclStore) StoreNamedIndividualDecl(iri string) (err error) {
	if _, ok := s.namedIndividualDecls[iri]; ok {
		err = errDoubleExplicitDecl(iri)
	} else {
		s.namedIndividualDecls[iri] = newNamedIndividualDecl(iri)
	}
	return
}

func (s *DeclStore) StoreObjectPropertyDecl(iri string) (err error) {
	if _, ok := s.objectPropertyDecls[iri]; ok {
		err = errDoubleExplicitDecl(iri)
	} else {
		s.objectPropertyDecls[iri] = newObjectPropertyDecl(iri)
	}
	return
}

// === end Store - methods =======

func newAnnotationPropertyDecl(iri string) *decl.AnnotationPropertyDecl {
	return &decl.AnnotationPropertyDecl{Declaration: decl.Declaration{IRI: iri}}
}

func newClassDecl(iri string) *decl.ClassDecl {
	return &decl.ClassDecl{Declaration: decl.Declaration{IRI: iri}}
}

func newDataPropertyDecl(iri string) *decl.DataPropertyDecl {
	return &decl.DataPropertyDecl{Declaration: decl.Declaration{IRI: iri}}
}

func newDatatypeDecl(iri string) *decl.DatatypeDecl {
	return &decl.DatatypeDecl{Declaration: decl.Declaration{IRI: iri}}
}

func newNamedIndividualDecl(iri string) *decl.NamedIndividualDecl {
	return &decl.NamedIndividualDecl{Declaration: decl.Declaration{IRI: iri}}
}

func newObjectPropertyDecl(iri string) *decl.ObjectPropertyDecl {
	return &decl.ObjectPropertyDecl{Declaration: decl.Declaration{IRI: iri}}
}

// === Exists - methods, intended for tests only =======

func (s *DeclStore) ClassDeclExists(iri string, imp bool) bool {
	_, ok := s.classDecls[iri]
	if !ok && imp {
		_, ok = s.impClassDecls[iri]
	}
	return ok
}

func (s *DeclStore) DataPropertyDeclExists(iri string, imp bool) bool {
	_, ok := s.dataPropertyDecls[iri]
	if !ok && imp {
		_, ok = s.impDataPropertyDecls[iri]
	}
	return ok
}

func (s *DeclStore) NamedIndividualDeclExists(iri string, imp bool) bool {
	_, ok := s.namedIndividualDecls[iri]
	if !ok && imp {
		_, ok = s.impNamedIndividualDecls[iri]
	}
	return ok
}

func (s *DeclStore) ObjectPropertyDeclExists(iri string, imp bool) bool {
	_, ok := s.objectPropertyDecls[iri]
	if !ok && imp {
		_, ok = s.impObjectPropertyDecls[iri]
	}
	return ok
}

// === End Exists - methods =======

func (s *DeclStore) String() string {
	return fmt.Sprintf("%d annotations, %d classes, %d object properties, %d data properties, %d named individuals, %d datatypes",
		len(s.annotationPropertyDecls),
		len(s.classDecls),
		len(s.objectPropertyDecls),
		len(s.dataPropertyDecls),
		len(s.namedIndividualDecls),
		len(s.datatypeDecls),
	)
}

type DefaultDecl struct {
	annotations []meta.Annotation
}

func (s *DefaultDecl) Annotations() []meta.Annotation {
	return s.annotations
}
