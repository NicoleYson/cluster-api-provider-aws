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

import "regexp"

var (
	re = map[string]*regexp.Regexp{
		"tagKey": regexp.MustCompile(`^.[\p{L}\d_.:/=+-@]{1,128}$`),

		// The tag value must be a minimum of 0 and a maximum of 256 Unicode characters in UTF-8.
		//Some services don't permit tags with an empty value (length of 0), going with the most stringent
		"tagValue": regexp.MustCompile(`^.[\p{L}\d_.:/=+-@]{1,256}$`),
	}
)

func IsValidTagKey(key string) bool {
	return re["tagKey"].MatchString(key)
}

func IsValidTagValue(value string) bool {
	return re["tagValue"].MatchString(value)
}
