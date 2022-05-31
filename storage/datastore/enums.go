package datastore

type StateT int

const (
	_ = iota
	Menu
	EmployerTextFill
	EmployerTags
)

type StageT int

const (
	Created = iota
	Payed
	Moderate
	Blocked
	Approved
	DStart
	DFinish
)

type StatusT int

const (
	_ = iota
	Paid
	HasCheck
	Back
)
