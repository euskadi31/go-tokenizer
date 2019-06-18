// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
//Plus Kenny Karnama that add the Restore function

package tokenizer

// Tokenizer interface
type Tokenizer interface {
	KeepSeparator()
	Tokenize(content string) []string
	Restore(tokens []string) string
}

type tokenizer struct {
	sep          [256]uint8
	keepSep      bool
	SepTokens    []byte
	SepTokensIdx []int
}

// New Tokenizer
func New() Tokenizer {
	return &tokenizer{
		sep: convertSeparator("\t\n\r ,.:?\"!;()"),
	}
}

// NewWithSeparator Tokenizer
func NewWithSeparator(sep string) Tokenizer {
	return &tokenizer{
		sep: convertSeparator(sep),
	}
}

func (t *tokenizer) KeepSeparator() {
	t.keepSep = true
}

//Restore try to construct again the original string (as match as possible)
//the given tokenized string
//by using previous delimiter saved
func (t *tokenizer) Restore(tokens []string) string {
	res := ""
	n := len(tokens)
	for idx, tokens := range tokens {
		if idx < n {
			res += tokens
			sep := ""

			if len(t.SepTokens) > 0 {
				isRepeated := false
				before := -1
				for {
					a := t.SepTokensIdx[0]
					b := a
					if len(t.SepTokensIdx) > 1 {
						b = t.SepTokensIdx[1]
					}
					diff := b - a
					if diff > 1 {
						if !isRepeated {
							sep += string(t.SepTokens[0])
							t.SepTokens = t.SepTokens[1:]
							t.SepTokensIdx = t.SepTokensIdx[1:]
							isRepeated = false
						} else {
							if before != -1 && (a-before) == 1 {
								sep += string(t.SepTokens[0])
								t.SepTokens = t.SepTokens[1:]
								t.SepTokensIdx = t.SepTokensIdx[1:]
								isRepeated = false
							}
						}
						break

					} else if diff == 1 {
						sep += string(t.SepTokens[0])
						sep += string(t.SepTokens[1])
						t.SepTokens = t.SepTokens[2:]
						t.SepTokensIdx = t.SepTokensIdx[2:]
						isRepeated = true
						before = b
					} else if diff == 0 {
						sep += string(t.SepTokens[0])
						t.SepTokens = t.SepTokens[1:]
						t.SepTokensIdx = t.SepTokensIdx[1:]
						isRepeated = false
						before = a
					}

					if len(t.SepTokens) < 1 {
						break
					}
				}
			}
			res += sep
		}
	}
	for len(t.SepTokens) > 0 {
		res += string(t.SepTokens[0])
		t.SepTokens = t.SepTokens[1:]
	}

	return res
}

//Tokenize string according to the delimiter
func (t *tokenizer) Tokenize(content string) []string {
	length := len(content)
	cut := make([]int, length+1)
	var sepInfo []byte
	var sepIdx []int
	for i := 0; i < length; i++ {
		r := content[i]

		if int(t.sep[r]) == 1 {
			cut[i] = 1
			sepInfo = append(sepInfo, r)
			sepIdx = append(sepIdx, i)
		}
	}

	t.SepTokens = sepInfo
	t.SepTokensIdx = sepIdx
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
