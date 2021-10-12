package main

import (
	"reflect"
	"testing"
)

func Test_grep(t *testing.T) {
	type args struct {
		input      string
		pattern    string
		parameters Flag
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Test 1",
			args: args{
				input:      "sd1 ff\nssa\nkk",
				pattern:    "sd",
				parameters: Flag{after: 0},
			},
			want: []string{"sd1 ff"},
		},
		{
			name: "Test 2",
			args: args{
				input:      "sd1 ff\nssa\nsd kk\nii\nss",
				pattern:    "sd",
				parameters: Flag{after: 2},
			},
			want: []string{"sd1 ff", "ssa", "sd kk", "ii", "ss"},
		},
		{
			name: "Test 3",
			args: args{
				input:      "sd1 ff\nssa\nsd kk\nii\nss",
				pattern:    "sd",
				parameters: Flag{after: 1},
			},
			want: []string{"sd1 ff", "ssa", "sd kk", "ii"},
		},
		{
			name: "Test 4",
			args: args{
				input:      "sd1 ff\nssa\nsd kk\nii\nss",
				pattern:    "sd",
				parameters: Flag{after: 3},
			},
			want: []string{"sd1 ff", "ssa", "sd kk", "ii", "ss"},
		},

		{
			name: "Test 5",
			args: args{
				input:      "sd1 ff\nssa\nsd kk\nii\nss",
				pattern:    "sd",
				parameters: Flag{before: 2},
			},
			want: []string{"sd1 ff", "ssa", "sd kk"},
		},
		{
			name: "Test 6",
			args: args{
				input:      "sd1 ff\nssa\nsd kk\nii\nss",
				pattern:    "sd",
				parameters: Flag{before: 1},
			},
			want: []string{"sd1 ff", "ssa", "sd kk"},
		},
		{
			name: "Test 7",
			args: args{
				input:      "sd1 ff\nssa\nsd kk\nii\nss",
				pattern:    "sd",
				parameters: Flag{context: 1},
			},
			want: []string{"sd1 ff", "ssa", "sd kk", "ii"},
		},
		{
			name: "Test 8",
			args: args{
				input:      "sd1 ff\nssa\nsd kk\nii\nss",
				pattern:    "sd",
				parameters: Flag{count: true},
			},
			want: []string{"2"},
		},
		{
			name: "Test 9",
			args: args{
				input:      "sd1 ff\nssa\nsd kk\nii\nss",
				pattern:    "Sd",
				parameters: Flag{ignoreCase: false},
			},
			want: nil,
		},
		{
			name: "Test 10",
			args: args{
				input:      "Sd1 ff\nssa\nSD kk\nii\nss",
				pattern:    "sd",
				parameters: Flag{ignoreCase: true},
			},
			want: []string{"Sd1 ff", "SD kk"},
		},
		{
			name: "Test 12",
			args: args{
				input:      "sd1 ff\nssa\nsd kk\nii\nss",
				pattern:    "sd",
				parameters: Flag{invert: true},
			},
			want: []string{"ssa", "ii", "ss"},
		},
		{
			name: "Test 13",
			args: args{
				input:      "sd1 ff\nssa\nsd kk\nii\nss\nsd",
				pattern:    "sd",
				parameters: Flag{fixed: true},
			},
			want: []string{"sd"},
		},
		{
			name: "Test 13",
			args: args{
				input:      "sd1 ff\nssa\nsd kk\nii\nss\nsd",
				pattern:    "sd",
				parameters: Flag{lineNum: true},
			},
			want: []string{"1:sd1 ff", "3:sd kk", "6:sd"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := grep(tt.args.input, tt.args.pattern, tt.args.parameters)
			if (err != nil) != tt.wantErr {
				t.Errorf("grep() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("grep() got = %v, want %v", got, tt.want)
			}
		})
	}
}
