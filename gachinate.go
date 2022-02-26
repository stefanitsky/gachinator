// Package gachinator provides utils to gachinate input text.
package gachinator

import (
	"fmt"
	"regexp"
)

type replacer []byte
type reGroupIndex int
type replacers map[reGroupIndex]replacer

// LangConfig contains *regexp.Regexp to find text to replace
// and replacers map to find a replace.
type LangConfig struct {
	re                 *regexp.Regexp
	replacers          replacers
	submatchIndexesLen int
}

var (
	russianReplacers = replacers{
		0:  replacer("ASS"),
		1:  replacer("ANAL"),
		2:  replacer("CUM"),
		3:  replacer("FUCK"),
		4:  replacer("DARK"),
		5:  replacer("GAY"),
		6:  replacer("DEEP"),
		7:  replacer("BUCKS"),
		8:  replacer("FANTASY"),
		9:  replacer("SWALL♂W"),
		10: replacer("MASTER"),
		11: replacer("DUNGE♂N"),
		12: replacer("B♂Y"),
		13: replacer("SLAVE"),
		14: replacer("♂"),
	}
	englishReplacers = replacers{
		0: replacer("ASS"),
		1: replacer("♂"),
		2: replacer("CUM"),
		3: replacer("FUCK"),
		4: replacer("DARK"),
	}
)

var (
	// RussianConfig is a config to work with a russian text
	RussianConfig = LangConfig{
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
		replacers:          russianReplacers,
		submatchIndexesLen: russianReplacers.getSubmatchIndexesLen(),
	}
	// EnglishConfig is a config to work with an english text
	EnglishConfig = LangConfig{
		re:                 regexp.MustCompile(`([eE]ss)|(o)|([cC][ou]m(?:e))|([fF]ac)|([dD]ark)`),
		replacers:          englishReplacers,
		submatchIndexesLen: englishReplacers.getSubmatchIndexesLen(),
	}
	langCodeToLangConfig = map[string]LangConfig{
		"ru": RussianConfig,
		"en": EnglishConfig,
	}
)

// LangConfigNotFoundError returns, when a config for specified language code is not found.
type LangConfigNotFoundError struct {
	lang string
}

func (e *LangConfigNotFoundError) Error() string {
	return fmt.Sprintf("language config for \"%v\" is not found (available configs are: ru, en).", e.lang)
}

var (
	insertOffset  int
	originalBytes []byte
)

var (
	currentReGroupIndex reGroupIndex
	startIndex          int
	endIndex            int
	repl                replacer
)

// Gachinate translates your input text with specified language config and returns gachinated variant.
func Gachinate(b []byte, lc LangConfig) []byte {
	return *gachinate(&b, lc)
}

// GachinateRU translates russian input
func GachinateRU(b []byte) []byte {
	return *gachinate(&b, RussianConfig)
}

// GachinateEN translates english input
func GachinateEN(b []byte) []byte {
	return *gachinate(&b, EnglishConfig)
}

// FindLangConfig finds language config by lang code
// Example:
// lc, err := FindLangConfig("ru")
// If config is not found, than LangConfigNotFoundError is returned
func FindLangConfig(lang string) (*LangConfig, error) {
	lc, found := langCodeToLangConfig[lang]
	if !found {
		return nil, &LangConfigNotFoundError{lang}
	}

	return &lc, nil
}

func gachinate(b *[]byte, lc LangConfig) *[]byte {
	allSubmatchIndexes := lc.re.FindAllSubmatchIndex(*b, -1)

	insertOffset = 0
	for _, loc := range allSubmatchIndexes {
		currentReGroupIndex = 0

		// We don't need first 2 values, just skip them
		for i := 2; i < lc.submatchIndexesLen; i += 2 {
			startIndex, endIndex = loc[i], loc[i+1]
			// Each loc contains all submatch indexes, even match is not found,
			// that's why we need to check, that match indexes not equal -1
			if startIndex != -1 && endIndex != -1 {
				// We don't need to check, that replacer exist,
				// because if not, it's not yet finished language config
				// and panic will be raised
				repl = lc.replacers[currentReGroupIndex]
				originalBytes = (*b)[startIndex:endIndex]
				*b = append((*b)[:startIndex+insertOffset], append(repl, (*b)[endIndex+insertOffset:]...)...)
				insertOffset += len(repl) - len(originalBytes)
				// Each loc contains only 1 found submatch indexes,
				// we don't need to iterate next, so we skip them
				continue
			}
			currentReGroupIndex++
		}
	}

	return b
}

// getSubmatchIndexesLen returns total submatch indexes length
// It means (length of replacers + 1) * 2, where 2 it's a multiplier,
// which says, that we will get 2 indexes on each group
// from the regexp.FindAllSubmatchIndex
func (r *replacers) getSubmatchIndexesLen() int {
	return (len(*r) + 1) * 2
}
