package mock

import (
	"testing"

	"github.com/amidgo/tester"
	"github.com/stretchr/testify/assert"
)

type MockedNamedTester struct {
	isCalled bool
}

func NewMockedNamedTester(t *testing.T) tester.NamedTester {
	tester := new(MockedNamedTester)

	t.Cleanup(func() { tester.assertCalled(t) })

	return tester
}

func (mnt *MockedNamedTester) Name() string {
	return ""
}

func (mnt *MockedNamedTester) Test(t *testing.T) {
	if mnt.isCalled {
		t.Fatal("tester called twice")
	}

	mnt.isCalled = true
}

func (mnt *MockedNamedTester) assertCalled(t *testing.T) {
	assert.True(t, mnt.isCalled, "mocked named tester not called")
}

func (mnt *MockedNamedTester) AssertNotCalled(t *testing.T) {
	assert.False(t, mnt.isCalled, "mocked named tester is called")
}
