package sulfur

type AuthRequest struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}

type AuthResponse struct {
	*ErrorResponse

	Token  string `json:"token"`
	Record struct {
		CollectionId    string `json:"collectionId"`
		CollectionName  string `json:"collectionName"`
		Id              string `json:"id"`
		Email           string `json:"email"`
		EmailVisibility bool   `json:"emailVisibility"`
		Verified        bool   `json:"verified"`
		Name            string `json:"name"`
		Avatar          string `json:"avatar"`
		Created         string `json:"created"`
		Updated         string `json:"updated"`
	} `json:"record"`
}
