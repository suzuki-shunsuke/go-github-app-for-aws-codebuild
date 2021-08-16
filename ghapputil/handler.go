package ghapputil

type Event struct {
	Body    string  `json:"body"`
	Headers Headers `json:"headers"`
}

type Headers struct {
	Event     string `json:"x-github-event"`
	Delivery  string `json:"x-github-delivery"`
	Signature string `json:"x-hub-signature-256"`
}
