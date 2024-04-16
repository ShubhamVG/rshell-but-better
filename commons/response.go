package commons

type Response struct {
	UniqueAddr UniqueConnAddr
	Status     StatusCode
	Body       string
}

func ParseIntoResponse(
	uniqueAddr UniqueConnAddr,
	contentBuffer []byte,
) Response {
	statusCode := contentBuffer[0]

	if len(contentBuffer) > 1 {
		content := string(contentBuffer[1:])
		return Response{UniqueAddr: uniqueAddr, Status: statusCode, Body: content}
	}

	return Response{UniqueAddr: uniqueAddr, Status: statusCode, Body: ""}
}
