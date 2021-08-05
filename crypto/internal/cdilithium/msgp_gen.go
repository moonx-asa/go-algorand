package cdilithium

// Code generated by github.com/algorand/msgp DO NOT EDIT.

import (
	_ "runtime/cgo"
	_ "unsafe"

	"github.com/algorand/msgp/msgp"
)

// The following msgp objects are implemented in this file:
// Dil2PrivateKey
//        |-----> (*) MarshalMsg
//        |-----> (*) CanMarshalMsg
//        |-----> (*) UnmarshalMsg
//        |-----> (*) CanUnmarshalMsg
//        |-----> (*) Msgsize
//        |-----> (*) MsgIsZero
//
// Dil2PublicKey
//       |-----> (*) MarshalMsg
//       |-----> (*) CanMarshalMsg
//       |-----> (*) UnmarshalMsg
//       |-----> (*) CanUnmarshalMsg
//       |-----> (*) Msgsize
//       |-----> (*) MsgIsZero
//
// Dil2Signature
//       |-----> (*) MarshalMsg
//       |-----> (*) CanMarshalMsg
//       |-----> (*) UnmarshalMsg
//       |-----> (*) CanUnmarshalMsg
//       |-----> (*) Msgsize
//       |-----> (*) MsgIsZero
//
// DilithiumKeyPair
//         |-----> (*) MarshalMsg
//         |-----> (*) CanMarshalMsg
//         |-----> (*) UnmarshalMsg
//         |-----> (*) CanUnmarshalMsg
//         |-----> (*) Msgsize
//         |-----> (*) MsgIsZero
//

// MarshalMsg implements msgp.Marshaler
func (z *Dil2PrivateKey) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendBytes(o, (*z)[:])
	return
}

func (_ *Dil2PrivateKey) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*Dil2PrivateKey)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Dil2PrivateKey) UnmarshalMsg(bts []byte) (o []byte, err error) {
	bts, err = msgp.ReadExactBytes(bts, (*z)[:])
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	o = bts
	return
}

func (_ *Dil2PrivateKey) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*Dil2PrivateKey)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Dil2PrivateKey) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize + (2528 * (msgp.ByteSize))
	return
}

// MsgIsZero returns whether this is a zero value
func (z *Dil2PrivateKey) MsgIsZero() bool {
	return (*z) == (Dil2PrivateKey{})
}

// MarshalMsg implements msgp.Marshaler
func (z *Dil2PublicKey) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendBytes(o, (*z)[:])
	return
}

func (_ *Dil2PublicKey) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*Dil2PublicKey)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Dil2PublicKey) UnmarshalMsg(bts []byte) (o []byte, err error) {
	bts, err = msgp.ReadExactBytes(bts, (*z)[:])
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	o = bts
	return
}

func (_ *Dil2PublicKey) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*Dil2PublicKey)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Dil2PublicKey) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize + (1312 * (msgp.ByteSize))
	return
}

// MsgIsZero returns whether this is a zero value
func (z *Dil2PublicKey) MsgIsZero() bool {
	return (*z) == (Dil2PublicKey{})
}

// MarshalMsg implements msgp.Marshaler
func (z *Dil2Signature) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendBytes(o, (*z)[:])
	return
}

func (_ *Dil2Signature) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*Dil2Signature)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Dil2Signature) UnmarshalMsg(bts []byte) (o []byte, err error) {
	bts, err = msgp.ReadExactBytes(bts, (*z)[:])
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	o = bts
	return
}

func (_ *Dil2Signature) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*Dil2Signature)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Dil2Signature) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize + (2420 * (msgp.ByteSize))
	return
}

// MsgIsZero returns whether this is a zero value
func (z *Dil2Signature) MsgIsZero() bool {
	return (*z) == (Dil2Signature{})
}

// MarshalMsg implements msgp.Marshaler
func (z *DilithiumKeyPair) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0003Len := uint32(2)
	var zb0003Mask uint8 /* 3 bits */
	if (*z).PublicKey == (Dil2PublicKey{}) {
		zb0003Len--
		zb0003Mask |= 0x2
	}
	if (*z).SecretKey == (Dil2PrivateKey{}) {
		zb0003Len--
		zb0003Mask |= 0x4
	}
	// variable map header, size zb0003Len
	o = append(o, 0x80|uint8(zb0003Len))
	if zb0003Len != 0 {
		if (zb0003Mask & 0x2) == 0 { // if not empty
			// string "pk"
			o = append(o, 0xa2, 0x70, 0x6b)
			o = msgp.AppendBytes(o, ((*z).PublicKey)[:])
		}
		if (zb0003Mask & 0x4) == 0 { // if not empty
			// string "sk"
			o = append(o, 0xa2, 0x73, 0x6b)
			o = msgp.AppendBytes(o, ((*z).SecretKey)[:])
		}
	}
	return
}

func (_ *DilithiumKeyPair) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*DilithiumKeyPair)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *DilithiumKeyPair) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0003 int
	var zb0004 bool
	zb0003, zb0004, bts, err = msgp.ReadMapHeaderBytes(bts)
	if _, ok := err.(msgp.TypeError); ok {
		zb0003, zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0003 > 0 {
			zb0003--
			bts, err = msgp.ReadExactBytes(bts, ((*z).SecretKey)[:])
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "SecretKey")
				return
			}
		}
		if zb0003 > 0 {
			zb0003--
			bts, err = msgp.ReadExactBytes(bts, ((*z).PublicKey)[:])
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "PublicKey")
				return
			}
		}
		if zb0003 > 0 {
			err = msgp.ErrTooManyArrayFields(zb0003)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array")
				return
			}
		}
	} else {
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0004 {
			(*z) = DilithiumKeyPair{}
		}
		for zb0003 > 0 {
			zb0003--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "sk":
				bts, err = msgp.ReadExactBytes(bts, ((*z).SecretKey)[:])
				if err != nil {
					err = msgp.WrapError(err, "SecretKey")
					return
				}
			case "pk":
				bts, err = msgp.ReadExactBytes(bts, ((*z).PublicKey)[:])
				if err != nil {
					err = msgp.WrapError(err, "PublicKey")
					return
				}
			default:
				err = msgp.ErrNoField(string(field))
				if err != nil {
					err = msgp.WrapError(err)
					return
				}
			}
		}
	}
	o = bts
	return
}

func (_ *DilithiumKeyPair) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*DilithiumKeyPair)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *DilithiumKeyPair) Msgsize() (s int) {
	s = 1 + 3 + msgp.ArrayHeaderSize + (2528 * (msgp.ByteSize)) + 3 + msgp.ArrayHeaderSize + (1312 * (msgp.ByteSize))
	return
}

// MsgIsZero returns whether this is a zero value
func (z *DilithiumKeyPair) MsgIsZero() bool {
	return ((*z).SecretKey == (Dil2PrivateKey{})) && ((*z).PublicKey == (Dil2PublicKey{}))
}