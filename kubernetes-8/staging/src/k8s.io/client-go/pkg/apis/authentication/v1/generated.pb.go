/*
Copyright 2017 The Kubernetes Authors.

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

// Code generated by protoc-gen-gogo.
// source: github.com/sourcegraph/monorepo-test-1/kubernetes-8/pkg/apis/authentication/v1/generated.proto
// DO NOT EDIT!

/*
	Package v1 is a generated protocol buffer package.

	It is generated from these files:
		github.com/sourcegraph/monorepo-test-1/kubernetes-8/pkg/apis/authentication/v1/generated.proto

	It has these top-level messages:
		ExtraValue
		TokenReview
		TokenReviewSpec
		TokenReviewStatus
		UserInfo
*/
package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strings "strings"
import reflect "reflect"
import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func (m *ExtraValue) Reset()                    { *m = ExtraValue{} }
func (*ExtraValue) ProtoMessage()               {}
func (*ExtraValue) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *TokenReview) Reset()                    { *m = TokenReview{} }
func (*TokenReview) ProtoMessage()               {}
func (*TokenReview) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *TokenReviewSpec) Reset()                    { *m = TokenReviewSpec{} }
func (*TokenReviewSpec) ProtoMessage()               {}
func (*TokenReviewSpec) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func (m *TokenReviewStatus) Reset()                    { *m = TokenReviewStatus{} }
func (*TokenReviewStatus) ProtoMessage()               {}
func (*TokenReviewStatus) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{3} }

func (m *UserInfo) Reset()                    { *m = UserInfo{} }
func (*UserInfo) ProtoMessage()               {}
func (*UserInfo) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{4} }

func init() {
	proto.RegisterType((*ExtraValue)(nil), "k8s.io.client-go.pkg.apis.authentication.v1.ExtraValue")
	proto.RegisterType((*TokenReview)(nil), "k8s.io.client-go.pkg.apis.authentication.v1.TokenReview")
	proto.RegisterType((*TokenReviewSpec)(nil), "k8s.io.client-go.pkg.apis.authentication.v1.TokenReviewSpec")
	proto.RegisterType((*TokenReviewStatus)(nil), "k8s.io.client-go.pkg.apis.authentication.v1.TokenReviewStatus")
	proto.RegisterType((*UserInfo)(nil), "k8s.io.client-go.pkg.apis.authentication.v1.UserInfo")
}
func (m ExtraValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m ExtraValue) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m) > 0 {
		for _, s := range m {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *TokenReview) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenReview) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ObjectMeta.Size()))
	n1, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Spec.Size()))
	n2, err := m.Spec.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x1a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Status.Size()))
	n3, err := m.Status.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	return i, nil
}

func (m *TokenReviewSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenReviewSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Token)))
	i += copy(dAtA[i:], m.Token)
	return i, nil
}

func (m *TokenReviewStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenReviewStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	if m.Authenticated {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.User.Size()))
	n4, err := m.User.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	dAtA[i] = 0x1a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Error)))
	i += copy(dAtA[i:], m.Error)
	return i, nil
}

func (m *UserInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Username)))
	i += copy(dAtA[i:], m.Username)
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.UID)))
	i += copy(dAtA[i:], m.UID)
	if len(m.Groups) > 0 {
		for _, s := range m.Groups {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Extra) > 0 {
		for k := range m.Extra {
			dAtA[i] = 0x22
			i++
			v := m.Extra[k]
			msgSize := 0
			if (&v) != nil {
				msgSize = (&v).Size()
				msgSize += 1 + sovGenerated(uint64(msgSize))
			}
			mapSize := 1 + len(k) + sovGenerated(uint64(len(k))) + msgSize
			i = encodeVarintGenerated(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64((&v).Size()))
			n5, err := (&v).MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n5
		}
	}
	return i, nil
}

func encodeFixed64Generated(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Generated(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m ExtraValue) Size() (n int) {
	var l int
	_ = l
	if len(m) > 0 {
		for _, s := range m {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *TokenReview) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Status.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *TokenReviewSpec) Size() (n int) {
	var l int
	_ = l
	l = len(m.Token)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *TokenReviewStatus) Size() (n int) {
	var l int
	_ = l
	n += 2
	l = m.User.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Error)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *UserInfo) Size() (n int) {
	var l int
	_ = l
	l = len(m.Username)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.UID)
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Groups) > 0 {
		for _, s := range m.Groups {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if len(m.Extra) > 0 {
		for k, v := range m.Extra {
			_ = k
			_ = v
			l = v.Size()
			mapEntrySize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + l + sovGenerated(uint64(l))
			n += mapEntrySize + 1 + sovGenerated(uint64(mapEntrySize))
		}
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *TokenReview) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TokenReview{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "TokenReviewSpec", "TokenReviewSpec", 1), `&`, ``, 1) + `,`,
		`Status:` + strings.Replace(strings.Replace(this.Status.String(), "TokenReviewStatus", "TokenReviewStatus", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TokenReviewSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TokenReviewSpec{`,
		`Token:` + fmt.Sprintf("%v", this.Token) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TokenReviewStatus) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TokenReviewStatus{`,
		`Authenticated:` + fmt.Sprintf("%v", this.Authenticated) + `,`,
		`User:` + strings.Replace(strings.Replace(this.User.String(), "UserInfo", "UserInfo", 1), `&`, ``, 1) + `,`,
		`Error:` + fmt.Sprintf("%v", this.Error) + `,`,
		`}`,
	}, "")
	return s
}
func (this *UserInfo) String() string {
	if this == nil {
		return "nil"
	}
	keysForExtra := make([]string, 0, len(this.Extra))
	for k := range this.Extra {
		keysForExtra = append(keysForExtra, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForExtra)
	mapStringForExtra := "map[string]ExtraValue{"
	for _, k := range keysForExtra {
		mapStringForExtra += fmt.Sprintf("%v: %v,", k, this.Extra[k])
	}
	mapStringForExtra += "}"
	s := strings.Join([]string{`&UserInfo{`,
		`Username:` + fmt.Sprintf("%v", this.Username) + `,`,
		`UID:` + fmt.Sprintf("%v", this.UID) + `,`,
		`Groups:` + fmt.Sprintf("%v", this.Groups) + `,`,
		`Extra:` + mapStringForExtra + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ExtraValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExtraValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExtraValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			*m = append(*m, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TokenReview) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TokenReview: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenReview: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TokenReviewSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TokenReviewSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenReviewSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TokenReviewStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TokenReviewStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenReviewStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authenticated", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Authenticated = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.User.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UserInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UserInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Groups", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Groups = append(m.Groups, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Extra", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthGenerated
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(dAtA[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.Extra == nil {
				m.Extra = make(map[string]ExtraValue)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var mapmsglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					mapmsglen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if mapmsglen < 0 {
					return ErrInvalidLengthGenerated
				}
				postmsgIndex := iNdEx + mapmsglen
				if mapmsglen < 0 {
					return ErrInvalidLengthGenerated
				}
				if postmsgIndex > l {
					return io.ErrUnexpectedEOF
				}
				mapvalue := &ExtraValue{}
				if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
					return err
				}
				iNdEx = postmsgIndex
				m.Extra[mapkey] = *mapvalue
			} else {
				var mapvalue ExtraValue
				m.Extra[mapkey] = mapvalue
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGenerated(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("k8s.io/client-go/pkg/apis/authentication/v1/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 655 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x53, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xb5, 0xf3, 0x53, 0x25, 0x93, 0xaf, 0x1f, 0x65, 0x24, 0xa4, 0x28, 0x12, 0x4e, 0x14, 0x58,
	0x74, 0x51, 0xc6, 0xa4, 0xa0, 0x52, 0x15, 0x10, 0xaa, 0x45, 0x85, 0xba, 0x00, 0xa4, 0x81, 0x22,
	0xc4, 0x06, 0x26, 0xce, 0xad, 0x63, 0x52, 0xff, 0x68, 0x3c, 0x36, 0xed, 0xae, 0x8f, 0xc0, 0x92,
	0x25, 0xaf, 0xc1, 0x1b, 0x74, 0x47, 0x77, 0xb0, 0x40, 0x15, 0x0d, 0x2f, 0x82, 0x66, 0x3c, 0xd4,
	0x2e, 0x69, 0x85, 0xda, 0xdd, 0xcc, 0x99, 0x7b, 0xce, 0xbd, 0xe7, 0xde, 0xb9, 0xe8, 0xc1, 0x64,
	0x35, 0x21, 0x7e, 0x64, 0x4f, 0xd2, 0x21, 0xf0, 0x10, 0x04, 0x24, 0x76, 0x3c, 0xf1, 0x6c, 0x16,
	0xfb, 0x89, 0xcd, 0x52, 0x31, 0x86, 0x50, 0xf8, 0x2e, 0x13, 0x7e, 0x14, 0xda, 0xd9, 0xc0, 0xf6,
	0x20, 0x04, 0xce, 0x04, 0x8c, 0x48, 0xcc, 0x23, 0x11, 0xe1, 0xa5, 0x9c, 0x4d, 0x0a, 0x36, 0x89,
	0x27, 0x1e, 0x91, 0x6c, 0x72, 0x9a, 0x4d, 0xb2, 0x41, 0xe7, 0x96, 0xe7, 0x8b, 0x71, 0x3a, 0x24,
	0x6e, 0x14, 0xd8, 0x5e, 0xe4, 0x45, 0xb6, 0x12, 0x19, 0xa6, 0xdb, 0xea, 0xa6, 0x2e, 0xea, 0x94,
	0x8b, 0x77, 0xee, 0xea, 0xd2, 0x58, 0xec, 0x07, 0xcc, 0x1d, 0xfb, 0x21, 0xf0, 0xbd, 0xa2, 0xb8,
	0x00, 0x04, 0x3b, 0xa3, 0xa4, 0x8e, 0x7d, 0x1e, 0x8b, 0xa7, 0xa1, 0xf0, 0x03, 0x98, 0x21, 0xac,
	0xfc, 0x8b, 0x90, 0xb8, 0x63, 0x08, 0xd8, 0x0c, 0xef, 0xce, 0x79, 0xbc, 0x54, 0xf8, 0x3b, 0xb6,
	0x1f, 0x8a, 0x44, 0xf0, 0x19, 0x52, 0xc9, 0x53, 0x02, 0x3c, 0x03, 0x5e, 0x18, 0x82, 0x5d, 0x16,
	0xc4, 0x3b, 0x70, 0x86, 0xa7, 0xfe, 0x3d, 0x84, 0x36, 0x76, 0x05, 0x67, 0xaf, 0xd8, 0x4e, 0x0a,
	0xb8, 0x8b, 0xea, 0xbe, 0x80, 0x20, 0x69, 0x9b, 0xbd, 0xea, 0x62, 0xd3, 0x69, 0x4e, 0x8f, 0xba,
	0xf5, 0x4d, 0x09, 0xd0, 0x1c, 0x5f, 0x6b, 0x7c, 0xfa, 0xdc, 0x35, 0xf6, 0x7f, 0xf4, 0x8c, 0xfe,
	0x97, 0x0a, 0x6a, 0xbd, 0x8c, 0x26, 0x10, 0x52, 0xc8, 0x7c, 0xf8, 0x80, 0xdf, 0xa1, 0x86, 0xec,
	0xdb, 0x88, 0x09, 0xd6, 0x36, 0x7b, 0xe6, 0x62, 0x6b, 0xf9, 0x36, 0xd1, 0x23, 0x2c, 0xdb, 0x28,
	0x86, 0x28, 0xa3, 0x49, 0x36, 0x20, 0xcf, 0x87, 0xef, 0xc1, 0x15, 0x4f, 0x41, 0x30, 0x07, 0x1f,
	0x1c, 0x75, 0x8d, 0xe9, 0x51, 0x17, 0x15, 0x18, 0x3d, 0x51, 0xc5, 0x6f, 0x51, 0x2d, 0x89, 0xc1,
	0x6d, 0x57, 0x94, 0xfa, 0x43, 0x72, 0x91, 0x0f, 0x42, 0x4a, 0xa5, 0xbe, 0x88, 0xc1, 0x75, 0xfe,
	0xd3, 0xa9, 0x6a, 0xf2, 0x46, 0x95, 0x30, 0xf6, 0xd0, 0x5c, 0x22, 0x98, 0x48, 0x93, 0x76, 0x55,
	0xa5, 0x78, 0x74, 0xf9, 0x14, 0x4a, 0xc6, 0xf9, 0x5f, 0x27, 0x99, 0xcb, 0xef, 0x54, 0xcb, 0xf7,
	0x57, 0xd0, 0x95, 0xbf, 0xea, 0xc1, 0x37, 0x50, 0x5d, 0x48, 0x48, 0xf5, 0xae, 0xe9, 0xcc, 0x6b,
	0x66, 0x3d, 0x8f, 0xcb, 0xdf, 0xfa, 0x5f, 0x4d, 0x74, 0x75, 0x26, 0x0b, 0xbe, 0x8f, 0xe6, 0x4b,
	0xc5, 0xc0, 0x48, 0x49, 0x34, 0x9c, 0x6b, 0x5a, 0x62, 0x7e, 0xbd, 0xfc, 0x48, 0x4f, 0xc7, 0xe2,
	0xd7, 0xa8, 0x96, 0x26, 0xc0, 0x75, 0x53, 0x57, 0x2e, 0xe6, 0x78, 0x2b, 0x01, 0xbe, 0x19, 0x6e,
	0x47, 0x45, 0x37, 0x25, 0x42, 0x95, 0xa2, 0x74, 0x04, 0x9c, 0x47, 0x5c, 0x35, 0xb3, 0xe4, 0x68,
	0x43, 0x82, 0x34, 0x7f, 0xeb, 0x7f, 0xab, 0xa0, 0xc6, 0x1f, 0x15, 0xbc, 0x84, 0x1a, 0x92, 0x19,
	0xb2, 0x00, 0x74, 0x1b, 0x16, 0x34, 0x49, 0xc5, 0x48, 0x9c, 0x9e, 0x44, 0xe0, 0xeb, 0xa8, 0x9a,
	0xfa, 0x23, 0x55, 0x78, 0xd3, 0x69, 0xe9, 0xc0, 0xea, 0xd6, 0xe6, 0x63, 0x2a, 0x71, 0xdc, 0x47,
	0x73, 0x1e, 0x8f, 0xd2, 0x58, 0x0e, 0x53, 0xfe, 0x65, 0x24, 0xe7, 0xf0, 0x44, 0x21, 0x54, 0xbf,
	0xe0, 0x6d, 0x54, 0x07, 0xf9, 0xf9, 0xdb, 0xb5, 0x5e, 0x75, 0xb1, 0xb5, 0xbc, 0x7e, 0x39, 0xf7,
	0x44, 0x2d, 0xd0, 0x46, 0x28, 0xf8, 0x5e, 0xc9, 0xa5, 0xc4, 0x68, 0x2e, 0xdf, 0xe1, 0x7a, 0xc9,
	0x54, 0x0c, 0x5e, 0x40, 0xd5, 0x09, 0xec, 0xe5, 0x0e, 0xa9, 0x3c, 0xe2, 0x67, 0xa8, 0x9e, 0xc9,
	0xfd, 0xd3, 0x53, 0x58, 0xbd, 0x58, 0x1d, 0xc5, 0xfe, 0xd2, 0x5c, 0x66, 0xad, 0xb2, 0x6a, 0x3a,
	0x37, 0x0f, 0x8e, 0x2d, 0xe3, 0xf0, 0xd8, 0x32, 0xbe, 0x1f, 0x5b, 0xc6, 0xfe, 0xd4, 0x32, 0x0f,
	0xa6, 0x96, 0x79, 0x38, 0xb5, 0xcc, 0x9f, 0x53, 0xcb, 0xfc, 0xf8, 0xcb, 0x32, 0xde, 0x54, 0xb2,
	0xc1, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x6b, 0x11, 0x20, 0xa4, 0x05, 0x00, 0x00,
}
