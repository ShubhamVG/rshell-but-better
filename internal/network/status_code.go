package network

type StatusCode = uint8

// TODO: Improve these maybe
const (
	PING              StatusCode = 10
	REQUESTING_CLOSE  StatusCode = 20
	REDIRECT          StatusCode = 30
	EXECUTE           StatusCode = 40
	OUTPUT            StatusCode = 50
	OUTPUT_WITH_ERROR StatusCode = 55
)
