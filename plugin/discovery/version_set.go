package discovery

import (
	version "github.com/hashicorp/go-version"
)

// A ConstraintsStr is a string containing a possibly-invalid representation
// of a version constraint provided in configuration. Call Parse on it to
// obtain a real Constraint object, or discover that it is invalid.
type ConstraintsStr string

// Parse transforms a ConstraintsStr into a VersionSet if it is
// syntactically valid. If it isn't then an error is returned instead.
func (s ConstraintsStr) Parse() (VersionSet, error) {
	raw, err := version.NewConstraint(string(s))
	if err != nil {
		return VersionSet{}, err
	}
	return VersionSet{raw}, nil
}

// MustParse is like Parse but it panics if the constraint string is invalid.
func (s ConstraintsStr) MustParse() VersionSet {
	ret, err := s.Parse()
	if err != nil {
		panic(err)
	}
	return ret
}

// VersionSet represents a set of versions which any given Version is either
// a member of or not.
type VersionSet struct {
	raw version.Constraints
}

// Has returns true if the given version is in the receiving set.
func (s VersionSet) Has(v Version) bool {
	return s.raw.Check(v.raw)
}

// Intersection combines the receving set with the given other set to produce a
// set that is the intersection of both sets, which is to say that it contains
// only the versions that are members of both sets.
func (s VersionSet) Intersection(other VersionSet) VersionSet {
	raw := make(version.Constraints, 0, len(s.raw)+len(other.raw))

	raw = append(raw, s.raw...)
	raw = append(raw, other.raw...)

	return VersionSet{raw}
}

// String returns a string representation of the set members as a set
// of range constraints.
func (c VersionSet) String() string {
	return c.raw.String()
}
