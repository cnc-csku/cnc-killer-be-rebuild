package models

type Action struct {
	ID        string `db:"action_id"`
	Detail    string `db:"action_detail"`
	Condition string `db:"action_condition"`
}
