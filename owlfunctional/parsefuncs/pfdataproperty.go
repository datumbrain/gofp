package parsefuncs

import (
	"github.com/datumbrain/gofp/owlfunctional/builtindatatypes"
	"github.com/datumbrain/gofp/owlfunctional/meta"
	"github.com/datumbrain/gofp/owlfunctional/parser"
	"github.com/datumbrain/gofp/owlfunctional/properties"
	"github.com/datumbrain/gofp/parsehelper"
	"github.com/datumbrain/gofp/store"
	"github.com/datumbrain/gofp/tech"
)

func ParseDataProperty(p *parser.Parser, decls store.Decls, prefixes tech.Prefixes) (expr meta.DataProperty, err error) {

	pos := p.Pos()
	// must be R
	var ident *tech.IRI
	ident, err = parsehelper.ParseAndResolveIRI(p, prefixes)

	if err != nil {
		return
	}

	if builtindatatypes.IsOWL(*ident) {
		// must be one of the predefined OWL property names
		switch ident.Fragment {
		case "topDataProperty":
			expr = &properties.OWLTopDataProperty{}
		case "bottomDataProperty":
			expr = &properties.OWLBottomDataProperty{}
		default:
			err = pos.Errorf(`unexpected OWL property "%v"`, ident.Fragment)
		}
		return
	}

	var ok bool
	expr, ok = decls.DataPropertyDecl(ident.String())
	if !ok {
		err = pos.Errorf("Unknown ref to %v. Expected datatype property.", ident)
	}
	return
}
