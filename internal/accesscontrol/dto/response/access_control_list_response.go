package response

type AccessControlResponse struct {
	Reminder AllowedMethods `json:"reminder"`
}

type AllowedMethods struct {
	Create bool `json:"create"`
	Read   bool `json:"read"`
	Update bool `json:"update"`
	Delete bool `json:"delete"`
}
