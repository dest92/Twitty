package models

//Contains the token to return to the client

type ResponseLogin struct {
	Token string `json:"token,omitempty"`
}