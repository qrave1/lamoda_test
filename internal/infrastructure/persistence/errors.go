package persistence

type DatabaseError struct {
	msg string
}

func (d DatabaseError) Error() string {
	return d.msg
}

var (
	TransactionStartError  = DatabaseError{msg: "Transaction start error"}
	TransactionCommitError = DatabaseError{msg: "Transaction commit error"}
	DuplicateError         = DatabaseError{msg: "Duplicate rows"}
	NoRowsFoundError       = DatabaseError{msg: "No rows found"}
	NoRowsAffected         = DatabaseError{msg: "No rows affected"}
)
