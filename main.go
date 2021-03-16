package main

import (
	Log "go-rest-api/logwrapper"
	"go-rest-api/restfullapi"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	Log.STDLog.Fatal(restfullapi.RunAPI("config.json"))
	Log.STDLog.Info("database run ...")

}
