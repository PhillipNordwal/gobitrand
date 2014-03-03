package gobitrand
import (
	"testing"
)

// TestRefillZerosCount checks to make user used is zero after refill
// is called.
func TestRefillZerosCount(t *testing.T) {
	refill()
	if used != 0 {
		t.Error("The used variable wasn't zerod on refill.")
	}
}

// TestTwo_bits_increasesCount checks to make sure that used goes up
// by two when two bits are requested
func TestTwo_bits_increasesCount(t *testing.T) {
	refill()
	Two_bits()
	if used != 2 {
		t.Error("The used variable didn't increase on use.")
	}
}

// TestTwo_bits_cyclesCount checks to make sure that after the
// entropy pool is exhausted that usde gets set to two
func TestTwo_bits_cyclesCount(t *testing.T) {
	refill()
	for i:=0; i<17; i++ {
		Two_bits()
	}
	if used != 2 {
		t.Error("The used variable didn't wrap on use.")
	}
}

// TestTwo_bits_successive sets the entropy pool and checks that
// successive calls return the proper number
func TestTwo_bits_successive(t *testing.T) {
	used = 0
	// should be to 0b11100100000110111110010000011011
	// or 4#3210012332100123
	src = 0xE41BE41B
	res := []uint8{3,2,1,0,0,1,2,3,3,2,1,0,0,1,2,3}
	for i:=0; i<16; i++ {
		if res[i] != Two_bits(){
			t.Error("Two_bits isn't properly returning src.")
		}
	}
}
