package datastore

import (
	"github.com/upper/db/v4"
	"time"
)

type (
	DatastoreT struct {
		db.Session
	}
	//User struct {
	//	ID       int             `json:"id" db:"id,omitempty"`
	//	Username string          `json:"username" db:"username"`
	//	Balance  decimal.Decimal `json:"balance" db:"balance"`
	//	//Wallets  WalletT         `json:"wallets" db:"wallets"`
	//	Created time.Time `json:"created" db:"created"`
	//}
	User struct {
		ID       int       `json:"id" db:"id,omitempty"`
		Name     string    `json:"name" db:"name"`
		Email    string    `json:"email" db:"email"`
		Phone    string    `json:"phone" db:"phone"`
		Password string    `json:"password" db:"password"`
		Created  time.Time `json:"created" db:"created"`
	}
	Image struct {
		ID      int       `json:"id" db:"id,omitempty"`
		OwnerID int       `json:"owner_id" db:"owner_id"`
		Name    string    `json:"name" db:"name"`
		Created time.Time `json:"created" db:"created"`
	}
	Favorite struct {
		UserID  int       `json:"user_id" db:"user_id,omitempty"`
		ImageID int       `json:"image_id" db:"image_id,omitempty"`
		Created time.Time `json:"created" db:"created"`
	}
	File struct {
		ID   int    `json:"id" db:"id,omitempty"`
		Size string `json:"size" db:"size"`
		Path string `json:"path" db:"path"`
	}

	//StateData struct {
	//	//Ad AdT `json:"ad"`
	//	//WithdrawalMethod WithdrawalMethodT `json:"withdrawal_method"`
	//	//Amount           decimal.Decimal   `json:"amount"`
	//	//Number           string            `json:"number"`
	//	//Count            int               `json:"count"` // -1 - All
	//	Text     string                   `json:"text"`
	//	Tags     []string                 `json:"tags"`
	//	Page     int                      `json:"page"`
	//	Entities []tgbotapi.MessageEntity `json:"entities"`
	//	*postgresql.JSONBConverter
	//}
	//
	//Job struct {
	//	ID      int      `json:"id" db:"id,omitempty"`
	//	OwnerID int      `json:"owner_id" db:"owner_id"`
	//	Tags    []string `json:"tags" db:"tags"`
	//	Stage   StageT   `json:"stage" db:"stage"`
	//	Data    JobData  `json:"job_data" db:"job_data"`
	//	//Wallets  WalletT         `json:"wallets" db:"wallets"`
	//	Created time.Time `json:"created" db:"created"`
	//}
	//JobData struct {
	//	//Ad AdT `json:"ad"`
	//	//WithdrawalMethod WithdrawalMethodT `json:"withdrawal_method"`
	//	//Amount           decimal.Decimal   `json:"amount"`
	//	//Number           string            `json:"number"`
	//	//Count            int               `json:"count"` // -1 - All
	//	Text string `json:"text"`
	//	//Tags     []string                 `json:"tags"`
	//	Entities []tgbotapi.MessageEntity `json:"entities"`
	//	*postgresql.JSONBConverter
	//}
	//Payment struct {
	//	ID      int       `json:"id" db:"id,omitempty"`
	//	JobID   int       `json:"job_id" db:"job_id,omitempty"`
	//	Amount  int       `json:"amount" db:"amount"`
	//	Status  StatusT   `json:"status"  db:"status"`
	//	Created time.Time `json:"created" db:"created"`
	//}
	//AdvUser struct {
	//	ID       int             `json:"id" db:"id,omitempty"`
	//	Username string          `json:"username" db:"username"`
	//	Balance  decimal.Decimal `json:"balance" db:"balance"`
	//	Token    uuid.UUID       `json:"token" db:"token"`
	//	Created  time.Time       `json:"created" db:"created"`
	//}
	////WalletT struct {
	////	Card string `json:"card"`
	////	YM   string `json:"ym"`
	////	WM   string `json:"wm"`
	////	Qiwi string `json:"qiwi"`
	////	*postgresql.JSONBConverter
	////}
	//Bot struct {
	//	ID       int    `json:"id" db:"id,omitempty"`
	//	Token    string `json:"token" db:"token"`
	//	ApiKey   string `json:"api_key" db:"api_key"`
	//	Username string `json:"username" db:"username"`
	//	Name     string `json:"name" db:"name"`
	//	OwnerID  int    `json:"owner_id" db:"owner_id"`
	//}
	//BotUser struct {
	//	UserID   int       `json:"user_id" db:"user_id,omitempty"`
	//	BotID    int       `json:"bot_id" db:"bot_id,omitempty"`
	//	IsActive bool      `json:"is_active" db:"is_active"`
	//	LastMsg  time.Time `json:"last_msg" db:"last_msg"`
	//}
	//TelegramUser struct {
	//	ID       int       `json:"id" db:"id,omitempty"`
	//	IsActive bool      `json:"is_active" db:"is_active"`
	//	LastMsg  time.Time `json:"last_msg" db:"last_msg"`
	//}
	//Stat struct {
	//	ID          int       `json:"id" db:"id,omitempty"`
	//	Users       int       `json:"users" db:"users"`
	//	ActiveUsers int       `json:"active_users" db:"active_users"`
	//	Uses        int       `json:"uses" db:"uses"`
	//	Created     time.Time `json:"created" db:"created"`
	//}
	//Transaction struct {
	//	ID      int             `json:"id" db:"id,omitempty"`
	//	BotID   int             `json:"bot_id" db:"bot_id"`
	//	Amount  decimal.Decimal `json:"amount" db:"amount"`
	//	T       DistributionT   `json:"t" db:"t"`
	//	Users   int             `json:"users" db:"users"`
	//	Created time.Time       `json:"created" db:"created"`
	//}
	//Withdraw struct {
	//	ID      int             `json:"id" db:"id,omitempty"`
	//	UserID  int             `json:"user_id" db:"user_id"`
	//	Amount  decimal.Decimal `json:"amount" db:"amount"`
	//	Method  MethodT         `json:"method"  db:"method"`
	//	Comment string          `json:"comment" db:"comment"`
	//	Created time.Time       `json:"created" db:"created"`
	//}
	//AdT struct {
	//	ID      int           `json:"id" db:"id,omitempty"`
	//	OwnerID int           `json:"owner_id" db:"owner_id"`
	//	Count   int           `json:"count" db:"count"` // -1 - All
	//	T       DistributionT `json:"t" db:"t"`
	//	Text    string        `json:"text" db:"text"`
	//	Img     string        `json:"img" db:"img"`
	//	//Buttons      postgresql.JSONBArray `json:"buttons" db:"buttons"`
	//	Amount  decimal.Decimal `json:"amount" db:"amount"`
	//	Invoice decimal.Decimal `json:"invoice" db:"invoice"`
	//	Stage   StageT          `json:"stage" db:"stage"`
	//	Created time.Time       `json:"created" db:"created"`
	//}
	//Button struct {
	//	Text string `json:"text"`
	//	Link string `json:"link"`
	//	*postgresql.JSONBConverter
	//}
	//Payment struct {
	//	ID      int             `json:"id" db:"id,omitempty"`
	//	AdID    int             `json:"ad_id" db:"ad_id,omitempty"`
	//	Amount  decimal.Decimal `json:"amount" db:"amount"`
	//	Method  MethodT         `json:"method"  db:"method"`
	//	Created time.Time       `json:"created" db:"created"`
	//}
)
