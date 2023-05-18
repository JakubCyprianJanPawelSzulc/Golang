package graph

import "testing"

type Flags map[string]bool // flags related to boolean options: ON/OFF state

// create new Flags object
// like this: someflags := Graph.Flags{}.New()
func (f Flags) New() Flags {
	return make(Flags)
}

// check if any flag is set inside
// if someflags.IsSet("thisflag") { ... }
func (F Flags) IsSet(flag string) bool {
	f, ok := F[flag]
	return ok && f
}

// unconditionally SET the flag to true
func (F Flags) Set(flag string) {
	F[flag] = true
}

// negate any flag, if it was there
// someflag.Neg("there")   - negate "there" flag, if was present
// if flag was not set there, it is set to False
func (F Flags) Neg(flag string) {
	if fl, ok := F[flag]; ok {
		if fl {
			F[flag] = false
		} else {
			F[flag] = true
		}
	} else {
		F[flag] = false
	}
}

// testing
func TestNewEmpty(t *testing.T) {
	F := Flags{}.New()
	if F == nil {
		t.Fatalf("Flags is created with New but returned nil!")
	}
}

func TestFlagsSet(t *testing.T) {
	F := Flags{}.New()
	F["test"] = true
	if !F.IsSet("test") {
		t.Fatalf("Flags has 'test' flag set, but it IsSet claims it is false.")
	}
}
