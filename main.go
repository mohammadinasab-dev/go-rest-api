package main

import (
	Log "go-rest-api/logwrapper"
	"go-rest-api/restfullapi"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	Log.ErrorLog.Fatal(restfullapi.RunAPI("config.json"))
	Log.InfoLog.Info("database run ...")

}
