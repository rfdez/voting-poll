package postgresql

const (
	sqlOptionTable = "options"
)

type sqlOption struct {
	ID          string `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	PollID      string `db:"poll_id"`
	Votes       int    `db:"votes"`
}
