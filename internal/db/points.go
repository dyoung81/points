package db

import "time"

type pointsTransactionRow struct {
	ID        string
	player    string
	points    points
	timestamp time.Time
}

type points struct {
	points int
}
