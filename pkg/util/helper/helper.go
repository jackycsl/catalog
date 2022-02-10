package helper

import "errors"

var ErrRecordNotFound = errors.New("record not found")

var BackfillGameTopic = "backfillGame"
var CreateGameTopic = "createGame"
