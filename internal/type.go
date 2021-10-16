package main

import (
	"fmt"
	"strings"
)

//go:generate stringer -type=GenType -trimprefix=GenType
type GenType int

const (
	GenTypeInvalid GenType = iota
	GenTypeInt8
	GenTypeInt16
	GenTypeInt32
	GenTypeInt
	GenTypeInt64
	GenTypeUint8
	GenTypeUint16
	GenTypeUint32
	GenTypeUint
	GenTypeUint64
	GenTypeFloat32
	GenTypeFloat64
	GenTypeString
	GenTypeTime
	GenTypeEndTypes
)

// Data return data type of GenType
func (t GenType) Data() string {
	switch {
	case t < GenTypeTime:
		return fmt.Sprintf("map[%s]struct{}", t.StringLower())
	case t == GenTypeTime:
		return "map[int64]time.Time"
	}

	return "unknown"
}

func (t GenType) StringLower() string {
	return strings.ToLower(t.String())
}

func (t GenType) Type() string {
	switch {
	case t < GenTypeTime:
		return t.StringLower()
	case t == GenTypeTime:
		return "time.Time"
	}

	return "unknown"
}

func (t GenType) MapKey(k string) string {
	switch {
	case t < GenTypeTime:
		return k
	case t == GenTypeTime:
		return fmt.Sprintf("%s.Unix()", k)
	}

	return "unknown"
}

func (t GenType) MapValue(v string) string {
	switch {
	case t < GenTypeTime:
		return "struct{}{}"
	case t == GenTypeTime:
		return v
	}

	return "unknown"
}
