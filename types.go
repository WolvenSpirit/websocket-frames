package ws

// Byte0 is the parsed first byte from the payload
type Byte0 struct {
	FIN  bool
	RSV1 bool
	RSV2 bool
	RSV3 bool
}

// Byte1 is the parsed second byte from the payload
type Byte1 struct {
	HasMask       bool
	InitialLength uint8
}

// Frame represents the base websocket frame
type Frame struct {
	// Byte0 is the parsed first byte from the payload
	Byte0 Byte0
	// Byte1 is the parsed second byte from the payload
	Byte1 Byte1
	// The payload, may be masked as indicated by Byte1.HasMask
	Payload []byte
	// The actual payload length
	PayloadLength uint64
	// The decoded payload
	DecodedPayload []byte
	// Flags if the DecodedPayload was parsed and is valid
	Decoded bool
}
