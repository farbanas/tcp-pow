package models

type Quote struct {
	QuoteText      string `json:"q"`
	AuthorName     string `json:"a"`
	AuthorImage    string `json:"i"`
	CharacterCount string `json:"c"`
	HTMLQuote      string `json:"h"`
}

type TCPData struct {
	Address     string      `json:"address"`
	PackageType PackageType `json:"package_type"`
	Nonce       string      `json:"nonce,omitempty"`
	Difficulty  int         `json:"difficulty,omitempty"`
	Prefix      string      `json:"prefix,omitempty"`
	Quote       string      `json:"quote,omitempty"`
	Error       string      `json:"error,omitempty"`
}

type PackageType int

const (
	// catchall
	Undefined PackageType = iota
	// RequestChallenge is used by the client to request a challenge from the server.
	RequestChallenge
	// SendChallenge is used by the server to send a challenge to the client.
	SendChallenge
	// SendSolution is used by the client to send a solution to the server for validation.
	SendSolution
	// SendQuote is used by the server in case everything is correct to send a quote to the client.
	SendQuote
	// SolutionIncorrect is used by the server to tell the client that the calculated solution is incorrect.
	SolutionIncorrect
	// Error is used to specify a more generic error from server to the client.
	Error
)
