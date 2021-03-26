package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersion(t *testing.T) {
	a := assert.New(t)

	var testString = "0.0.0"

	SetVersion(testString)
	v := GetVersion()
	a.Equal(testString, v)
}
