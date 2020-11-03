package jpush

type Audience struct {
	Tag    []string
	TagAnd []string
	Alias  []string
	IDS    []string
}

type Obj struct {
	Type    string
	Notice  Notice
	Message Message
}

type Notice struct {
	Alert string
}

type Message struct {
	Title   string
	Content string
}

type RespFail struct {
	Error RespFailError
	MsgId string
}

type RespFailError struct {
	Code    int
	Message string
}
