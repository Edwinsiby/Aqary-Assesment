package models

type StudentData struct {
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	StudentID   string `json:"studentid"`
	MathsMark   int    `json:"mathsmark"`
	EnglishMark int    `json:"englishmark"`
	Age         int    `json:"age"`
}
