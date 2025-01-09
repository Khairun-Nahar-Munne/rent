package routers

import (
    "github.com/beego/beego/v2/server/web"
    "rent/controllers"
)

func init() {
    web.Router("/api/locations/fetch", &controllers.LocationController{}, "get:FetchAndStoreLocations")
    web.Router("/api/properties/fetch", &controllers.RentalPropertyController{}, "get:FetchAndStoreProperties")
}
