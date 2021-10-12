package main

import "testing"

func Test_unpackString(t *testing.T) {

	tests := []struct {
		name    string
		input   string
		wantOutput    string
		wantErr bool
	}{
		{
			name: "Test #1",
			input: "a4bc2d5e",
			wantOutput: "aaaabccddddde",
			wantErr: false,
		},
		{
			name: "Test #2",
			input: "abcd",
			wantOutput: "abcd",
			wantErr: false,
		},
		{
			name: "Test #3",
			input: "45",
			wantOutput: "",
			wantErr: true,
		},
		{
			name: "Test #4",
			input: "",
			wantOutput: "",
			wantErr: false,
		},
		{
			name: "Test #5",
			input: "界3世2",
			wantOutput: "界界界世世",
			wantErr: false,
		},
		{
			name: "Test #6",
			input: "qwe\\4\\5",
			wantOutput: "qwe45",
			wantErr: false,
		},
		{
			name: "Test #7",
			input: "qwe\\45",
			wantOutput: "qwe44444",
			wantErr: false,
		},
		{
			name: "Test #8",
			input: "qwe\\\\5",
			wantOutput: "qwe\\\\\\\\\\",
			wantErr: false,
		},
		{
			name: "Test #9",
			input: "qwe\\\\10",
			wantOutput: "qwe\\\\\\\\\\\\\\\\\\\\",
			wantErr: false,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			gotOutput, err := unpackString(testCase.input)
			if (err != nil) != testCase.wantErr {
				t.Errorf("unpack() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if gotOutput != testCase.wantOutput {
				t.Errorf("unpack() got = %v, want %v", gotOutput, testCase.wantOutput)
			}
		})
	}
}