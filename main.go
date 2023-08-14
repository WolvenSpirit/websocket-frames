/**
*	For education purpose only, by Mihai Dragusin
* 	Do NOT copy
 */

package ws

import (
	"encoding/binary"
	"fmt"
)

// Parses the websocket masked frame decoding it using
//
//	`D_i = E_i XOR M_(i mod 4)` or `returnData[i] = encodedData[i] ^ mask32bit[i%4]`
func ParseFrame(p []byte) Frame {
	firstByte := p[0]
	fin := p[0]&FIN == FIN
	secondByte := p[1]
	maskSet := secondByte&HasMask == HasMask
	// fmt.Println("mask set is ", maskSet)
	msgLen := uint64(secondByte & 0b01111111)
	startOfPayload := 2 // after first two bytes we already parsed

	decoded := make([]byte, msgLen+1)
	masks := make([]byte, 4)

	if msgLen < 126 {

		// fmt.Println("msg len is ", msgLen)
		if maskSet {
			masks = p[2:6] // 2 - 5 inclusive - 4 bytes, 32 bit
		}
		// fmt.Println("mask is ", masks)
		startOfPayload += len(masks)

	} else if msgLen == 126 {
		// fmt.Println("second byte len is", msgLen)
		// if 126 then the actual msg len is next 2 bytes as 16 bit uint
		msgLen = uint64(binary.BigEndian.Uint16(p[2:4])) // network data is always in big endian
		if maskSet {
			masks = p[4:8]
		}
		startOfPayload += (len(masks) + len(p[2:4])) // start of payload data should be after mask and length bytes, at the first byte after
		decoded = make([]byte, msgLen+1)

	} else if msgLen > 126 {
		// if over 126 then we need to parse the next eight bytes as an unsigned integer 64 bit
		msgLen := binary.BigEndian.Uint64(p[2:10])
		if maskSet {
			masks = p[10:14]
		}
		startOfPayload += (len(masks) + len(p[2:10])) + 1
		decoded = make([]byte, msgLen+1)

	}

	for i := 0; i < int(msgLen); i++ {
		decoded[i] = p[startOfPayload+i] ^ masks[i%4]
	}

	return Frame{
		Byte0: Byte0{
			FIN:  fin,
			RSV1: (firstByte&RSV1 != RSV1),
			RSV2: (firstByte&RSV2 != RSV2),
			RSV3: (firstByte&RSV3 != RSV3),
		},
		Byte1: Byte1{
			HasMask:       maskSet,
			InitialLength: (secondByte & 0b01111111),
		},
		DecodedPayload: decoded,
		PayloadLength:  msgLen,
	}
}

// SendNewFrame assembles an unmasked server frame with data as payload
func SendSerializedFrame(data []byte) []byte {
	var payload []byte
	payloadLen := len(data)
	if payloadLen < 126 {
		payload = make([]byte, payloadLen+2)
		payload[0] ^= FIN
		payload[0] ^= OPCODE_TEXT
		// FIN and OPCODE_TEXT combination informs it is a full frame with text
		payload[1] ^= uint8(payloadLen)
		offset := 2
		for i := 0; i < payloadLen; i++ {
			payload[offset+i] = data[i]
		}
	}
	fmt.Println("sending", string(payload))
	return payload
}
