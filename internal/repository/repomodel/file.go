package repomodel

type File struct {
	ID   int     `db:"id"`
	Name *string `db:"name"`
	Path *string `db:"path"`
}
