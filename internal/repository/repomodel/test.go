package repomodel

type TestRequest struct {
	Message string `db:"message"`
}

type TestResponse struct {
	Message string `db:"message"`
}
