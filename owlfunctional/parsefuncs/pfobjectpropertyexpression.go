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

func ParseObjectPropertyExpression(p *parser.Parser, decls store.Decls, prefixes tech.Prefixes) (expr meta.ObjectPropertyExpression, err error) {
	tok, lit, pos := p.ScanIgnoreWSAndComment()
	switch tok {
	case parser.ObjectInverseOf:
		p.Unscan()
		expr, err = parseObjectInverseOf(p, prefixes)
	default:
		p.Unscan()
		// must be PN
		var ident *tech.IRI
		ident, err = parsehelper.ParseAndResolveIRI(p, prefixes)
		if err != nil {
			err = pos.ErrorfUnexpectedToken(tok, lit, "Object Property Name")
			return
		}

		if builtindatatypes.IsOWL(*ident) {
			// must be one of the predefined OWL property names
			switch ident.Fragment {
			case "topObjectProperty":
				expr = &properties.OWLTopObjectProperty{}
			case "bottomObjectProperty":
				expr = &properties.OWLBottomObjectProperty{}
			default:
				err = pos.Errorf(`unexpected OWL property "%v"`, ident.Fragment)
			}
			return
		}
		var ok bool
		expr, ok = decls.ObjectPropertyDecl(ident.String())
		if !ok {
			err = pos.Errorf("Unknown ref to %v. Expected object property name.", ident)
		}
	}
	return
}

func parseObjectInverseOf(p *parser.Parser, prefixes tech.Prefixes) (expr meta.ObjectPropertyExpression, err error) {
	if err = p.ConsumeTokens(parser.ObjectInverseOf, parser.B1); err != nil {
		return
	}
	pos := p.Pos()
	var ident *tech.IRI
	ident, err = parsehelper.ParseAndResolveIRI(p, prefixes)
	if err != nil {
		err = pos.Errorf("parsing IRI in ObjectInverseOf: %v", err)
		return
	}
	if err = p.ConsumeTokens(parser.B2); err != nil {
		return
	}
	expr = &properties.ObjectInverseOf{PN: ident.String()}
	return
}
