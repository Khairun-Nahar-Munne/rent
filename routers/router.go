package routers

import (
    beego "github.com/beego/beego/v2/server/web"
    "rent/controllers"
)

func init() {
    beego.Router("/api/locations/fetch", &controllers.LocationController{})
    beego.Router("/api/properties/fetch", &controllers.RentalPropertyController{})

}
