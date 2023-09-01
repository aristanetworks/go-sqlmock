package sqlmock

type statement struct {
	conn  *sqlmock
	ex    *ExpectedPrepare
	query string
}

func (stmt *statement) Close() error {
	stmt.ex.Lock()
	stmt.ex.wasClosed = true
	stmt.ex.Unlock()
	return stmt.ex.closeErr
}

func (stmt *statement) NumInput() int {
	return -1
}
