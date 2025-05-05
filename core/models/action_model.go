package models

type Action struct {
	ActionID        string `db:"action_id"`
	ActionDetail    string `db:"action_detail"`
	ActionCondition string `db:"action_condition"`
}
