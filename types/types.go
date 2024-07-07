package types

type Quote struct {
	ID         string `json:"_id"`
	Source     string `json:"source"`
	Philosophy string `json:"philosophy"`
	Quote      string `json:"quote"`
}
