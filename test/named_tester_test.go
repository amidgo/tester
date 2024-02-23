package test_test

import (
	"testing"

	"github.com/amidgo/tester"
	"github.com/amidgo/tester/mock"
)

func Test_NamedTesterContainer(t *testing.T) {
	var tester tester.NamedTesterContainer

	tester.AddNamedTester(mock.NewMockedNamedTester(t))
	tester.AddNamedTester(mock.NewMockedNamedTester(t))
	tester.AddNamedTester(mock.NewMockedNamedTester(t))
	tester.AddNamedTester(mock.NewMockedNamedTester(t))

	tester.AddNamedTesters(
		mock.NewMockedNamedTester(t),
		mock.NewMockedNamedTester(t),
		mock.NewMockedNamedTester(t),
		mock.NewMockedNamedTester(t),
	)

	tester.Test(t)
}

func Test_NamedTesterContainer_Not_Test_Call_Case(t *testing.T) {
	var tester tester.NamedTesterContainer

	nt := new(mock.MockedNamedTester)

	tester.AddNamedTester(nt)

	tester.AddNamedTesters(
		nt,
	)

	nt.AssertNotCalled(t)
}

func Test_RunNamedTesters(t *testing.T) {
	tester.RunNamedTesters(
		t,
		mock.NewMockedNamedTester(t),
		mock.NewMockedNamedTester(t),
		mock.NewMockedNamedTester(t),
		mock.NewMockedNamedTester(t),
	)
}
