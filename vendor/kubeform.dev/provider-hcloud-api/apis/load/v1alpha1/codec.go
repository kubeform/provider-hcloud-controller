/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

func GetEncoder() map[string]jsoniter.ValEncoder {
	return map[string]jsoniter.ValEncoder{
		jsoniter.MustGetKind(reflect2.TypeOf(BalancerSpecAlgorithm{}).Type1()):              BalancerSpecAlgorithmCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BalancerServiceSpecHealthCheck{}).Type1()):     BalancerServiceSpecHealthCheckCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BalancerServiceSpecHealthCheckHttp{}).Type1()): BalancerServiceSpecHealthCheckHttpCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BalancerServiceSpecHttp{}).Type1()):            BalancerServiceSpecHttpCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(BalancerSpecAlgorithm{}).Type1()):              BalancerSpecAlgorithmCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BalancerServiceSpecHealthCheck{}).Type1()):     BalancerServiceSpecHealthCheckCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BalancerServiceSpecHealthCheckHttp{}).Type1()): BalancerServiceSpecHealthCheckHttpCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BalancerServiceSpecHttp{}).Type1()):            BalancerServiceSpecHttpCodec{},
	}
}

func getEncodersWithout(typ string) map[string]jsoniter.ValEncoder {
	origMap := GetEncoder()
	delete(origMap, typ)
	return origMap
}

func getDecodersWithout(typ string) map[string]jsoniter.ValDecoder {
	origMap := GetDecoder()
	delete(origMap, typ)
	return origMap
}

// +k8s:deepcopy-gen=false
type BalancerSpecAlgorithmCodec struct {
}

func (BalancerSpecAlgorithmCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*BalancerSpecAlgorithm)(ptr) == nil
}

func (BalancerSpecAlgorithmCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*BalancerSpecAlgorithm)(ptr)
	var objs []BalancerSpecAlgorithm
	if obj != nil {
		objs = []BalancerSpecAlgorithm{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BalancerSpecAlgorithm{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (BalancerSpecAlgorithmCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*BalancerSpecAlgorithm)(ptr) = BalancerSpecAlgorithm{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []BalancerSpecAlgorithm

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BalancerSpecAlgorithm{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*BalancerSpecAlgorithm)(ptr) = objs[0]
			} else {
				*(*BalancerSpecAlgorithm)(ptr) = BalancerSpecAlgorithm{}
			}
		} else {
			*(*BalancerSpecAlgorithm)(ptr) = BalancerSpecAlgorithm{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj BalancerSpecAlgorithm

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BalancerSpecAlgorithm{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*BalancerSpecAlgorithm)(ptr) = obj
		} else {
			*(*BalancerSpecAlgorithm)(ptr) = BalancerSpecAlgorithm{}
		}
	default:
		iter.ReportError("decode BalancerSpecAlgorithm", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type BalancerServiceSpecHealthCheckCodec struct {
}

func (BalancerServiceSpecHealthCheckCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*BalancerServiceSpecHealthCheck)(ptr) == nil
}

func (BalancerServiceSpecHealthCheckCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*BalancerServiceSpecHealthCheck)(ptr)
	var objs []BalancerServiceSpecHealthCheck
	if obj != nil {
		objs = []BalancerServiceSpecHealthCheck{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BalancerServiceSpecHealthCheck{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (BalancerServiceSpecHealthCheckCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*BalancerServiceSpecHealthCheck)(ptr) = BalancerServiceSpecHealthCheck{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []BalancerServiceSpecHealthCheck

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BalancerServiceSpecHealthCheck{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*BalancerServiceSpecHealthCheck)(ptr) = objs[0]
			} else {
				*(*BalancerServiceSpecHealthCheck)(ptr) = BalancerServiceSpecHealthCheck{}
			}
		} else {
			*(*BalancerServiceSpecHealthCheck)(ptr) = BalancerServiceSpecHealthCheck{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj BalancerServiceSpecHealthCheck

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BalancerServiceSpecHealthCheck{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*BalancerServiceSpecHealthCheck)(ptr) = obj
		} else {
			*(*BalancerServiceSpecHealthCheck)(ptr) = BalancerServiceSpecHealthCheck{}
		}
	default:
		iter.ReportError("decode BalancerServiceSpecHealthCheck", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type BalancerServiceSpecHealthCheckHttpCodec struct {
}

func (BalancerServiceSpecHealthCheckHttpCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*BalancerServiceSpecHealthCheckHttp)(ptr) == nil
}

func (BalancerServiceSpecHealthCheckHttpCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*BalancerServiceSpecHealthCheckHttp)(ptr)
	var objs []BalancerServiceSpecHealthCheckHttp
	if obj != nil {
		objs = []BalancerServiceSpecHealthCheckHttp{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BalancerServiceSpecHealthCheckHttp{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (BalancerServiceSpecHealthCheckHttpCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*BalancerServiceSpecHealthCheckHttp)(ptr) = BalancerServiceSpecHealthCheckHttp{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []BalancerServiceSpecHealthCheckHttp

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BalancerServiceSpecHealthCheckHttp{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*BalancerServiceSpecHealthCheckHttp)(ptr) = objs[0]
			} else {
				*(*BalancerServiceSpecHealthCheckHttp)(ptr) = BalancerServiceSpecHealthCheckHttp{}
			}
		} else {
			*(*BalancerServiceSpecHealthCheckHttp)(ptr) = BalancerServiceSpecHealthCheckHttp{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj BalancerServiceSpecHealthCheckHttp

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BalancerServiceSpecHealthCheckHttp{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*BalancerServiceSpecHealthCheckHttp)(ptr) = obj
		} else {
			*(*BalancerServiceSpecHealthCheckHttp)(ptr) = BalancerServiceSpecHealthCheckHttp{}
		}
	default:
		iter.ReportError("decode BalancerServiceSpecHealthCheckHttp", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type BalancerServiceSpecHttpCodec struct {
}

func (BalancerServiceSpecHttpCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*BalancerServiceSpecHttp)(ptr) == nil
}

func (BalancerServiceSpecHttpCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*BalancerServiceSpecHttp)(ptr)
	var objs []BalancerServiceSpecHttp
	if obj != nil {
		objs = []BalancerServiceSpecHttp{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BalancerServiceSpecHttp{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (BalancerServiceSpecHttpCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*BalancerServiceSpecHttp)(ptr) = BalancerServiceSpecHttp{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []BalancerServiceSpecHttp

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BalancerServiceSpecHttp{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*BalancerServiceSpecHttp)(ptr) = objs[0]
			} else {
				*(*BalancerServiceSpecHttp)(ptr) = BalancerServiceSpecHttp{}
			}
		} else {
			*(*BalancerServiceSpecHttp)(ptr) = BalancerServiceSpecHttp{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj BalancerServiceSpecHttp

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BalancerServiceSpecHttp{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*BalancerServiceSpecHttp)(ptr) = obj
		} else {
			*(*BalancerServiceSpecHttp)(ptr) = BalancerServiceSpecHttp{}
		}
	default:
		iter.ReportError("decode BalancerServiceSpecHttp", "unexpected JSON type")
	}
}