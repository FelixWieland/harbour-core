package harbourcore

//response for an unauthenticated user
type respForbidden struct {
	Code        int    `json:"code"`
	Cessage     string `json:"message"`
	Description string `json:"description"`
}
