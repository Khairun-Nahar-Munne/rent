package main

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
	beego "github.com/beego/beego/v2/server/web"
	_ "rent/routers"
)

func main() {
	// Register database
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=munne password=munne123 host=localhost port=5432 dbname=rent sslmode=disable")
	orm.RunSyncdb("default", false, true)



	// Start the Beego application
	beego.Run()
}
