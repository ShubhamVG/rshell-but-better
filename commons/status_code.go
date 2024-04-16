package commons

type StatusCode = uint8

const (
	PING             StatusCode = 10
	REQUESTING_CLOSE StatusCode = 20
	REDIRECT         StatusCode = 30
	EXECUTE          StatusCode = 40
)
