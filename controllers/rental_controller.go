package controllers

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
    "time"
    "github.com/beego/beego/v2/client/orm"
    "github.com/beego/beego/v2/server/web"
    "rent/models"
)

type RentalPropertyController struct {
    web.Controller
}

// Add new struct for the first API response
type FirstAPIResponse struct {
    Data struct {
        HotelPhotos []struct {
            LargeUrl string `json:"large_url"`
        } `json:"hotelPhotos"`
        HotelTranslation []struct {
            Description string `json:"description"`
        } `json:"hotelTranslation"`
    } `json:"data"`
    Status  bool   `json:"status"`
    Message string `json:"message"`
}

// Update PropertyDetailResponse to include city_in_trans
type PropertyDetailResponse struct {
    Data struct {
        FacilitiesBlock struct {
            Facilities []struct {
                Name string `json:"name"`
                Icon string `json:"icon"`
            } `json:"facilities"`
        } `json:"facilities_block"`
        TpiBlock []struct {
            GuestCount int `json:"guest_count"`
        } `json:"tpi_block"`
        AccommodationTypeName string `json:"accommodation_type_name"`
        CityInTrans          string `json:"city_in_trans"`
    } `json:"data"`
}

type SearchResponse struct {
    Data struct {
        Results []struct {
            BasicPropertyData struct {
                ID      int64 `json:"id"`
                Reviews struct {
                    TotalScore   float64 `json:"totalScore"`
                    ReviewsCount int     `json:"reviewsCount"`
                }
            } `json:"basicPropertyData"`
            DisplayName struct {
                Text string `json:"text"`
            } `json:"displayName"`
            MatchingUnitConfigurations struct {
                CommonConfiguration struct {
                    NbBathrooms int `json:"nbBathrooms"`
                    NbBedrooms int `json:"nbBedrooms"`
                }
            } `json:"matchingUnitConfigurations"`
            PriceDisplayInfoIrene struct {
                DisplayPrice struct {
                    AmountPerStay struct {
                        Amount string `json:"amount"`
                    } `json:"amountPerStay"`
                }
            } `json:"priceDisplayInfoIrene"`
            Location struct {
                DisplayLocation string `json:"displayLocation"`
            } `json:"location"`
            ID string `json:"id"`
        } `json:"results"`
        Breadcrumbs []struct {
            Name string `json:"name"`
        } `json:"breadcrumbs"`
    } `json:"data"`
}


func (c *RentalPropertyController) FetchAndStoreProperties() {
    o := orm.NewOrm()
    
    // Get all locations from the database
    var locations []models.Location
    _, err := o.QueryTable("location").All(&locations)
    if err != nil {
        c.Data["json"] = map[string]interface{}{"error": fmt.Sprintf("Error fetching locations: %v", err)}
        c.ServeJSON()
        return
    }

    client := &http.Client{}
    
    checkIn := time.Now().AddDate(0, 0, 1).Format("2006-01-02")
    checkOut := time.Now().AddDate(0, 0, 3).Format("2006-01-02")

    var properties []models.RentalProperty

    for _, location := range locations {
        url := fmt.Sprintf("https://booking-com18.p.rapidapi.com/web/stays/search?destId=%s&destType=%s&checkIn=%s&checkOut=%s",
            location.DestId, location.DestType, checkIn, checkOut)

        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
            fmt.Printf("Error creating request for location %s: %v\n", location.Value, err)
            continue
        }

        req.Header.Add("x-rapidapi-host", "booking-com18.p.rapidapi.com")
        req.Header.Add("x-rapidapi-key", "87b0822e28msh2a36d071068c413p1b1a5ejsn6091577e3bce")

        resp, err := client.Do(req)
        if err != nil {
            fmt.Printf("Error making request for location %s: %v\n", location.Value, err)
            continue
        }
        defer resp.Body.Close()

        var searchResp SearchResponse
        if err := json.NewDecoder(resp.Body).Decode(&searchResp); err != nil {
            fmt.Printf("Error decoding response for location %s: %v\n", location.Value, err)
            continue
        }

        for _, result := range searchResp.Data.Results {
            // First API call to get photos and description
            firstDetailURL := fmt.Sprintf("https://booking-com18.p.rapidapi.com/web/stays/details?id=%s&checkIn=%s&checkOut=%s", result.ID, checkIn, checkOut)

            firstDetailReq, err := http.NewRequest("GET", firstDetailURL, nil)
            if err != nil {
                fmt.Printf("Error creating first detail request for property %d: %v\n", result.BasicPropertyData.ID, err)
                continue
            }

            firstDetailReq.Header.Add("x-rapidapi-host", "booking-com18.p.rapidapi.com")
            firstDetailReq.Header.Add("x-rapidapi-key", "87b0822e28msh2a36d071068c413p1b1a5ejsn6091577e3bce")

            firstDetailResp, err := client.Do(firstDetailReq)
            if err != nil {
                fmt.Printf("Error making first detail request for property %d: %v\n", result.BasicPropertyData.ID, err)
                continue
            }
            defer firstDetailResp.Body.Close()

            var firstDetailData FirstAPIResponse
            firstDetailBody, _ := ioutil.ReadAll(firstDetailResp.Body)
            if err := json.Unmarshal(firstDetailBody, &firstDetailData); err != nil {
                fmt.Printf("Error decoding first detail response for property %d: %v\n", result.BasicPropertyData.ID, err)
                continue
            }

            // Second API call (existing detail API)
            detailURL := fmt.Sprintf("https://booking-com18.p.rapidapi.com/stays/detail?hotelId=%d&checkinDate=%s&checkoutDate=%s&units=metric",
                result.BasicPropertyData.ID, checkIn, checkOut)

            detailReq, err := http.NewRequest("GET", detailURL, nil)
            if err != nil {
                fmt.Printf("Error creating detail request for property %d: %v\n", result.BasicPropertyData.ID, err)
                continue
            }

            detailReq.Header.Add("x-rapidapi-host", "booking-com18.p.rapidapi.com")
            detailReq.Header.Add("x-rapidapi-key", "87b0822e28msh2a36d071068c413p1b1a5ejsn6091577e3bce")

            detailResp, err := client.Do(detailReq)
            if err != nil {
                fmt.Printf("Error making detail request for property %d: %v\n", result.BasicPropertyData.ID, err)
                continue
            }
            defer detailResp.Body.Close()

            var detailRespData PropertyDetailResponse
            detailBody, _ := ioutil.ReadAll(detailResp.Body)
            if err := json.Unmarshal(detailBody, &detailRespData); err != nil {
                fmt.Printf("Error decoding detail response for property %d: %v\n", result.BasicPropertyData.ID, err)
                continue
            }

            // Create and store the RentalProperty
            amenities := []string{}
            for i, facility := range detailRespData.Data.FacilitiesBlock.Facilities {
                if i >= 3 {
                    break
                }
                amenities = append(amenities, facility.Name)
            }

            guestCount := 0
            if len(detailRespData.Data.TpiBlock) > 0 {
                guestCount = detailRespData.Data.TpiBlock[0].GuestCount
            }

            property := models.RentalProperty{
                Location:        &location,
                PropertyId:      result.BasicPropertyData.ID,
                PropertySlugId:  result.ID,
                HotelName:       result.DisplayName.Text,
                Bedrooms:        result.MatchingUnitConfigurations.CommonConfiguration.NbBedrooms,
                Bathrooms:       result.MatchingUnitConfigurations.CommonConfiguration.NbBathrooms,
                GuestCount:      guestCount,
                Rating:          result.BasicPropertyData.Reviews.TotalScore,
                ReviewCount:     result.BasicPropertyData.Reviews.ReviewsCount,
                Price:           result.PriceDisplayInfoIrene.DisplayPrice.AmountPerStay.Amount,
                DisplayLocation: result.Location.DisplayLocation,
                Amenities:       strings.Join(amenities, ", "),
                Type:           detailRespData.Data.AccommodationTypeName,
            }

            for i, breadcrumb := range searchResp.Data.Breadcrumbs {
                switch i {
                case 0:
                    property.Breadcrumb1 = breadcrumb.Name
                case 1:
                    property.Breadcrumb2 = breadcrumb.Name
                case 2:
                    property.Breadcrumb3 = breadcrumb.Name
                case 3:
                    property.Breadcrumb4 = breadcrumb.Name
                }
            }

            // Insert the property
            id, err := o.Insert(&property)
            if err != nil {
                fmt.Printf("Error inserting property %s: %v\n", property.HotelName, err)
                continue
            }

            // Create and store PropertyDetails
            propertyDetails := &models.PropertyDetails{
                RentalProperty: &property,
                CityInTrans:    detailRespData.Data.CityInTrans,
            }

            // Set description if available
            if len(firstDetailData.Data.HotelTranslation) > 0 {
                propertyDetails.Description = firstDetailData.Data.HotelTranslation[0].Description
            }

            // Set up to 5 image URLs
            for i, photo := range firstDetailData.Data.HotelPhotos {
                if i >= 5 {
                    break
                }
                switch i {
                case 0:
                    propertyDetails.ImageUrl1 = photo.LargeUrl
                case 1:
                    propertyDetails.ImageUrl2 = photo.LargeUrl
                case 2:
                    propertyDetails.ImageUrl3 = photo.LargeUrl
                case 3:
                    propertyDetails.ImageUrl4 = photo.LargeUrl
                case 4:
                    propertyDetails.ImageUrl5 = photo.LargeUrl
                }
            }

            // Insert property details
            _, err = o.Insert(propertyDetails)
            if err != nil {
                fmt.Printf("Error inserting property details for property %d: %v\n", id, err)
                continue
            }

            properties = append(properties, property)
            fmt.Printf("Inserted property and details with ID: %d\n", id)
        }
    }

    c.Data["json"] = map[string]interface{}{
        "success":    true,
        "message":    "Properties and details fetched and stored successfully",
        "properties": properties,
    }
    c.ServeJSON()
}