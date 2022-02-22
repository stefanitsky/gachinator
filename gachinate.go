// Package gachinator provides utils to gachinate input text.
package gachinator

import (
	"fmt"
	"regexp"
)

type match struct {
	start    int
	end      int
	replacer []byte
	found    bool
}

type langConfig struct {
	re        *regexp.Regexp
	replacers map[int][]byte
}

var (
	RussianConfig = langConfig{
		re: regexp.MustCompile(`` +
			`([эЭ]с)` +
			`|(онал)` +
			`|([кК][ао]м)` +
			`|([фФ]ак)` +
			`|([тТ][её]мн(?:ый|ое|ая|о|ые|ых))` +
			`|([гГ]ей)` +
			`|([гГ]лубок(?:ий|ое|ая|о|и|ие|ого)?)` +
			`|([дД]оллар(?:ов|ы)?)` +
			`|([фФ]антази[яйи])` +
			`|([гГ]лот(?:ает|ать|ай|ок))` +
			`|([мМ]астер(?:а|ы|ов)?)` +
			`|([пП]одзем(?:ный|ное|ная|ные|ных|ного|елье|елья))` +
			`|([бБ]ой|[мМ]альчик(?:а|и|ов)?|[пП]арен(?:ь|[её]к))` +
			`|(?:^|\s)([сС]луг[аиеу])` +
			`|(о)`,
		),
		replacers: map[int][]byte{
			0:  []byte("ASS"),
			1:  []byte("ANAL"),
			2:  []byte("CUM"),
			3:  []byte("FUCK"),
			4:  []byte("DARK"),
			5:  []byte("GAY"),
			6:  []byte("DEEP"),
			7:  []byte("BUCKS"),
			8:  []byte("FANTASY"),
			9:  []byte("SWALL♂W"),
			10: []byte("MASTER"),
			11: []byte("DUNGE♂N"),
			12: []byte("B♂Y"),
			13: []byte("SLAVE"),
			14: []byte("♂"),
		},
	}
	EnglishConfig = langConfig{
		re: regexp.MustCompile(`([eE]ss)|(o)|([cC][ou]m(e))|([fF]ac)|([dD]ark)`),
		replacers: map[int][]byte{
			0: []byte("ASS"),
			1: []byte("♂"),
			2: []byte("CUM"),
			// 3: useless suffix group
			4: []byte("FUCK"),
			5: []byte("DARK"),
		},
	}
	langCodeToLangConfig = map[string]langConfig{
		"ru": RussianConfig,
		"en": EnglishConfig,
	}
)

type LangConfigNotFoundError struct {
	lang string
}

func (e *LangConfigNotFoundError) Error() string {
	return fmt.Sprintf("language config for \"%v\" is not found (available configs are: ru, en).", e.lang)
}

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

// Gachinates your input text with specified language config and returns gachinated variant.
func Gachinate(b []byte, lc langConfig) []byte {
	return *gachinate(&b, lc)
}

// Gachinate russian input
func GachinateRU(b []byte) []byte {
	return *gachinate(&b, RussianConfig)
}

// Gachinate english input
func GachinateEN(b []byte) []byte {
	return *gachinate(&b, EnglishConfig)
}

// Finds language config by lang code
// Example:
// lc, err := FindLangConfig("ru")
func FindLangConfig(lang string) (*langConfig, error) {
	lc, found := langCodeToLangConfig[lang]
	if !found {
		return nil, &LangConfigNotFoundError{lang}
	}

	return &lc, nil
}

func gachinate(b *[]byte, lc langConfig) *[]byte {
	allSubmatchIndexes := lc.re.FindAllSubmatchIndex(*b, -1)

	offset = 0
	for _, loc := range allSubmatchIndexes {
		m = findMatch(&loc, b, &lc)
		if m.found {
			orig = (*b)[m.start:m.end]
			*b = append((*b)[:m.start+offset], append(m.replacer, (*b)[m.end+offset:]...)...)
			offset += len(m.replacer) - len(orig)
		}
	}

	return b
}

// Finds match by found regex submatch indexes and returns match struct.
func findMatch(indexes *[]int, b *[]byte, lc *langConfig) (m match) {
	subIndex = 0

	for i := 2; i < len(*indexes); i += 2 {
		start, end = (*indexes)[i], (*indexes)[i+1]
		if start != -1 && end != -1 {
			repl, found = lc.replacers[subIndex]
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
