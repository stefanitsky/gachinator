package gachinate

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
	re        = regexp.MustCompile(`(эс)|(о)|(к[ао]м)|([фФ]ак)|(т[её]мный)|(гей)|(глубокий|глубокое|глубоко|глубокая)`)
	replacers = map[int][]byte{
		0: []byte("ASS"),
		1: []byte("♂"),
		2: []byte("CUM"),
		3: []byte("FUCK"),
		4: []byte("DARK"),
		5: []byte("GAY"),
		6: []byte("DEEP"),
	}
)

var (
	offset   int
	m        match
	orig     []byte
	subIndex int
	start    int
	end      int
	repl     []byte
	found    bool
)

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

func findMatch(indexes *[]int, b *[]byte) (m match) {
	if len(*indexes) < 4 {
		return m
	}

	subIndex = 0

	for i := 2; i < len(*indexes); i += 2 {
		start, end = (*indexes)[i], (*indexes)[i+1]
		if start != -1 && end != -1 {
			repl, found = replacers[subIndex]
			if found {
				m.start = start
				m.end = end
				m.replacer = repl
				m.found = true
				break
			}
		}
		subIndex++
	}

	return m
}
