package facets

import "github.com/datumbrain/gofp/owlfunctional/literal"

type Facet int

const (
	Xsd_minInclusive Facet = iota
	Xsd_maxInclusive
	Xsd_minExclusive
	Xsd_maxExclusive
	//todo more Facets
)

type FVPair struct {
	F Facet
	V literal.OWLLiteral
}
