package server

type Request struct {
	Addr          UniqueConnAddr
	ContentBuffer []byte
}

func parseIntoRequest(
	uniqueAddr UniqueConnAddr,
	contentBuffer []byte,
) Request {
	addr := 
}