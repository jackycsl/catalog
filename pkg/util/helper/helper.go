package helper

import "errors"

var ErrRecordNotFound = errors.New("record not found")

var BackfillGameTopic = "backfillGame"
var BackfillListGameTopic = "backfillListGame"
var CreateGameTopic = "createGame"
