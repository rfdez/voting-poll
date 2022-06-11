package postgresql

const (
	sqlPollTable = "polls"
)

type sqlPoll struct {
	ID          string `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Voters      int    `db:"voters"`
}
