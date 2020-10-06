package assertions

import (
	"github.com/datumbrain/gofp/owlfunctional/individual"
	"github.com/datumbrain/gofp/owlfunctional/meta"
)

type NegativeObjectPropertyAssertion struct {
	P  meta.ObjectPropertyExpression
	A1 individual.Individual
	A2 individual.Individual
}

type ObjectPropertyAssertion struct {
	PN string
	A1 individual.Individual
	A2 individual.Individual
}
