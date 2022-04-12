package ports

type DbPort interface {
	CloseDbConnection()
	AddHistory(answer int32, operation string) error
}
