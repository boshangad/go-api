package helpers

import "strconv"

type Int int

type Int8 int8

type Int16 int16

type Int32 int32

type Int64 int64

type Uint uint

type Uint8 uint8

type Uint16 uint16

type Uint32 uint32

type Uint64 uint64

func (u *Int) UnmarshalJSON(bs []byte) (err error) {
	var v int64
	if v, err = unmarshalInt64(bs); err != nil {
		return err
	}
	*u = Int(v)
	return err
}

func (u *Int8) UnmarshalJSON(bs []byte) (err error) {
	var v int64
	if v, err = unmarshalInt64(bs); err != nil {
		return err
	}
	*u = Int8(v)
	return err
}

func (u *Int16) UnmarshalJSON(bs []byte) (err error) {
	var v int64
	if v, err = unmarshalInt64(bs); err != nil {
		return err
	}
	*u = Int16(v)
	return err
}

func (u *Int32) UnmarshalJSON(bs []byte) (err error) {
	var v int64
	if v, err = unmarshalInt64(bs); err != nil {
		return err
	}
	*u = Int32(v)
	return err
}

func (u *Int64) UnmarshalJSON(bs []byte) (err error) {
	var v int64
	if v, err = unmarshalInt64(bs); err != nil {
		return err
	}
	*u = Int64(v)
	return err
}

func (u *Uint) UnmarshalJSON(bs []byte) (err error) {
	var v uint64
	if v, err = unmarshalUint64(bs); err != nil {
		return err
	}
	*u = Uint(v)
	return nil
}

func (u *Uint8) UnmarshalJSON(bs []byte) (err error) {
	var v uint64
	if v, err = unmarshalUint64(bs); err != nil {
		return err
	}
	*u = Uint8(v)
	return nil
}

func (u *Uint16) UnmarshalJSON(bs []byte) (err error) {
	var v uint64
	if v, err = unmarshalUint64(bs); err != nil {
		return err
	}
	*u = Uint16(v)
	return nil
}

func (u *Uint32) UnmarshalJSON(bs []byte) (err error) {
	var v uint64
	if v, err = unmarshalUint64(bs); err != nil {
		return err
	}
	*u = Uint32(v)
	return nil
}

func (u *Uint64) UnmarshalJSON(bs []byte) (err error) {
	var v uint64
	if v, err = unmarshalUint64(bs); err != nil {
		return err
	}
	*u = Uint64(v)
	return nil
}

// String returns a string representation of
func unmarshalUint64(bs []byte) (uint64, error) {
	// Parse plain numbers directly.
	str := string(bs)
	if bs[0] == '"' && bs[len(bs)-1] == '"' {
		// Unwrap the quotes from string numbers.
		str = string(bs[1 : len(bs)-1])
	}
	x, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return x, nil
}

// String returns a string representation of
func unmarshalInt64(bs []byte) (int64, error) {
	// Parse plain numbers directly.
	str := string(bs)
	if bs[0] == '"' && bs[len(bs)-1] == '"' {
		// Unwrap the quotes from string numbers.
		str = string(bs[1 : len(bs)-1])
	}
	x, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return x, nil
}
