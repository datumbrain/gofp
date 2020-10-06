package store

// store has interfaces for container types, used by the parser to hold Declarations and Axioms.
// See the storedefaults package for the reference implementation.

import (
	"github.com/datumbrain/gofp/owlfunctional/individual"
	"github.com/datumbrain/gofp/owlfunctional/literal"
	"github.com/datumbrain/gofp/owlfunctional/meta"
)

// Decls gives read access to all Declarations that were parsed yet.
type Decls interface {
	// By Key - methods:
	AnnotationPropertyDecl(ident string) (meta.AnnotationProperty, bool)
	ClassDecl(ident string) (meta.ClassExpression, bool)
	DataPropertyDecl(ident string) (meta.DataProperty, bool)
	DatatypeDecl(ident string) (meta.DataRange, bool)
	NamedIndividualDecl(ident string) (interface{}, bool)
	ObjectPropertyDecl(ident string) (meta.ObjectPropertyExpression, bool)
}

// DeclStore is used by the parser to store explicit declarations.
// The store functions should return error if the declaration was already explicitly given. The "should" wording is because a custom implementation
// of DeclStore may choose to silently ignore double declarations.
// No error is returned if the declaration is already known, but was implicitly declared only.
type DeclStore interface {
	StoreAnnotationPropertyDecl(iri string) error
	StoreClassDecl(iri string) error
	StoreDataPropertyDecl(iri string) error
	StoreDatatypeDecl(iri string) error
	StoreNamedIndividualDecl(iri string) error
	StoreObjectPropertyDecl(iri string) error
}

// AxiomStore takes all possible axioms and encapsulates the data structures to store them.
type AxiomStore interface {
	StoreAnnotationAssertion(A meta.AnnotationProperty, S string, t string, anns []meta.Annotation)
	StoreAnnotationPropertyDomain(A meta.AnnotationProperty, U string, anns []meta.Annotation)
	StoreAnnotationPropertyRange(A meta.AnnotationProperty, U string, anns []meta.Annotation)
	StoreAsymmetricObjectProperty(P meta.ObjectPropertyExpression, anns []meta.Annotation)
	StoreClassAssertion(C meta.ClassExpression, a individual.Individual, anns []meta.Annotation)
	StoreDataPropertyAssertion(R meta.DataProperty, a individual.Individual, v literal.OWLLiteral, anns []meta.Annotation)
	StoreFunctionalDataProperty(R meta.DataProperty, anns []meta.Annotation)
	StoreFunctionalObjectProperty(P meta.ObjectPropertyExpression, anns []meta.Annotation)
	StoreInverseFunctionalObjectProperty(P meta.ObjectPropertyExpression, anns []meta.Annotation)
	StoreInverseObjectProperties(P1, P2 meta.ObjectPropertyExpression, anns []meta.Annotation)
	StoreIrreflexiveObjectProperty(P meta.ObjectPropertyExpression, anns []meta.Annotation)
	StoreDataPropertyDomain(R meta.DataProperty, C meta.ClassExpression, anns []meta.Annotation)
	StoreDataPropertyRange(R meta.DataProperty, D meta.DataRange, anns []meta.Annotation)
	StoreDisjointClasses(Cs []meta.ClassExpression, anns []meta.Annotation)
	StoreDifferentIndividuals(as []individual.Individual, anns []meta.Annotation)
	StoreEquivalentClasses(Cs []meta.ClassExpression, anns []meta.Annotation)
	StoreNegativeObjectPropertyAssertion(P meta.ObjectPropertyExpression, a1 individual.Individual, a2 individual.Individual)
	StoreObjectPropertyAssertion(PN string, a1 individual.Individual, a2 individual.Individual)
	StoreObjectPropertyDomain(P meta.ObjectPropertyExpression, C meta.ClassExpression, anns []meta.Annotation)
	StoreObjectPropertyRange(P meta.ObjectPropertyExpression, C meta.ClassExpression, anns []meta.Annotation)
	StoreReflexiveObjectProperty(P meta.ObjectPropertyExpression, anns []meta.Annotation)
	StoreSubAnnotationPropertyOf(A1, A2 string, anns []meta.Annotation)
	StoreSubClassOf(Csub, Csuper meta.ClassExpression, anns []meta.Annotation)
	StoreSubDataPropertyOf(P1, P2 meta.DataProperty, anns []meta.Annotation)
	StoreSubObjectPropertyOf(P1, P2 meta.ObjectPropertyExpression, anns []meta.Annotation)
	StoreSymmetricObjectProperty(P meta.ObjectPropertyExpression, anns []meta.Annotation)
	StoreTransitiveObjectProperty(P meta.ObjectPropertyExpression, anns []meta.Annotation)
}
