// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package tokenizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenizer(t *testing.T) {
	tt := []struct {
		name     string
		opts     []Option
		content  string
		expected []string
	}{
		{
			name:     "Without keeping separator",
			content:  "I believe life is an intelligent thing: that things aren't random.",
			expected: []string{"I", "believe", "life", "is", "an", "intelligent", "thing", "that", "things", "aren't", "random"},
		},
		{
			name:     "With keeping separator",
			opts:     []Option{KeepSeparator()},
			content:  "I believe life is an intelligent thing: that things aren't random.",
			expected: []string{"I", " ", "believe", " ", "life", " ", "is", " ", "an", " ", "intelligent", " ", "thing", ":", " ", "that", " ", "things", " ", "aren't", " ", "random", "."},
		},
		{
			name:     "With UTF-8 separator",
			opts:     []Option{WithSeparator("’")},
			content:  "s’ajoute",
			expected: []string{"s", "ajoute"},
		},
		{
			name:     "With UTF-8 separator and keeping separator",
			opts:     []Option{WithSeparator("’"), KeepSeparator()},
			content:  "s’ajoute",
			expected: []string{"s", "’", "ajoute"},
		},
		{
			name:     "With space separator",
			opts:     []Option{WithSeparator(" ")},
			content:  "I believe life is an intelligent thing: that things aren't random.",
			expected: []string{"I", "believe", "life", "is", "an", "intelligent", "thing:", "that", "things", "aren't", "random."},
		},
		{
			name:     "With Japanese text",
			opts:     []Option{WithSeparator("。、")},
			content:  "デザインとは、見た目や感触だけではありません。デザインがどのように機能するかが重要です。",
			expected: []string{"デザインとは", "見た目や感触だけではありません", "デザインがどのように機能するかが重要です"},
		},
		{
			name:     "With Japanese text and keeping separator",
			opts:     []Option{WithSeparator("。、"), KeepSeparator()},
			content:  "デザインとは、見た目や感触だけではありません。デザインがどのように機能するかが重要です。",
			expected: []string{"デザインとは", "、", "見た目や感触だけではありません", "。", "デザインがどのように機能するかが重要です", "。"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tokenizer := New(tc.opts...)
			tokens := tokenizer.Tokenize(tc.content)

			assert.Equal(t, tc.expected, tokens)
		})
	}
}

func TestConvertSeparator(t *testing.T) {
	expected := map[rune]bool{
		'\t': true,
		'\n': true,
		' ':  true,
	}

	assert.Equal(t, expected, convertSeparator("\t\n "))
}

func BenchmarkTokenizer(b *testing.B) {
	tokenizer := New()

	for n := 0; n < b.N; n++ {
		tokenizer.Tokenize("I believe life is an intelligent thing: that things aren't random.")
	}
}
