package test_test

import (
	"testing"

	"github.com/amidgo/tester"
	"github.com/amidgo/tester/mock"
)

func Test_RunNamedTesters(t *testing.T) {
	tester.RunNamedTesters(
		t,
		mock.NewMockedNamedTester(t),
		mock.NewMockedNamedTester(t),
		mock.NewMockedNamedTester(t),
		mock.NewMockedNamedTester(t),
	)
}
