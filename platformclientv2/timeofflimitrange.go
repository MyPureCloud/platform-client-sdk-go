package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Timeofflimitrange - A single range filled with time off limit interval values
type Timeofflimitrange struct { 
	// StartDate - Start date of the range. The end date is determined by 'granularity' and the size of 'limitMinutesPerInterval'. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	StartDate *time.Time `json:"startDate,omitempty"`


	// Granularity - Granularity choice for the time off limit
	Granularity *string `json:"granularity,omitempty"`


	// LimitMinutesPerInterval - The list of time off limit values in minutes per granularity interval. If 'null' is specified, then interval specific value is cleared. Such interval will have 'defaultLimitMinutes' value
	LimitMinutesPerInterval *[]int `json:"limitMinutesPerInterval,omitempty"`

}

func (o *Timeofflimitrange) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Timeofflimitrange
	
	StartDate := new(string)
	if o.StartDate != nil {
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%d")
	} else {
		StartDate = nil
	}
	
	return json.Marshal(&struct { 
		StartDate *string `json:"startDate,omitempty"`
		
		Granularity *string `json:"granularity,omitempty"`
		
		LimitMinutesPerInterval *[]int `json:"limitMinutesPerInterval,omitempty"`
		*Alias
	}{ 
		StartDate: StartDate,
		
		Granularity: o.Granularity,
		
		LimitMinutesPerInterval: o.LimitMinutesPerInterval,
		Alias:    (*Alias)(o),
	})
}

func (o *Timeofflimitrange) UnmarshalJSON(b []byte) error {
	var TimeofflimitrangeMap map[string]interface{}
	err := json.Unmarshal(b, &TimeofflimitrangeMap)
	if err != nil {
		return err
	}
	
	if startDateString, ok := TimeofflimitrangeMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02", startDateString)
		o.StartDate = &StartDate
	}
	
	if Granularity, ok := TimeofflimitrangeMap["granularity"].(string); ok {
		o.Granularity = &Granularity
	}
	
	if LimitMinutesPerInterval, ok := TimeofflimitrangeMap["limitMinutesPerInterval"].([]interface{}); ok {
		LimitMinutesPerIntervalString, _ := json.Marshal(LimitMinutesPerInterval)
		json.Unmarshal(LimitMinutesPerIntervalString, &o.LimitMinutesPerInterval)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Timeofflimitrange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
