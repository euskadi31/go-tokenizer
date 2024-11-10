// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package tokenizer

import (
	"unicode/utf8"
)

type Option func(*tokenizer)

func WithSeparator(sep string) Option {
	return func(t *tokenizer) {
		t.sep = convertSeparator(sep)
	}
}

func KeepSeparator() Option {
	return func(t *tokenizer) {
		t.keepSep = true
	}
}

// Tokenizer interface.
type Tokenizer interface {
	Tokenize(content string) []string
}

type tokenizer struct {
	sep     map[rune]bool
	keepSep bool
}

// New Tokenizer with options.
func New(opts ...Option) Tokenizer {
	t := &tokenizer{
		sep: convertSeparator("\t\n\r ,.:?\"!;()"),
	}

	for _, opt := range opts {
		opt(t)
	}

	return t
}

func (t tokenizer) getCutsList(content string) []int {
	length := len(content)

	cut := make([]int, length)

	lastPos := 0

	for {
		r, size := utf8.DecodeRuneInString(content[lastPos:])

		if t.sep[r] {
			cut[lastPos] = size
		}

		lastPos += size

		if lastPos >= length {
			break
		}
	}

	return cut
}

func (t tokenizer) Tokenize(content string) []string {
	cut := t.getCutsList(content)

	tokens := []string{}
	lastCut := 0
	previousIsSep := false
	remainingSkip := 0

	for i, size := range cut {
		if size == 0 {
			if remainingSkip > 0 {
				remainingSkip--

				continue
			}

			previousIsSep = false

			continue
		}

		if !previousIsSep {
			tokens = append(tokens, content[lastCut:i])
		}

		lastCut = i + size
		remainingSkip = size

		if t.keepSep {
			tokens = append(tokens, content[i:lastCut])
		}

		previousIsSep = true
	}

	if !previousIsSep {
		tokens = append(tokens, content[lastCut:])
	}

	return tokens
}

func convertSeparator(sep string) map[rune]bool {
	separators := map[rune]bool{}

	b := sep

	for len(b) > 0 {
		r, size := utf8.DecodeRuneInString(b)

		separators[r] = true

		b = b[size:]
	}

	return separators
}
