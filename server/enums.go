package server

type ErrorCode int

const (
	SomethingWentWrong = iota
	InvalidData
	AccessDenied
	UnallowableFileExtension
)

var (
	AllowableFileExtensions = []string{".png", ".jpeg", ".jpg"}
	uploadPath              = "./resources/temp/"
	sizes                   = []ResizeConfig{
		{
			Name:   "HD",
			Width:  1280,
			Height: 720,
		},
		{
			Name:   "FHD",
			Width:  1920,
			Height: 1080,
		},
		{
			Name:   "2K",
			Width:  2560,
			Height: 1440,
		},
		{
			Name:   "4K",
			Width:  3840,
			Height: 2160,
		},
	}
)
