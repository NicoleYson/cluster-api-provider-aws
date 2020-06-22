/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package awsvalidation

import (
	"math/rand"
	"testing"
	"time"
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func TestIsValidTagKey(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Valid 128 Character Tag Key",
			args: args{StringWithCharset(128, "abcd123")},
			want: true,
		},
		{
			name: "Tag Key is too long",
			args: args{StringWithCharset(500, "abcd123")},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidTagKey(tt.args.key); got != tt.want {
				t.Errorf("IsValidTagKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidTagValue(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Valid 128 Character Tag Value",
			args: args{StringWithCharset(128, "abcd123")},
			want: true,
		},
		{
			name: "Tag Value is too long",
			args: args{StringWithCharset(500, "abcd123")},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidTagValue(tt.args.value); got != tt.want {
				t.Errorf("IsValidTagValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
