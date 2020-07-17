package utils

import (
	"math"
	"math/rand"
	"time"
)

const MULTI int64 = 1024 * 1024

var MAP = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
	'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w',
	'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
}

func indexOf(candidate rune) int {
	for index, c := range MAP {
		if c == candidate {
			return index
		}
	}
	return -1
}

func GenerateID() string {
	now := time.Now()
	calc := now.Unix() * MULTI
	randsrc := rand.NewSource(now.UnixNano())
	randr := rand.New(randsrc)
	calc += randr.Int63n(MULTI)
	s := ""
	for calc != 0 {
		s += string(MAP[int(calc%int64(len(MAP)))])
		calc /= int64(len(MAP))
	}
	randsrc = rand.NewSource(now.UnixNano())
	randr = rand.New(randsrc)
	s += string(MAP[randr.Intn(len(MAP))]) + string(MAP[randr.Intn(len(MAP))])
	return s
}

func TimeFromID(s string) *time.Time {
	var calc int64 = 0
	rslc := []rune(s[:len(s)-2])
	for index, val := range rslc {
		calc += int64(indexOf(val)) * int64(math.Pow(float64(len(MAP)), float64(index)))
	}
	t := time.Unix((calc-(calc%MULTI))/MULTI, 0)
	return &t
}
