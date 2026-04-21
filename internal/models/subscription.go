package models

import "time"

type Subsciption struct {
	id               int
	chat_id          int
	interval_minutes int
	is_active        bool
	last_send_at     time.Time
}
