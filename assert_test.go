// Copyright 2023 Christian Banse
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// This file is part of The Money Gopher.

// package assert contains logic to assert test values.
package assert

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type SomeStruct struct {
	A int
	B string
}

func TestEquals(t *testing.T) {
	type args struct {
		t        *testing.T
		expected *SomeStruct
		actual   *SomeStruct
		opts     []cmp.Option
	}
	tests := []struct {
		name   string
		args   args
		wantOk bool
	}{
		{
			name: "happy path",
			args: args{
				t:        &testing.T{},
				expected: &SomeStruct{A: 1, B: "foo"},
				actual:   &SomeStruct{A: 1, B: "foo"},
			},
			wantOk: true,
		},
		{
			name: "sad path",
			args: args{
				t:        &testing.T{},
				expected: &SomeStruct{A: 1, B: "foo"},
				actual:   &SomeStruct{A: 2, B: "bar"},
			},
			wantOk: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOk := Equals(tt.args.t, tt.args.expected, tt.args.actual, tt.args.opts...); gotOk != tt.wantOk {
				t.Errorf("Equals() = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
