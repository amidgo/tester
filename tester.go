package tester

import "testing"

type Tester interface {
	Test(t *testing.T)
}

type TesterFunc func(t *testing.T)

func (tf TesterFunc) Test(t *testing.T) {
	tf(t)
}

type TesterContainer struct {
	testers []Tester
}

func RunTesters(t *testing.T, testers ...Tester) {
	for _, tester := range testers {
		if tester == nil {
			continue
		}

		tester.Test(t)
	}
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
