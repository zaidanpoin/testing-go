package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorld1(t *testing.T) {
	result := HelloWorld("ashel")
	ekspetasi := "hello ashel"

	assert.Equal(t, ekspetasi, result, "result must be hello ashel")
	fmt.Println("testd done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("ashel")
	ekspetasi := "hello ashel"

	require.Equal(t, ekspetasi, result, "result must be hello ashel")
	fmt.Println("testd done")
}
