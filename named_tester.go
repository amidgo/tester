package tester

import "testing"

type NamedTester interface {
	Name() string
	Test(t *testing.T)
}

type NamedTesterContainer struct {
	namedTesters []NamedTester
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
