// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package tokenizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenizer(t *testing.T) {
	tokenizer := New()

	tokens := tokenizer.Tokenize("I believe life is an intelligent thing: that things aren't random.")

	assert.Equal(t, []string{"I", "believe", "life", "is", "an", "intelligent", "thing", "that", "things", "aren't", "random"}, tokens)
}

func TestTokenizerWithSeparator(t *testing.T) {
	tokenizer := NewWithSeparator(" ")

	tokens := tokenizer.Tokenize("I believe life is an intelligent thing: that things aren't random.")

	assert.Equal(t, []string{"I", "believe", "life", "is", "an", "intelligent", "thing:", "that", "things", "aren't", "random."}, tokens)
}

func TestTokenizerWithKeepingSeparator(t *testing.T) {
	tokenizer := New()
	tokenizer.KeepSeparator()

	tokens := tokenizer.Tokenize("I believe life is an intelligent thing: that things aren't random.")

	assert.Equal(t, []string{"I", " ", "believe", " ", "life", " ", "is", " ", "an", " ", "intelligent", " ", "thing", ":", " ", "that", " ", "things", " ", "aren't", " ", "random", "."}, tokens)
}

func TestConvertSeparator(t *testing.T) {
	assert.Equal(t, [256]uint8{'\t': 1, '\n': 1, ' ': 1}, convertSeparator("\t\n "))
}

func BenchmarkTokenizer(b *testing.B) {
	tokenizer := New()

	for n := 0; n < b.N; n++ {
		tokenizer.Tokenize("I believe life is an intelligent thing: that things aren't random.")
	}
}
