package response

type UserData struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
}

type UserDetails struct {
	Name              string
	Email             string
	Mobile            string
	IsBlocked         bool
	BlockedAt         string
	BlockedBy         uint
	ReasonForBlocking string
}
