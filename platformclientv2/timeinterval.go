package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Timeinterval
type Timeinterval struct { 
	// Months
	Months *int `json:"months,omitempty"`


	// Weeks
	Weeks *int `json:"weeks,omitempty"`


	// Days
	Days *int `json:"days,omitempty"`


	// Hours
	Hours *int `json:"hours,omitempty"`

}

func (o *Timeinterval) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Timeinterval
	
	return json.Marshal(&struct { 
		Months *int `json:"months,omitempty"`
		
		Weeks *int `json:"weeks,omitempty"`
		
		Days *int `json:"days,omitempty"`
		
		Hours *int `json:"hours,omitempty"`
		*Alias
	}{ 
		Months: o.Months,
		
		Weeks: o.Weeks,
		
		Days: o.Days,
		
		Hours: o.Hours,
		Alias:    (*Alias)(o),
	})
}

func (o *Timeinterval) UnmarshalJSON(b []byte) error {
	var TimeintervalMap map[string]interface{}
	err := json.Unmarshal(b, &TimeintervalMap)
	if err != nil {
		return err
	}
	
	if Months, ok := TimeintervalMap["months"].(float64); ok {
		MonthsInt := int(Months)
		o.Months = &MonthsInt
	}
	
	if Weeks, ok := TimeintervalMap["weeks"].(float64); ok {
		WeeksInt := int(Weeks)
		o.Weeks = &WeeksInt
	}
	
	if Days, ok := TimeintervalMap["days"].(float64); ok {
		DaysInt := int(Days)
		o.Days = &DaysInt
	}
	
	if Hours, ok := TimeintervalMap["hours"].(float64); ok {
		HoursInt := int(Hours)
		o.Hours = &HoursInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Timeinterval) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
