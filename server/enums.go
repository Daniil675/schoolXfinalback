package server

type ErrorCode int

const (
	SomethingWentWrong = iota
	InvalidData
	AccessDenied
)

var AllowableFileExtensions = []string{"png", "jpeg", "jpg", "gif", "bmp", "doc", "rtf", "pdf", "txt"}
