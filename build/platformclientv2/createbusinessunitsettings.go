package platformclientv2
import (
	"encoding/json"
)

// Createbusinessunitsettings
type Createbusinessunitsettings struct { 
	// StartDayOfWeek - The start day of week for this business unit
	StartDayOfWeek *string `json:"startDayOfWeek,omitempty"`


	// TimeZone - The time zone for this business unit, using the Olsen tz database format
	TimeZone *string `json:"timeZone,omitempty"`


	// ShortTermForecasting - Short term forecasting settings
	ShortTermForecasting *Bushorttermforecastingsettings `json:"shortTermForecasting,omitempty"`

}

// String returns a JSON representation of the model
func (o *Createbusinessunitsettings) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
