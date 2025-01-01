// Copyright 2023-2024 Christian Banse
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
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Want is a function type that can be executed to chain several assertions
// together.
type Want[T any] func(*testing.T, T) bool

// Equals compares expected to actual and returns true if they are equal.
func Equals[T any](t *testing.T, expected T, actual T, opts ...cmp.Option) (ok bool) {
	t.Helper()

	return EqualsFunc(t, expected, actual, func(expected T, actual T) bool {
		return cmp.Equal(expected, actual, opts...)
	})
}

// Equals compares expected to actual using the equals function and returns true
// if they are equal.
func EqualsFunc[T any](t testing.TB, expected T, actual T, equals func(expected T, actual T) bool) (ok bool) {
	t.Helper()

	ok = equals(expected, actual)

	if !ok {
		t.Errorf("%T = %v, want %v", actual, actual, expected)
	}

	return ok
}

// NotEquals compares expected to actual and returns true if they are not equal.
func NotEquals[T any](t *testing.T, expected T, actual T, opts ...cmp.Option) (ok bool) {
	t.Helper()

	ok = !cmp.Equal(expected, actual, opts...)

	if !ok {
		t.Errorf("%T != %v, want %v", actual, actual, expected)
	}

	return ok
}

// Is asserts that value is of type T. If it succeeds, it returns the value
// casted to T. If it fails, we fatally fail the test, because we cannot
// continue.
func Is[T any](t *testing.T, value any) T {
	t.Helper()

	cast, ok := value.(T)
	if !ok {
		// We cannot continue
		t.Fatalf("%v is not of type %T", value, new(T))
	}

	return cast
}

// NoError asserts that err does not contain an error.
func NoError(t *testing.T, err error) bool {
	return Equals(t, nil, err)
}

// NotNil asserts that value is not nil. If it fails, we fatally fail the test,
// because we will probably run into a panic otherwise anyway.
func NotNil(t *testing.T, value any) bool {
	t.Helper()

	ok := NotEquals(t, nil, &value)
	if !ok {
		// We cannot continue
		t.Fatalf("variable of type %T should not be nil", value)
	}

	return ok
}

// NotNil asserts that value is nil
func Nil(t *testing.T, value any) bool {
	t.Helper()

	return Equals(t, nil, value)
}

// ErrorIs asserts that an error is the expected error using [errors.Is].
func ErrorIs(t testing.TB, expected error, actual error) bool {
	t.Helper()

	return EqualsFunc(t, expected, actual, func(expected error, actual error) bool {
		return errors.Is(actual, expected)
	})
}
