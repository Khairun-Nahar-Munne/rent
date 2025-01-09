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
        SearchMeta struct {
            NbAdults   int `json:"nbAdults"`
            NbChildren int `json:"nbChildren"`
        } `json:"searchMeta"`
    } `json:"data"`
}

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
    
    // Get tomorrow's date for check-in and the day after tomorrow for check-out
    checkIn := time.Now().AddDate(0, 0, 1).Format("2006-01-02")
    checkOut := time.Now().AddDate(0, 0, 2).Format("2006-01-02")

    var properties []models.RentalProperty  // Slice to store fetched properties

    for _, location := range locations {
        

        url := fmt.Sprintf("https://booking-com18.p.rapidapi.com/web/stays/search?destId=%s&destType=%s&checkIn=%s&checkOut=%s",
            location.DestId, location.DestType, checkIn, checkOut)

        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
            fmt.Printf("Error creating request for location %s: %v\n", location.Value, err)
            continue
        }

        req.Header.Add("x-rapidapi-host", "booking-com18.p.rapidapi.com")
        req.Header.Add("x-rapidapi-key", "a086eb4944mshde04dbeb7635d1dp153b0fjsn73c31892df55")

        resp, err := client.Do(req)
        if err != nil {
            fmt.Printf("Error making request for location %s: %v\n", location.Value, err)
            continue
        }
        defer resp.Body.Close()

        var searchResp SearchResponse
        decoder := json.NewDecoder(resp.Body)
        if err := decoder.Decode(&searchResp); err != nil {
            fmt.Printf("Error decoding response for location %s: %v\n", location.Value, err)
            continue
        }

        // Process and store each property
        for _, result := range searchResp.Data.Results {
            // Fetch additional details from the second API
            detailURL := fmt.Sprintf("https://booking-com18.p.rapidapi.com/stays/detail?hotelId=%d&checkinDate=%s&checkoutDate=%s&units=metric",
                result.BasicPropertyData.ID, checkIn, checkOut)

            detailReq, err := http.NewRequest("GET", detailURL, nil)
            if err != nil {
                fmt.Printf("Error creating detail request for property %d: %v\n", result.BasicPropertyData.ID, err)
                continue
            }

            detailReq.Header.Add("x-rapidapi-host", "booking-com18.p.rapidapi.com")
            detailReq.Header.Add("x-rapidapi-key", "a086eb4944mshde04dbeb7635d1dp153b0fjsn73c31892df55")

            detailResp, err := client.Do(detailReq)
            if err != nil {
                fmt.Printf("Error making detail request for property %d: %v\n", result.BasicPropertyData.ID, err)
                continue
            }
            defer detailResp.Body.Close()

            // Read the raw response body
            rawDetailRespBody, err := ioutil.ReadAll(detailResp.Body)
            if err != nil {
                fmt.Printf("Error reading detail response body for property %d: %v\n", result.BasicPropertyData.ID, err)
                continue
            }

            // Print the raw response body for debugging
            fmt.Printf("Raw detail response body for property %d: %s\n", result.BasicPropertyData.ID, rawDetailRespBody)

            var detailRespData PropertyDetailResponse
            if err := json.Unmarshal(rawDetailRespBody, &detailRespData); err != nil {
                fmt.Printf("Error decoding detail response for property %d: %v\n", result.BasicPropertyData.ID, err)
                continue
            }

            // Debug print the detail response
            fmt.Printf("Detail response for property %d: %+v\n", result.BasicPropertyData.ID, detailRespData)

            // Extract amenities (first three names from facilities)
            amenities := []string{}
            for i, facility := range detailRespData.Data.FacilitiesBlock.Facilities {
                if i >= 3 {
                    break
                }
                amenities = append(amenities, facility.Name)
            }

            // Extract type
            propertyType := detailRespData.Data.AccommodationTypeName

            // Extract guest count
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
                Type:            propertyType,
            }

            // Store breadcrumbs
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
            fmt.Printf("Inserted property with ID: %d\n", id)

            // Append the property to the slice
            properties = append(properties, property)
        }
    }

    c.Data["json"] = map[string]interface{}{
        "success":  true,
        "message":  "Properties fetched and stored successfully",
        "properties": properties,  // Include the fetched properties in the response
    }
    c.ServeJSON()
}