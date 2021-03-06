// tech are types not required by OWL but internally.
package tech

import (
	"fmt"
	"strings"
)

type Prefixes interface {
	// ResolvePrefix returns the IRI part which is associated with the prefix
	// false if prefix was unknown.
	ResolvePrefix(prefix string) (resolved string, ok bool)
}

//todo: Eventually remove the IRI type, it comes from my early wrong idea that IRI fragments play a special role (RR)
// Simply a string may work.

// IRI is an identifier used by all OWL declarations.
// In case the IRI has a fragment, it is stored in two pieces - the fragment (without hash sign), and everything before.
// The intention is that the part before the fragment, in some cases, has a meaning for itself,
// e.g. to answer the question "is this an element of the OWL namespace". While not beeing of great value, storing these two pices sometimes saves some Strings.startsWith - operations.
type IRI struct {
	// Fragment is empty for an IRI without fragment.
	// Example: for http://www.w3.org/2002/07/owl#Thing, the Fragment ist "Thing" - without the separating Hash.
	Fragment string // e.g."Thing"

	// Head + Fragment forms the whole IRI String.
	// In case there's no Fragment, Head is the whole IRI.
	// In case there is a fragment, Head MUST end with Hash (#).
	Head string // e.g."http://www.w3.org/2002/07/owl#"
}

// MustNewFragmentedIRI expects a head ending with "#".
// Panics otherwise.
func MustNewFragmentedIRI(head, fragment string) *IRI {
	if !strings.HasSuffix(head, "#") {
		panic(fmt.Sprintf("fragmented IRI needs head with Suffix '#'. (Got head=%v and fragment=%v)", head, fragment))
	}
	return &IRI{Head: head, Fragment: fragment}
}

// NewIRIFromString separates the fragment from the first part (Head), if the given value has a fragment.
// Otherwise, Fragment remains empty.
// error if val is no valid IRI. Note that most error conditions are not checked.
func NewIRIFromString(val string) (*IRI, error) {
	parts := strings.Split(val, "#")
	switch len(parts) {
	case 1: // no "#"
		return &IRI{Head: parts[0]}, nil
	case 2: // had "#" : keep the # at the end of Head
		return &IRI{Head: parts[0] + "#", Fragment: parts[1]}, nil
	default:
		return nil, fmt.Errorf("invalid IRI string with multiple # (%v)", val)
	}
}

func (s *IRI) String() string {
	return s.Head + s.Fragment
}

// ZeroBasedPosWord is "first" for 0, then "second" ... 4th ... and so on
func ZeroBasedPosWord(i int) string {
	switch i {
	case 0:
		return "first"
	case 1:
		return "second"
	case 2:
		return "third"
	default:
		return fmt.Sprintf("%dth", i)
	}
}
