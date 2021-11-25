package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Localdaterange
type Localdaterange struct { 
	// StartDate - The inclusive start of a date range in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - The inclusive end of a date range in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	EndDate *time.Time `json:"endDate,omitempty"`

}

func (o *Localdaterange) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Localdaterange
	
	StartDate := new(string)
	if o.StartDate != nil {
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%d")
	} else {
		StartDate = nil
	}
	
	EndDate := new(string)
	if o.EndDate != nil {
		*EndDate = timeutil.Strftime(o.EndDate, "%Y-%m-%d")
	} else {
		EndDate = nil
	}
	
	return json.Marshal(&struct { 
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		*Alias
	}{ 
		StartDate: StartDate,
		
		EndDate: EndDate,
		Alias:    (*Alias)(o),
	})
}

func (o *Localdaterange) UnmarshalJSON(b []byte) error {
	var LocaldaterangeMap map[string]interface{}
	err := json.Unmarshal(b, &LocaldaterangeMap)
	if err != nil {
		return err
	}
	
	if startDateString, ok := LocaldaterangeMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := LocaldaterangeMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02", endDateString)
		o.EndDate = &EndDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Localdaterange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
