package app

import (
	"crucial/banner/app/config"
	"crucial/banner/app/db"
	"crucial/banner/app/server"
)

// Run - начальная загрузка
func Run() {
	var dbCfg = config.Peek().Database

	db.Connect(dbCfg.User, dbCfg.Password, dbCfg.DBName, dbCfg.Host, dbCfg.Port)
	defer db.Disconnect()

	server.Listen()
}
