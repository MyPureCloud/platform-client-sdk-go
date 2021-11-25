package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Querywaitlistpositionsrequest
type Querywaitlistpositionsrequest struct { 
	// TimeOffRequests - The list of the time off request ids for which to fetch the daily waitlist positions
	TimeOffRequests *[]Usertimeoffrequestreference `json:"timeOffRequests,omitempty"`

}

func (o *Querywaitlistpositionsrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Querywaitlistpositionsrequest
	
	return json.Marshal(&struct { 
		TimeOffRequests *[]Usertimeoffrequestreference `json:"timeOffRequests,omitempty"`
		*Alias
	}{ 
		TimeOffRequests: o.TimeOffRequests,
		Alias:    (*Alias)(o),
	})
}

func (o *Querywaitlistpositionsrequest) UnmarshalJSON(b []byte) error {
	var QuerywaitlistpositionsrequestMap map[string]interface{}
	err := json.Unmarshal(b, &QuerywaitlistpositionsrequestMap)
	if err != nil {
		return err
	}
	
	if TimeOffRequests, ok := QuerywaitlistpositionsrequestMap["timeOffRequests"].([]interface{}); ok {
		TimeOffRequestsString, _ := json.Marshal(TimeOffRequests)
		json.Unmarshal(TimeOffRequestsString, &o.TimeOffRequests)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Querywaitlistpositionsrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
