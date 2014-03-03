// Package gobitrand provides a method for consuming only some of the
// bits of entropy provided by math/rand
package gobitrand

import (
	"math/rand"
)

var src uint32
var used uint8 = 32

// refill refills the internal 32bit entropy buffer, discarding the
// remaining bits in it.
func refill() {
	src = rand.Uint32()
	used = 0
}

// two_bit returns two bits from the entropy buffer and refills it if
// necessaray.
func Two_bits() uint8 {

	if used >= 32 - 1 {
		refill()
	}
	// select the least two bits from the entropy pool after
	// it has been shifted past the used bits
	bits := (uint8) (3 & (src>>used))
	used +=2
	return bits
}
