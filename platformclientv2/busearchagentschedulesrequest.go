package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Busearchagentschedulesrequest
type Busearchagentschedulesrequest struct { 
	// StartDate - Start date of the range to search. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - End date of the range to search. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndDate *time.Time `json:"endDate,omitempty"`


	// UserIds - IDs of the users for whose schedules to search
	UserIds *[]string `json:"userIds,omitempty"`

}

func (o *Busearchagentschedulesrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Busearchagentschedulesrequest
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	EndDate := new(string)
	if o.EndDate != nil {
		
		*EndDate = timeutil.Strftime(o.EndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndDate = nil
	}
	
	return json.Marshal(&struct { 
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		
		UserIds *[]string `json:"userIds,omitempty"`
		*Alias
	}{ 
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		UserIds: o.UserIds,
		Alias:    (*Alias)(o),
	})
}

func (o *Busearchagentschedulesrequest) UnmarshalJSON(b []byte) error {
	var BusearchagentschedulesrequestMap map[string]interface{}
	err := json.Unmarshal(b, &BusearchagentschedulesrequestMap)
	if err != nil {
		return err
	}
	
	if startDateString, ok := BusearchagentschedulesrequestMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := BusearchagentschedulesrequestMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	
	if UserIds, ok := BusearchagentschedulesrequestMap["userIds"].([]interface{}); ok {
		UserIdsString, _ := json.Marshal(UserIds)
		json.Unmarshal(UserIdsString, &o.UserIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Busearchagentschedulesrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
