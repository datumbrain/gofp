package parsefuncs

import (
	"reifenberg.de/gofp/owlfunctional/meta"
	"reifenberg.de/gofp/owlfunctional/parser"
	"reifenberg.de/gofp/owlfunctional/properties"
	"reifenberg.de/gofp/parsehelper"
	"reifenberg.de/gofp/tech"
)

func ParseDataProperty(p *parser.Parser, decls tech.Declarations, prefixes tech.Prefixes) (expr meta.DataProperty, err error) {

	pos := p.Pos()
	// must be R
	var ident *tech.IRI
	ident, err = parsehelper.ParseAndResolveIRI(p, prefixes)

	if err != nil {
		return
	}

	if ident.IsOWL() {
		// must be one of the predefined OWL property names
		switch ident.Name {
		case "topDataProperty":
			expr = &properties.OWLTopDataProperty{}
		case "bottomDataProperty":
			expr = &properties.OWLBottomDataProperty{}
		default:
			err = pos.Errorf(`unexpected OWL property "%v"`, ident.Name)
		}
		return
	}

	var ok bool
	expr, ok = decls.GetDataPropertyDecl(*ident)
	if !ok {
		err = pos.Errorf("Unknown ref to %v", ident)
	}
	return
}
