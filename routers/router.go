package routers

import (
    "github.com/beego/beego/v2/server/web"
    "rent/controllers"
)

func init() {
    web.Router("/api/locations/fetch", &controllers.LocationController{}, "get:FetchAndStoreLocations")
    web.Router("/api/properties/fetch", &controllers.RentalPropertyController{}, "get:FetchAndStoreProperties")
    web.Router("/v1/property/list", &controllers.PropertyController{}, "get:ListProperties")
    web.Router("/v1/property/details", &controllers.PropertyController{}, "get:PropertyDetails")
    web.Router("/", &controllers.PropertyController{}, "get:ListPage")
    web.Router("/property-details", &controllers.PropertyController{}, "get:DetailsPage")

}
