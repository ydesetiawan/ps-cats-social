package dto

type RegisterResp struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	AccessToken string `json:"accessToken"`
}
