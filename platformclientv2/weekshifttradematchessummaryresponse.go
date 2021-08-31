package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Weekshifttradematchessummaryresponse
type Weekshifttradematchessummaryresponse struct { 
	// WeekDate - The schedule week date in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	WeekDate *time.Time `json:"weekDate,omitempty"`


	// Count - The number of trades in the Matched state for the given week
	Count *int `json:"count,omitempty"`

}

func (o *Weekshifttradematchessummaryresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Weekshifttradematchessummaryresponse
	
	WeekDate := new(string)
	if o.WeekDate != nil {
		*WeekDate = timeutil.Strftime(o.WeekDate, "%Y-%m-%d")
	} else {
		WeekDate = nil
	}
	
	return json.Marshal(&struct { 
		WeekDate *string `json:"weekDate,omitempty"`
		
		Count *int `json:"count,omitempty"`
		*Alias
	}{ 
		WeekDate: WeekDate,
		
		Count: o.Count,
		Alias:    (*Alias)(o),
	})
}

func (o *Weekshifttradematchessummaryresponse) UnmarshalJSON(b []byte) error {
	var WeekshifttradematchessummaryresponseMap map[string]interface{}
	err := json.Unmarshal(b, &WeekshifttradematchessummaryresponseMap)
	if err != nil {
		return err
	}
	
	if weekDateString, ok := WeekshifttradematchessummaryresponseMap["weekDate"].(string); ok {
		WeekDate, _ := time.Parse("2006-01-02", weekDateString)
		o.WeekDate = &WeekDate
	}
	
	if Count, ok := WeekshifttradematchessummaryresponseMap["count"].(float64); ok {
		CountInt := int(Count)
		o.Count = &CountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Weekshifttradematchessummaryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
