package token

import "time"

//maker is an interface for managing tokens
type Maker interface {
	//createtoken creates a token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	//Verifytoken checks if the input token is valid or not
	VerifyToken(token string) (*Payload, error)
}
