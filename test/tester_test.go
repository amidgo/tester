package test_test

import (
	"testing"

	"github.com/amidgo/tester"
	"github.com/amidgo/tester/mock"
)

func Test_TesterContainer(t *testing.T) {
	var tester tester.TesterContainer

	tester.AddTester(mock.NewMockedNamedTester(t))
	tester.AddTester(mock.NewMockedNamedTester(t))
	tester.AddTester(mock.NewMockedNamedTester(t))
	tester.AddTester(mock.NewMockedNamedTester(t))

	tester.Test(t)
}

func Test_TesterContainer_Not_Test_Call_Case(t *testing.T) {
	var tester tester.TesterContainer

	nt := new(mock.MockedNamedTester)

	tester.AddTester(nt)

	nt.AssertNotCalled(t)
}
