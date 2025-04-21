package requests

const (
	GameStatusWaiting = "waiting"
	GameStatusStart   = "start"
	GameStatusEnd     = "end"
)

var ValidGameStatus = map[string]bool{
	GameStatusWaiting: true,
	GameStatusStart:   true,
	GameStatusEnd:     true,
}

const (
	MsgTypeUpdateStatus = "update-game-status"
	MsgTypeAction       = "actions"
	MsgTypeKill         = "kill"
	MsgTypeRevive       = "revive"
)
