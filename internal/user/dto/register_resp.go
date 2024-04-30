package dto

type RegisterResp struct {
	Message string   `json:"message"`
	Data    struct{} `json:"data"`
}

type Data struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	AccessToken string `json:"accessToken"`
}
