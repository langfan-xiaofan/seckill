package dto

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResp struct{}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginResp struct {
	Token string `json:"token"`
}
