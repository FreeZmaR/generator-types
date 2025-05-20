package main

import (
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

const quote = '"'

func B2S(b []byte) (string, error) {
	if len(b) == 0 {
		return "", nil
	}

	if b[0] == quote {
		b = b[1:]
	}

	if b[len(b)-1] == quote {
		b = b[:len(b)-1]
	}

	return unsafe.String(unsafe.SliceData(b), len(b)), nil
}

func B2I(b []byte) (int, error) {
	v, _ := B2S(b)
	if v == "" {
		return 0, nil
	}

	return strconv.Atoi(v)
}

func B2Bool(b []byte) (bool, error) {
	v, _ := B2S(b)
	if v == "" {
		return false, nil
	}

	val := strings.ToLower(v)

	switch val {
	case "true", "1":
		return true, nil
	case "false", "0":
		return false, nil
	default:
		return false, fmt.Errorf("invalid boolean value: %s", v)
	}
}

func B2Float64(b []byte) (float64, error) {
	v, _ := B2S(b)
	if v == "" {
		return 0, nil
	}

	return strconv.ParseFloat(v, 64)
}
