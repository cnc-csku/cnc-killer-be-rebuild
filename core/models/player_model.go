package models

type Player struct {
	ID         string  `db:"player_id"`
	SecretCode string  `db:"secret_code"`
	VictimID   *string `db:"victim_id"`
	IsAlive    bool    `db:"is_alive"`
	Score      int     `db:"score"`
	ActionID   *string `db:"action_id"`
	Email      string  `db:"email"`
}
