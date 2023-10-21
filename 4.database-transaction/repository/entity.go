package repository

type School struct {
	StudentID  string
	SchoolID   string
	SchoolName string
	StateID    string
}

type SchoolMarkSheet struct {
	StudentID   string
	StudentName string
	AdmissionNo int
	SchoolID    string
	TotalMark   int
}

type StateMarkSheet struct {
	StudentID string
	SchoolID  string
	StateID   string
	TotalMark int
}

type SchoolRank struct {
	SchoolID  string
	StudentID string
	TotalMark int
}

type StateRank struct {
	SchoolID  string
	StudentID string
	StateID   string
	TotalMark int
}

type Prohbation struct {
	StudentID string
	SchoolID  string
}

type StudentInfo struct {
	StudentID  string
	SchoolID   string
	SchoolRank int
	StateRank  int
	Prohbation bool
}
