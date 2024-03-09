// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package tokenizer

// Tokenizer interface.
type Tokenizer interface {
	KeepSeparator()
	Tokenize(content string) []string
}

type tokenizer struct {
	sep     [256]uint8
	keepSep bool
}

// New Tokenizer.
func New() Tokenizer {
	return &tokenizer{
		sep: convertSeparator("\t\n\r ,.:?\"!;()"),
	}
}

// NewWithSeparator Tokenizer.
func NewWithSeparator(sep string) Tokenizer {
	return &tokenizer{
		sep: convertSeparator(sep),
	}
}

func (t *tokenizer) KeepSeparator() {
	t.keepSep = true
}

func (t tokenizer) Tokenize(content string) []string {
	length := len(content)
	cut := make([]int, length+1)

	for i := 0; i < length; i++ {
		r := content[i]

		if int(t.sep[r]) == 1 {
			cut[i] = 1
		}
	}

	tokens := []string{}
	lastCut := 0
	previousIsSep := false

	for i := 0; i < length; i++ {
		if cut[i] == 1 {
			if previousIsSep && !t.keepSep {
				previousIsSep = true
				lastCut++

				continue
			}

			tokens = append(tokens, content[lastCut:i])
			lastCut = i

			if !t.keepSep {
				lastCut++
			}

			previousIsSep = true
		} else {
			if t.keepSep && previousIsSep {
				tokens = append(tokens, content[lastCut:i])
				lastCut = i
			}

			previousIsSep = false
		}
	}

	if previousIsSep && t.keepSep || !previousIsSep {
		tokens = append(tokens, content[lastCut:])
	}

	return tokens
}

func convertSeparator(sep string) [256]uint8 {
	separators := [256]uint8{}

	for _, r := range sep {
		separators[r] = 1
	}

	return separators
}
