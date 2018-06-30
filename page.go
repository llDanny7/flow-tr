package main

type Page struct {
	Text      string     `json:"text"`
	InState   int        `json:"intStates"`
	Character string     `json:"character"`
	Decisions []Decision `json:"decisions"`
	IsFinal   bool       `json:"isFinal"`
}

type Decision struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type Pages []Page
