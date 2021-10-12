package main

import (
	"reflect"
	"testing"
)

func Test_cut(t *testing.T) {
	type args struct {
		param Flag
		input []string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput []string
		wantErr    bool
	}{
		{
			name: "Test 1",
			args: args{
				param: Flag{
					fields:       "1",
					delimiter:    "\t",
					withSeparate: false,
				},
				input: []string {"1\txxx\twww\tkkk","2\taaa\tyyy\tbbb"},
			},
			wantOutput: []string {"1","\n","2","\n"},
			wantErr: false,
		},
		{
			name: "Test 2",
			args: args{
				param: Flag{
					fields:       "1-3",
					delimiter:    "\t",
					withSeparate: false,
				},
				input: []string {"1\txxx\twww\tkkk","2\taaa\tyyy\tbbb"},
			},
			wantOutput: []string {"1","\t","xxx","\t","www","\t","\n","2","\t","aaa","\t","yyy","\t","\n"},
			wantErr: false,
		},
		{
			name: "Test 3",
			args: args{
				param: Flag{
					fields:       "1-3",
					delimiter:    " ",
					withSeparate: false,
				},
				input: []string {"1 xxx www kkk","2 aaa yyy bbb"},
			},
			wantOutput: []string {"1"," ","xxx"," ","www"," ","\n","2"," ","aaa"," ","yyy"," ","\n"},
			wantErr: false,
		},
		{
			name: "Test 4",
			args: args{
				param: Flag{
					fields:       "1",
					delimiter:    " ",
					withSeparate: false,
				},
				input: []string {"1 xxx www kkk","2 aaa yyy bbb", "aaa"},
			},
			wantOutput: []string {"1","\n","2","\n","aaa","\n"},
			wantErr: false,
		},
		{
			name: "Test 5",
			args: args{
				param: Flag{
					fields:       "1",
					delimiter:    " ",
					withSeparate: true,
				},
				input: []string {"1 xxx www kkk","2 aaa yyy bbb", "aaa"},
			},
			wantOutput: []string {"1","\n","2","\n"},
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := cut(test.args.param, test.args.input)
			if (err != nil) != test.wantErr {
				t.Errorf("cut() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.wantOutput) {
				t.Errorf("cut() got = %v, wantOutput %v", got, test.wantOutput)
			}
		})
	}
}