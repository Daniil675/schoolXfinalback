package server

type (
	Server struct {
		//Sessions  *session.SessionStorage
		//API       *api.API
		//TG        *telegram.Bot
		//Bot       *vk.Bot
	}

	Err struct {
		Code ErrorCode `json:"error_code"`
	}
)
