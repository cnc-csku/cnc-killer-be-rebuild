package requests

type AddActionRequest struct {
	Detail    string `json:"action_detail"`
	Condition string `json:"action_condition"`
}
