// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package tokenizer

import (
	"unicode"
	"unicode/utf8"
)

type Option func(*tokenizer)

func WithSeparator(sep string) Option {
	return func(t *tokenizer) {
		t.mode = separatorModeLegacy
		t.sep = convertSeparator(sep)
	}
}

func KeepSeparator() Option {
	return func(t *tokenizer) {
		t.keepSep = true
	}
}

func WithUnicodeSeparator(tables ...*unicode.RangeTable) Option {
	return func(t *tokenizer) {
		t.mode = separatorModeUnicode

		if len(tables) > 0 {
			t.tables = tables

			return
		}

		t.tables = []*unicode.RangeTable{
			unicode.Punct,
			unicode.Symbol,
			unicode.Space,
			unicode.Cc,
		}
	}
}

func WithIgnoreSeparators(runes ...rune) Option {
	return func(t *tokenizer) {
		t.ignoreSeparators = map[rune]bool{}

		for _, r := range runes {
			t.ignoreSeparators[r] = true
		}
	}
}

// Tokenizer interface.
type Tokenizer interface {
	Tokenize(content string) []string
}

type separatorMode int8

const (
	separatorModeLegacy separatorMode = iota
	separatorModeUnicode
)

type tokenizer struct {
	mode             separatorMode
	sep              map[rune]bool
	ignoreSeparators map[rune]bool
	tables           []*unicode.RangeTable
	keepSep          bool
}

// New Tokenizer with options.
func New(opts ...Option) Tokenizer {
	t := &tokenizer{
		mode:             separatorModeLegacy,
		sep:              convertSeparator("\t\n\r ,.:?\"!;()"),
		ignoreSeparators: convertSeparator("'"),
	}

	for _, opt := range opts {
		opt(t)
	}

	return t
}

func (t tokenizer) isSeparator(r rune) bool {
	if t.ignoreSeparators[r] {
		return false
	}

	if t.mode == separatorModeUnicode {
		return unicode.IsOneOf(t.tables, r)
	}

	return t.sep[r]
}

func (t tokenizer) getCutsList(content string) []int {
	length := len(content)

	cut := make([]int, length)

	lastPos := 0

	for {
		r, size := utf8.DecodeRuneInString(content[lastPos:])

		if t.isSeparator(r) {
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
		remainingSkip = (size - 1)

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
