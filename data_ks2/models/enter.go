package models

type Model struct {
	ID int `json:"id" gromL:"primaryKey;not null"`
}
