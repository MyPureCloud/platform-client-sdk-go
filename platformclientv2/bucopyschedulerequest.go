package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bucopyschedulerequest
type Bucopyschedulerequest struct { 
	// Description - The description for the new schedule
	Description *string `json:"description,omitempty"`


	// WeekDate - The start weekDate for the new copy of the schedule. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	WeekDate *time.Time `json:"weekDate,omitempty"`

}

func (o *Bucopyschedulerequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bucopyschedulerequest
	
	WeekDate := new(string)
	if o.WeekDate != nil {
		*WeekDate = timeutil.Strftime(o.WeekDate, "%Y-%m-%d")
	} else {
		WeekDate = nil
	}
	
	return json.Marshal(&struct { 
		Description *string `json:"description,omitempty"`
		
		WeekDate *string `json:"weekDate,omitempty"`
		*Alias
	}{ 
		Description: o.Description,
		
		WeekDate: WeekDate,
		Alias:    (*Alias)(o),
	})
}

func (o *Bucopyschedulerequest) UnmarshalJSON(b []byte) error {
	var BucopyschedulerequestMap map[string]interface{}
	err := json.Unmarshal(b, &BucopyschedulerequestMap)
	if err != nil {
		return err
	}
	
	if Description, ok := BucopyschedulerequestMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if weekDateString, ok := BucopyschedulerequestMap["weekDate"].(string); ok {
		WeekDate, _ := time.Parse("2006-01-02", weekDateString)
		o.WeekDate = &WeekDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Bucopyschedulerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
