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

type Float32 float32

type Float64 float64

func (u *Int) UnmarshalJSON(bs []byte) (err error) {
	var v int64
	if v, err = unmarshalInt64(bs); err != nil {
		return err
	}
	*u = Int(v)
	return err
}

func (u *Int) UnmarshalYAML(bs []byte) (err error) {
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

func (u *Int8) UnmarshalYAML(bs []byte) (err error) {
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

func (u *Int16) UnmarshalYAML(bs []byte) (err error) {
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

func (u *Int32) UnmarshalYAML(bs []byte) (err error) {
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

func (u *Int64) UnmarshalYAML(bs []byte) (err error) {
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

func (u *Uint) UnmarshalYAML(bs []byte) (err error) {
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

func (u *Uint8) UnmarshalYAML(bs []byte) (err error) {
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

func (u *Uint16) UnmarshalYAML(bs []byte) (err error) {
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

func (u *Uint32) UnmarshalYAML(bs []byte) (err error) {
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

func (u *Uint64) UnmarshalYAML(bs []byte) (err error) {
	var v uint64
	if v, err = unmarshalUint64(bs); err != nil {
		return err
	}
	*u = Uint64(v)
	return nil
}

func (u *Float32) UnmarshalJSON(bs []byte) (err error) {
	var v float32
	if v, err = unmarshalFloat32(bs); err != nil {
		return err
	}
	*u = Float32(v)
	return nil
}

func (u *Float32) UnmarshalYAML(bs []byte) (err error) {
	var v float32
	if v, err = unmarshalFloat32(bs); err != nil {
		return err
	}
	*u = Float32(v)
	return nil
}

func (u *Float64) UnmarshalJSON(bs []byte) (err error) {
	var v float64
	if v, err = unmarshalFloat64(bs); err != nil {
		return err
	}
	*u = Float64(v)
	return nil
}

func (u *Float64) UnmarshalYAML(bs []byte) (err error) {
	var v float64
	if v, err = unmarshalFloat64(bs); err != nil {
		return err
	}
	*u = Float64(v)
	return nil
}

func unmarshalUint64(bs []byte) (uint64, error) {
	str := string(bs)
	if bs[0] == '"' && bs[len(bs)-1] == '"' {
		str = string(bs[1 : len(bs)-1])
	}
	x, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return x, nil
}

func unmarshalInt64(bs []byte) (int64, error) {
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

func unmarshalFloat32(bs []byte) (float32, error) {
	str := string(bs)
	if bs[0] == '"' && bs[len(bs)-1] == '"' {
		str = string(bs[1 : len(bs)-1])
	}
	x, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0, err
	}
	return float32(x), nil
}

func unmarshalFloat64(bs []byte) (float64, error) {
	str := string(bs)
	if bs[0] == '"' && bs[len(bs)-1] == '"' {
		str = string(bs[1 : len(bs)-1])
	}
	x, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}
	return x, nil
}
