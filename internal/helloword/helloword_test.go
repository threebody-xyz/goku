package helloword

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelloWord(t *testing.T) {
	helloWord := HelloWord()
	targetString := "Hello Word!"
	assert.Equal(t, helloWord, targetString, "The two words should be the same.")

}
