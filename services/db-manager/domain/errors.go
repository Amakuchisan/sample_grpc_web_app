package domain

import "errors"

// ErrNotFound はリポジトリにエンティティが見つからなかったときに返される
var ErrNotFound = errors.New("not found")
