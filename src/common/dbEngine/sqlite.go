package dbEngine

func StartSqlite(dbName string) {
	startDBEngine(DRIVER_SQLITE3, dbName)
}
