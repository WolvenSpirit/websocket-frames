/**
*	For education purpose only, by Mihai Dragusin
* 	Do NOT copy
 */

package bit_util

import (
	"strconv"
)

// Dumb way of getting a mask at bit position
func GetPositiveMaskForBitPosition(bit int) uint8 {
	var bin string = "1"
	for i := bit; i > 1; i-- {
		bin += "0"
	}
	binI64, _ := strconv.ParseInt(bin, 2, 64)
	return uint8(binI64)
}

// Dumb way of getting a mask for bit range
func GetPositiveMaskForBitRange(bitLeft, bitRight int) uint8 {
	var bin string = "1"
	for i := bitLeft; i > 1; i-- {
		if i > bitRight {
			bin += "1"
		} else {
			bin += "0"
		}
	}
	binI64, _ := strconv.ParseInt(bin, 2, 8) // base 2, 8 bit
	return uint8(binI64)
}

// Flip the bit at mask bit position
/*
func main() {
	Byte := uint8(0x80)
	fmt.Printf("%b\n", Xor(&Byte, getPositiveMaskForBitPosition(4))) // Equivalent to fmt.Printf("%b\n", Byte^8)
}
*/
func Xor(b *byte, mask uint8) byte {
	return (*b) ^ mask
}

// Check if bit was flipped on
/*
 // helper for something like
 myByte := x07F
 is128 := (myByte & 128) == 128 // would be false since 128 is 0x80 and mask wasn't applied
 // can be written as
 isMaskSet(myByte,128)
*/
func IsMaskSet(b *byte, mask uint8) bool {
	return ((*b) & mask) == mask
}
