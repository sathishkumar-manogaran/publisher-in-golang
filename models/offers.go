package models

import (
	"database/sql/driver"
	"gorm.io/datatypes"
)

type Offers struct {
	Offer []Offer `json:"offers"`
}

type Offer struct {
	CmOfferId    string `json:"cm_offer_id"`
	Hotel        `json:"hotel"`
	Room         `json:"room"`
	RatePlan     `json:"rate_plan"`
	OriginalData `json:"original_data"`
	Capacity     `json:"capacity"`
	Number       `json:"number"`
	Price        `json:"price"`
	Currency     `json:"currency"`
	CheckIn      `json:"check_in"`
	CheckOut     `json:"check_out"`
	Fees         []Fee `json:"fees"`
}

type Hotel struct {
	HotelId     string         `json:"hotel_id" gorm:"primaryKey;size:10"`
	Name        string         `json:"name"`
	Country     string         `json:"country"`
	Address     string         `json:"address"`
	Latitude    float64        `json:"latitude"`
	Longitude   float64        `json:"longitude"`
	Telephone   string         `json:"telephone"`
	Amenities   datatypes.JSON `json:"amenities"`
	Description string         `json:"description"`
	RoomCount   int8           `json:"room_count"`
	Currency    string         `json:"currency"`
	Rooms       []Room         `gorm:"foreignKey:hotel_id"`
	RatePlans   []RatePlan     `gorm:"foreignKey:hotel_id"`
}

type Room struct {
	RoomId      string         `json:"room_id" gorm:"primaryKey"`
	HotelId     string         `json:"hotel_id" gorm:"size:10"`
	Description string         `json:"description"`
	Name        string         `json:"name"`
	Capacity    datatypes.JSON `json:"capacity"`
}

type RatePlan struct {
	RatePlanId         string         `json:"rate_plan_id" gorm:"primaryKey"`
	HotelId            string         `json:"hotel_id" gorm:"size:10"`
	CancellationPolicy datatypes.JSON `json:"cancellation_policy"`
	Name               string         `json:"name"`
	OtherConditions    datatypes.JSON `json:"other_conditions"`
	MealPlan           string         `json:"meal_plan"`
}

type Capacity struct {
	MaxAdults     int8 `json:"max_adults"`
	ExtraChildren int8 `json:"extra_children"`
}

type CancellationPolicy struct {
	Type              string `json:"type"`
	ExpiresDaysBefore int8   `json:"expires_days_before"`
}

type OriginalData struct {
	GuaranteePolicy `json:"GuaranteePolicy"`
}

type GuaranteePolicy struct {
	Required bool `json:"Required"`
}

type Number int
type Price int
type Currency string
type CheckIn string
type CheckOut string

type Fee struct {
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Included    bool    `json:"included"`
	Percent     float32 `json:"percent"`
}




// Scan scan value into Jsonb, implements sql.Scanner interface
//func (j *Capacity) Scan(value interface{}) error {
//	bytes, ok := value.([]byte)
//	if !ok {
//		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
//	}
//
//	result := json.RawMessage{}
//	err := json.Unmarshal(bytes, &result)
//	*j = Capacity{
//		MaxAdults:     result.MarshalJSON(),
//		ExtraChildren: 0,
//	}
//	return err
//}
//
//// Value return json value, implement driver.Valuer interface
//func (j Capacity) Value() (driver.Value, error) {
//	if len(j) == 0 {
//		return nil, nil
//	}
//	return json.RawMessage(j).MarshalJSON()
//}

func (capacityValue *Capacity) Value() (driver.Value, error) {
	//fmt.Print(driver.Value)
	//adults := capacityValue.MaxAdults
	//children := capacityValue.ExtraChildren
	//
	//buf, ok := adults.([]byte)
	//
	//mysqlEncoding := make([]byte, 4)
	return capacityValue.ConvertJSONToString(), nil
}

func (capacityValue *Capacity) Scan(value interface{}) error {
	//*data = data
	if value == nil {
		capacityValue.MaxAdults, capacityValue.ExtraChildren = 5, 6
		return nil
	}
	capacity := value.(Capacity)
	capacityValue = &capacity
	return nil
}

func (capacityValue *Capacity) ConvertJSONToString() driver.Value {
	//value := driver.Value()
	//value.
	return capacityValue
}

//func (data *Capacity) ConvertStringToJson(valueString interface{}) Capacity {
//	//return data
//}
