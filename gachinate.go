// Package gachinator provides utils to gachinate input text.
package gachinator

import (
	"regexp"
)

type match struct {
	start    int
	end      int
	replacer []byte
	found    bool
}

var (
	re        = regexp.MustCompile(`([эЭ]с)|(о)|([кК][ао]м)|([фФ]ак)|(т[её]мн(ый|ое|ая|о|ые|ых))|(гей)|(глубок(ий|ое|ая|о|и|ие|ого))|(доллар(ов|ы))`)
	replacers = map[int][]byte{
		0: []byte("ASS"),
		1: []byte("♂"),
		2: []byte("CUM"),
		3: []byte("FUCK"),
		4: []byte("DARK"),
		// 5: useless suffixes group
		6: []byte("GAY"),
		7: []byte("DEEP"),
		// 8: useless suffixes group
		9: []byte("BUCKS"),
	}
)

var (
	offset int
	m      match
	orig   []byte
)

var (
	subIndex int
	start    int
	end      int
	repl     []byte
	found    bool
)

// Gachinates your input text and returns gachinated variant.
func Gachinate(b []byte) []byte {
	allSubmatchIndexes := re.FindAllSubmatchIndex(b, -1)

	offset = 0
	for _, loc := range allSubmatchIndexes {
		m = findMatch(&loc, &b)
		if m.found {
			orig = b[m.start:m.end]
			b = append(b[:m.start+offset], append(m.replacer, b[m.end+offset:]...)...)
			offset += len(m.replacer) - len(orig)
		}
	}

	return b
}

// Finds match by found regex submatch indexes and returns match struct.
func findMatch(indexes *[]int, b *[]byte) (m match) {
	subIndex = 0

	for i := 2; i < len(*indexes); i += 2 {
		start, end = (*indexes)[i], (*indexes)[i+1]
		if start != -1 && end != -1 {
			repl, found = replacers[subIndex]
			if found {
				m.start = start
				m.end = end
				m.replacer = repl
				m.found = found
				break
			}
		}
		subIndex++
	}

	return m
}
