package main

import (
	Log "go-rest-api/logwrapper"
	"go-rest-api/restfullapi"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// _, b, _, _ := runtime.Caller(0)
	// basepath := filepath.Dir(b)
	// fmt.Println(basepath)
	Log.STDLog.Fatal(restfullapi.RunAPI("."))

}
