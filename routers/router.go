package routers

import (
    "github.com/beego/beego/v2/server/web"
    "rent/controllers"
)

func init() {
    web.Router("/fetch_locations", &controllers.LocationController{}, "get:FetchAndStoreLocations")
    web.Router("/api/properties/fetch", &controllers.RentalPropertyController{}, "get:FetchAndStoreProperties")
}
