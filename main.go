package main

import (
	"fmt"

	"altaproject2/config"
	"altaproject2/factory"
	"altaproject2/utils/database/mysql"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.Getconfig()
	db := mysql.InitDB(cfg)
	mysql.MigrateDB(db)

	e := echo.New()

	factory.InitFactory(e, db)

	fmt.Println("==== STARTING PROGRAM ====")
	address := fmt.Sprintf(":%d", config.SERVERPORT)
	e.Logger.Fatal(e.Start(address))
}
