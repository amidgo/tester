package tester

import "testing"

type Tester interface {
	Test(t *testing.T)
}

type TesterContainer struct {
	testers []Tester
}

func NewTesterContainer(t *testing.T) *TesterContainer {
	tester := &TesterContainer{
		testers: make([]Tester, 0),
	}

	t.Cleanup(func() { tester.Test(t) })

	return tester
}

func (tt *TesterContainer) AddTester(t Tester) {
	tt.testers = append(tt.testers, t)
}

func (tt *TesterContainer) Test(t *testing.T) {
	for _, tester := range tt.testers {
		if tester == nil {
			continue
		}
		tester.Test(t)
	}
}
