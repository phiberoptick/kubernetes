/*
Copyright 2025 The Kubernetes Authors.

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

// +k8s:validation-gen=TypeMeta
// +k8s:validation-gen-scheme-registry=k8s.io/code-generator/cmd/validation-gen/testscheme.Scheme

// This is a test package.
package list

import "k8s.io/code-generator/cmd/validation-gen/testscheme"

var localSchemeBuilder = testscheme.New()

type StructSlice struct {
	TypeMeta int

	// +k8s:eachVal=+k8s:validateFalse="field AtomicSliceStringField[*]"
	AtomicSliceStringField []StringType `json:"atomicSliceStringField"`

	// +k8s:eachVal=+k8s:validateFalse="field AtomicSliceTypeField[*]"
	AtomicSliceTypeField IntSliceType `json:"atomicSliceTypeField"`

	// +k8s:eachVal=+k8s:validateFalse="field AtomicSliceComparableField[*]"
	AtomicSliceComparableField []ComparableStruct `json:"atomicSliceComparableField"`

	// +k8s:eachVal=+k8s:validateFalse="field AtomicSliceNonComparableField[*]"
	AtomicSliceNonComparableField []NonComparableStruct `json:"atomicSliceNonComparableField"`

	// +k8s:listType=set
	// +k8s:eachVal=+k8s:validateFalse="field SetSliceComparableField[*]"
	SetSliceComparableField []ComparableStruct `json:"setSliceComparableField"`

	// +k8s:listType=set
	// +k8s:eachVal=+k8s:validateFalse="field SetSliceNonComparableField[*]"
	SetSliceNonComparableField []NonComparableStruct `json:"setSliceNonComparableField"`

	// +k8s:listType=map
	// +k8s:listMapKey=key
	// +k8s:eachVal=+k8s:validateFalse="field MapSliceComparableField[*]"
	MapSliceComparableField []ComparableStructWithKey `json:"mapSliceComparableField"`

	// +k8s:listType=map
	// +k8s:listMapKey=key
	// +k8s:eachVal=+k8s:validateFalse="field MapSliceNonComparableField[*]"
	MapSliceNonComparableField []NonComparableStructWithKey `json:"mapSliceNonComparableField"`
}

type StringType string

type IntSliceType []int

type ComparableStruct struct {
	IntField int `json:"intField"`
}

// +k8s:validateFalse="type NonComparableStruct"
type NonComparableStruct struct {
	IntPtrField *int `json:"intPtrField"`
}

type ComparableStructWithKey struct {
	Key      string `json:"key"`
	IntField int    `json:"intField"`
}

// +k8s:validateFalse="type NonComparableStructWithKey"
type NonComparableStructWithKey struct {
	Key         string `json:"key"`
	IntPtrField *int   `json:"intPtrField"`
}
