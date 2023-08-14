package ws

/**
*	Copyright: Only for preview, by Mihai Dragusin
 */

/*
References:

https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API/Writing_WebSocket_server
https://datatracker.ietf.org/doc/html/rfc6455#section-5.2
*/

const (
	ContinuationFrame = 0x0
	TextFrame         = 0b_0000_0001 // 0x1
	BinaryFrame       = 0b_0000_0010 // 0x2
	ConnectionClose   = 0b_0000_1000 // 0x8
	Ping              = 0x9
	Pong              = 0xA
	Payload126        = 0x7E // Next 2 bytes as uint16 are length
	PayloadContinued  = 0x7F // Next 8 bytes are the length as uint64

)

const (
	FIN                       uint8 = 0b_1000_0000 // bit_util.GetPositiveMaskForBitPosition(8)
	RSV1                      uint8 = 0b_0100_0000 // bit_util.GetPositiveMaskForBitPosition(7)
	RSV2                      uint8 = 0b_0010_0000 // bit_util.GetPositiveMaskForBitPosition(6)
	RSV3                      uint8 = 0b_0001_0000 // bit_util.GetPositiveMaskForBitPosition(5)
	HasMask                   uint8 = 0b_1000_0000
	OPCODE_TEXT               uint8 = 0b_0000_0001 // text data opcode
	OPCODE_BINARY             uint8 = 0b_0000_0010 // text data opcode
	OPCODE_CONTINUATION_FRAME uint8 = 0b_0000_0000
)
