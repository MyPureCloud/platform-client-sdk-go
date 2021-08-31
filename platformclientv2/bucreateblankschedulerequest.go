package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bucreateblankschedulerequest
type Bucreateblankschedulerequest struct { 
	// Description - The description for the schedule
	Description *string `json:"description,omitempty"`


	// ShortTermForecast - The forecast to use when generating the schedule.  Note that the forecast must fully encompass the schedule's start week + week count
	ShortTermForecast *Bushorttermforecastreference `json:"shortTermForecast,omitempty"`


	// WeekCount - The number of weeks in the schedule. One extra day is added at the end
	WeekCount *int `json:"weekCount,omitempty"`

}

func (o *Bucreateblankschedulerequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bucreateblankschedulerequest
	
	return json.Marshal(&struct { 
		Description *string `json:"description,omitempty"`
		
		ShortTermForecast *Bushorttermforecastreference `json:"shortTermForecast,omitempty"`
		
		WeekCount *int `json:"weekCount,omitempty"`
		*Alias
	}{ 
		Description: o.Description,
		
		ShortTermForecast: o.ShortTermForecast,
		
		WeekCount: o.WeekCount,
		Alias:    (*Alias)(o),
	})
}

func (o *Bucreateblankschedulerequest) UnmarshalJSON(b []byte) error {
	var BucreateblankschedulerequestMap map[string]interface{}
	err := json.Unmarshal(b, &BucreateblankschedulerequestMap)
	if err != nil {
		return err
	}
	
	if Description, ok := BucreateblankschedulerequestMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if ShortTermForecast, ok := BucreateblankschedulerequestMap["shortTermForecast"].(map[string]interface{}); ok {
		ShortTermForecastString, _ := json.Marshal(ShortTermForecast)
		json.Unmarshal(ShortTermForecastString, &o.ShortTermForecast)
	}
	
	if WeekCount, ok := BucreateblankschedulerequestMap["weekCount"].(float64); ok {
		WeekCountInt := int(WeekCount)
		o.WeekCount = &WeekCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Bucreateblankschedulerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
