package main

import (
	"reflect"
	"testing"
)

func Test_findAnagram(t *testing.T) {

	tests := []struct {
		name       string
		input      []string
		wantOutput map[string][]string
	}{
		{
			name:       "Test 1",
			input:      []string{"пятак", "ПЯТКА", "тяпка", "пятак", "слиток", "Листок", "столик", "сок"},
			wantOutput: map[string][]string{"пятак": {"пятак", "пятка", "тяпка"}, "слиток": {"листок", "слиток", "столик"}},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			if gotOutput := findAnagram(testCase.input); !reflect.DeepEqual(gotOutput, testCase.wantOutput) {
				t.Errorf("findAnagram() = %v, wantOutput %v", gotOutput, testCase.wantOutput)
			}
		})
	}
}
