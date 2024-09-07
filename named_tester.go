package tester

import "testing"

type NamedTester interface {
	Name() string
	Test(t *testing.T)
}

type NamedTesterFunc struct {
	name string
	f    func(t *testing.T)
}

func (nt *NamedTesterFunc) Name() string {
	return nt.name
}

func (nt *NamedTesterFunc) Test(t *testing.T) {
	nt.f(t)
}

func NamedTester(name string, f func(t *testing.T)) NamedTester {
	return &NamedTesterFunc{
		name: name,
		f:    f,
	}
}

func RunNamedTesters(t *testing.T, namedTesters ...NamedTester) {
	for _, tester := range namedTesters {
		if tester == nil {
			continue
		}

		t.Run(tester.Name(), tester.Test)
	}
}
