package service

const (
	MissingNameParam Code = 100001
)

var codeText = map[Code]string{
	MissingNameParam: "Missing name parameter",
}

func NewError(code Code) error {
	return NewRPCError(code, codeText[code])
}
