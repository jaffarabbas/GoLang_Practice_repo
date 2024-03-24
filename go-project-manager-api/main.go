package main

import "github.com/go-sql-driver/mysql"

func main() {
	cfg := mysql.Config{
		User:                 Envs.DBUser,
		Passwd:               Envs.DBPassword,
		Addr:                 Envs.DBAddress,
		DBName:               Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	sqlStorage := NewMySqlStorage(cfg)
	db, err := sqlStorage.init()
	if err != nil {
		panic(err)
	}
	store := NewStore(db)

	api := NewApiServer(":8080", store)
	api.Serve()
}
