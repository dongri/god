package barikata

import (
	"fmt"
	"math"
	"strconv"
)

// major: 16 bit unsigned integer (0~65535)
// minor: 16 bit unsigned integer (0~65535)

// Beacon Bit
const (
	BitTotal = 32
)

// PackIBeacon is pack major, minor values
// params x, y is sequence value (24bit, 8bit)
// returns major, minor is real beacon value (16bit, 16bit)
func PackBeacon(x, y int, minorBit uint) (int, int) {
	s := fmt.Sprintf("%024b%08b", x, y)
	sum, _ := strconv.ParseInt(s, 2, 0)
	major := int(sum >> (minorBit + uint(BitTotal/2-minorBit)))
	minor := int(sum) & int(math.Pow(2, BitTotal/2)-1)
	return major, minor
}

// UnpackIBeacon is unpack major, minor values
// params x, y is real beacon value (16bit, 16bit)
// returns major, minor is sequence value (24bit, 8bit)
func UnpackBeacon(x, y int, minorBit uint) (int, int) {
	s := fmt.Sprintf("%016b%016b", x, y)
	sum, _ := strconv.ParseInt(s, 2, 0)
	major := int(sum) >> minorBit
	minor := int(sum) & int(math.Pow(2, float64(minorBit))-1)
	return major, minor
}
