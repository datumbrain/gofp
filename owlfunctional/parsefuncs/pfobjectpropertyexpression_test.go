package parsefuncs

import (
	"testing"

	"github.com/datumbrain/gofp/mock"
	"github.com/datumbrain/gofp/owlfunctional/meta"
	"github.com/datumbrain/gofp/owlfunctional/parser"
	"github.com/datumbrain/gofp/owlfunctional/properties"
)

func TestParseObjectPropertyExpression(t *testing.T) {
	var p *parser.Parser
	var err error
	decls, prefixes := mock.NewBuilder().AddPrefixes("").AddOWLStandardPrefixes().Get()

	p = mock.NewTestParser(`owl:topObjectProperty`)
	var expr meta.ObjectPropertyExpression
	expr, err = ParseObjectPropertyExpression(p, decls, prefixes)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := expr.(*properties.OWLTopObjectProperty); !ok {
		t.Fatal(expr)
	}
}
