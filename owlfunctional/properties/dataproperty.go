package properties

import (
	"github.com/datumbrain/gofp/owlfunctional/meta"
)

type OWLBottomDataProperty struct {
}

var _ meta.DataProperty = (*OWLBottomDataProperty)(nil)

func (*OWLBottomDataProperty) IsNamedDataProperty() bool {
	return false
}

type OWLTopDataProperty struct {
}

var _ meta.DataProperty = (*OWLTopDataProperty)(nil)

func (*OWLTopDataProperty) IsNamedDataProperty() bool {
	return false
}
