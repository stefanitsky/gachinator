package gachinator

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"testing"
)

type GachinateTestCase struct {
	input          string
	expectedOutput string
}

type GachinateSimpleTestCase struct {
	input          string
	expectedOutput string
	langConfig     langConfig
}

func TestGachinateRu(t *testing.T) {
	cases := []GachinateTestCase{
		{
			input:          "фактор",
			expectedOutput: "FUCKт♂р",
		},
		{
			input:          "соответственно",
			expectedOutput: "с♂♂тветственн♂",
		},
		{
			input:          "комбайнер",
			expectedOutput: "CUMбайнер",
		},
		{
			input:          "респиратор",
			expectedOutput: "респират♂р",
		},
		{
			input:          "темный",
			expectedOutput: "DARK",
		},
		{
			input:          "тёмный",
			expectedOutput: "DARK",
		},
		{
			input:          "гей",
			expectedOutput: "GAY",
		},
		{
			input:          "Сергей",
			expectedOutput: "СерGAY",
		},
		{
			input:          " Фактор - это причина",
			expectedOutput: " FUCKт♂р - эт♂ причина",
		},
		{
			input:          "нырнуть глубоко",
			expectedOutput: "нырнуть DEEP",
		},
		{
			input:          "глубокий",
			expectedOutput: "DEEP",
		},
		{
			input:          "глубокое",
			expectedOutput: "DEEP",
		},
		{
			input:          "глубокая кроличья нора",
			expectedOutput: "DEEP кр♂личья н♂ра",
		},
		{
			input:          "доллары",
			expectedOutput: "BUCKS",
		},
		{
			input:          "долларов",
			expectedOutput: "BUCKS",
		},
	}

	for _, c := range cases {
		result := string(GachinateRU([]byte(c.input)))

		if result != c.expectedOutput {
			t.Errorf("\nExpected: %v\nGot:%v\nOriginal:%v\n", c.expectedOutput, result, c.input)
		}
	}
}

func TestGachinateEn(t *testing.T) {
	cases := []GachinateTestCase{
		{
			input:          "cool",
			expectedOutput: "c♂♂l",
		},
		{
			input:          "manufacturable",
			expectedOutput: "manuFUCKturable",
		},
		{
			input:          "message",
			expectedOutput: "mASSage",
		},
		{
			input:          "come",
			expectedOutput: "CUM",
		},
		{
			input:          "become",
			expectedOutput: "beCUM",
		},
		{
			input:          "semidarkness",
			expectedOutput: "semiDARKnASS",
		},
	}

	for _, c := range cases {
		result := string(GachinateEN([]byte(c.input)))

		if result != c.expectedOutput {
			t.Errorf("\nExpected: %v\nGot:%v\nOriginal:%v\n", c.expectedOutput, result, c.input)
		}
	}
}

// Just test that it executes with different configs (no need complex test duplication)
func TestGachinate(t *testing.T) {
	cases := []GachinateSimpleTestCase{
		{
			input:          "круто",
			expectedOutput: "крут♂",
			langConfig:     RussianConfig,
		},
		{
			input:          "cool",
			expectedOutput: "c♂♂l",
			langConfig:     EnglishConfig,
		},
	}

	for _, c := range cases {
		result := string(Gachinate([]byte(c.input), c.langConfig))

		if result != c.expectedOutput {
			t.Errorf("\nExpected: %v\nGot: %v\nOriginal: %v\nConfig: %v", c.expectedOutput, result, c.input, c.langConfig)
		}
	}
}

func TestFindLangConfig(t *testing.T) {
	for lang, expectedConfig := range langCodeToLangConfig {
		lc, err := FindLangConfig(lang)
		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
		}

		// reflect.DeepEqual will not work, that's why we need to test manually
		if lc.re != expectedConfig.re {
			t.Errorf("Regex missmatch, expected %v, got %v", expectedConfig.re, lc.re)
		}

		foundReplacersLen, expectedReplacersLen := len(lc.replacers), len(expectedConfig.replacers)
		if foundReplacersLen != expectedReplacersLen {
			t.Errorf("Replacers missmatch, expected %v, got %v", lc.replacers, expectedConfig.replacers)
		}

		for k, v := range lc.replacers {
			if v2, ok := expectedConfig.replacers[k]; !ok || !bytes.Equal(v, v2) {
				t.Errorf("Language configs are not match, expected %v, got %v", expectedConfig, lc)
			}
		}
	}
}

func TestFindLangConfigNotFound(t *testing.T) {
	expectedErr := `language config for "fake" is not found (available configs are: ru, en).`
	lc, err := FindLangConfig("fake")
	if lc != nil || err.Error() != expectedErr {
		t.Errorf(`Expected error: "%v", got language config: %v and error: %v`, expectedErr, lc, err)
	}
}

func BenchmarkGachinate(b *testing.B) {
	b.StopTimer()
	f, err := os.Open("testdata/benchmark.txt")
	if err != nil {
		b.Error(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	input, err := io.ReadAll(r)
	if err != nil {
		b.Error(err)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		GachinateRU(input)
	}
	b.StopTimer()
}
