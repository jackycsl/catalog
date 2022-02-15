package helper

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrUserNotFound   = errors.New("user not found")
)

var (
	BackfillGameTopic     = "backfillGame"
	BackfillListGameTopic = "backfillListGame"
	CreateGameTopic       = "createGame"
	UpdateGameTopic       = "updateGame"
)
