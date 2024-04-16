package commons

// Design it first goddamnit
type Request struct {
	UniqueAddr UniqueConnAddr
	Status     StatusCode
	Payload    string
}

func ParseIntoRequest(buffer []byte) Request {
	statusCode := buffer[0]

	if len(buffer) == 1 {
		return Request{UniqueAddr: "", Status: statusCode, Payload: ""}
	}

	payload := string(buffer[1:])
	return Request{UniqueAddr: "", Status: statusCode, Payload: payload}
}
