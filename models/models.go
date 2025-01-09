package models

import (
    "github.com/beego/beego/v2/client/orm"
)

type Location struct {
    Id       int64  `orm:"auto;pk"`
    DestId   string `orm:"size(100);column(dest_id)"`
    DestType string `orm:"size(100);column(dest_type)"`
    Value    string `orm:"size(255);column(value)"`
}

type RentalProperty struct {
    Id              int64    `orm:"auto;pk"`
    Location        *Location `orm:"rel(fk);column(location_id)"` // Foreign key to Location table
    PropertyId      int64    `orm:"column(property_id)"`
    PropertySlugId  string   `orm:"size(255);column(property_slug_id)"`
    HotelName       string   `orm:"size(255);column(hotel_name)"`
    Bedrooms        int      `orm:"default(0);column(bedrooms)"`
    Bathrooms       int      `orm:"default(0);column(bathrooms)"`
    GuestCount      int      `orm:"default(0);column(guest_count)"`
    Rating          float64  `orm:"digits(3);decimals(1);column(rating)"`
    ReviewCount     int      `orm:"default(0);column(review_count)"`
    Price           string   `orm:"size(50);column(price)"`
    Breadcrumb1     string   `orm:"size(255);null;column(breadcrumb_1)"`
    Breadcrumb2     string   `orm:"size(255);null;column(breadcrumb_2)"`
    Breadcrumb3     string   `orm:"size(255);null;column(breadcrumb_3)"`
    Breadcrumb4     string   `orm:"size(255);null;column(breadcrumb_4)"`
    DisplayLocation string   `orm:"size(255);column(display_location)"`
    Amenities       string   `orm:"type(text);column(amenities)"` // New field for amenities
    Type            string   `orm:"size(255);column(type)"`       // New field for type
}

type PropertyDetails struct {
    Id             int64          `orm:"auto;pk"`
    RentalProperty *RentalProperty `orm:"rel(fk);column(rental_property_id)"` // Foreign key to RentalProperty table
    ImageUrl       string         `orm:"size(500);column(image_url)"`
    Description    string         `orm:"type(text);column(description)"`
}

func init() {
    // Register models with orm
    orm.RegisterModel(new(RentalProperty))
    orm.RegisterModel(new(PropertyDetails))
    orm.RegisterModel(new(Location))
}