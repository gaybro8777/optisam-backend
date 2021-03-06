// Copyright (C) 2019 Orange
// 
// This software is distributed under the terms and conditions of the 'Apache License 2.0'
// license which can be found in the file 'License.txt' in this package distribution 
// or at 'http://www.apache.org/licenses/LICENSE-2.0'. 

package helper

import (
	"reflect"
	"testing"
)

func TestRemoveSlice(t *testing.T) {
	type args struct {
		originalSlice      []string
		removeElementSlice []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "SUCCESS - Both slice empty",
			args: args{originalSlice: []string{}, removeElementSlice: []string{}},
			want: []string{},
		},
		{
			name: "SUCCESS - removeElementSlice empty",
			args: args{originalSlice: []string{"a", "b", "c"}, removeElementSlice: []string{}},
			want: []string{"a", "b", "c"},
		},
		{
			name: "SUCCESS - OriginalSlice empty",
			args: args{originalSlice: []string{}, removeElementSlice: []string{"a", "b", "c"}},
			want: []string{},
		},
		{
			name: "SUCCESS - Both Slices have some overlappig elements",
			args: args{originalSlice: []string{"a", "b", "y", "z"}, removeElementSlice: []string{"a", "b", "c"}},
			want: []string{"y", "z"},
		},
		{
			name: "SUCCESS - No Overlapping Element",
			args: args{originalSlice: []string{"a", "b", "c"}, removeElementSlice: []string{"x", "y", "z"}},
			want: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveElements(tt.args.originalSlice, tt.args.removeElementSlice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendIfNotExists(t *testing.T) {
	type args struct {
		slice           []string
		addElementSlice []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "SUCCESS - elem and target slice empty",
			args: args{slice: []string{}, addElementSlice: []string{}},
			want: []string{},
		},
		{
			name: "SUCCESS - elem provided and target slice empty",
			args: args{slice: []string{}, addElementSlice: []string{"a"}},
			want: []string{"a"},
		},
		{
			name: "SUCCESS - elem not provided and target slice provided",
			args: args{slice: []string{"a"}, addElementSlice: []string{}},
			want: []string{"a"},
		},
		{
			name: "SUCCESS - elem and target slice provided",
			args: args{slice: []string{"a", "b", "c"}, addElementSlice: []string{"d"}},
			want: []string{"a", "b", "c", "d"},
		},
		{
			name: "SUCCESS - provided element already exists",
			args: args{slice: []string{"a", "b", "c"}, addElementSlice: []string{"a", "c"}},
			want: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AppendElementsIfNotExists(tt.args.slice, tt.args.addElementSlice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendIfNotExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	type args struct {
		slice []string
		val   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "SUCCESS - string match regex",
			args: args{slice: []string{`^products\.csv$`, `^applications\.csv$`},
				val: "products.csv"},
			want: true,
		},
		{name: "SUCCESS - string does not match regex",
			args: args{slice: []string{`^products\.csv$`, `^applications\.csv$`},
				val: "products1.csv"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.slice, tt.args.val); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
