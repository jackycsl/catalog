package helper

import (
	"errors"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
)

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

func PasswordFilterFunc(level log.Level, keyvals ...interface{}) bool {
	for i := 0; i < len(keyvals); i++ {
		if keyvals[i] == "args" {
			if strings.Contains(keyvals[i+1].(string), "password") {
				keyvals[i+1] = "***"
			}
		}
	}
	return false
}
