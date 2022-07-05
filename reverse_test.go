package main

import (
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "default",
			args: args{
				input: `
				{
					
					
					"other": "3",
						"1170": "1",
						"70#": "2",
						"*70":"0"
			
				}
				
				`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverse(tt.args.input)
		})
	}
}

func Test_generateJavaMap(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "default case",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generateJavaMap()
		})
	}
}
