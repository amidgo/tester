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

func NewNamedTester(name string, f func(t *testing.T)) NamedTester {
	return &NamedTesterFunc{
		name: name,
		f:    f,
	}
}

type NamedTesterContainer struct {
	namedTesters []NamedTester
}

func RunNamedTesters(t *testing.T, namedTesters ...NamedTester) {
	for _, tester := range namedTesters {
		if tester == nil {
			continue
		}

		t.Run(tester.Name(), tester.Test)
	}
}

func (c *NamedTesterContainer) AddNamedTester(tester NamedTester) {
	c.namedTesters = append(c.namedTesters, tester)
}

func (c *NamedTesterContainer) AddNamedTesters(testers ...NamedTester) {
	c.namedTesters = append(c.namedTesters, testers...)
}

func (c *NamedTesterContainer) Test(t *testing.T) {
	for _, tester := range c.namedTesters {
		if tester == nil {
			continue
		}

		t.Run(tester.Name(), tester.Test)
	}
}
