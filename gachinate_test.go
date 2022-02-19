package gachinate

import (
	"bufio"
	"io"
	"os"
	"testing"
)

type GachinateTestCase struct {
	input          string
	expectedOutput string
}

func TestGachinate(t *testing.T) {
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
	}

	for _, c := range cases {
		result := string(Gachinate([]byte(c.input)))

		if result != c.expectedOutput {
			t.Errorf("\nExpected: %v\nGot:%v\nOriginal:%v\n", c.expectedOutput, result, c.input)
		}
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
		Gachinate(input)
	}
	b.StopTimer()
}
