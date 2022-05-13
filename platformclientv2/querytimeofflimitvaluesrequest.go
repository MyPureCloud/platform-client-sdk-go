package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Querytimeofflimitvaluesrequest - This sets up a filter to request date ranges of time off limit, allocated and waitlisted minutes
type Querytimeofflimitvaluesrequest struct { 
	// TimeOffLimitId - The time off limit object id to retrieve values for. Required if activityCodeId is not specified
	TimeOffLimitId *string `json:"timeOffLimitId,omitempty"`


	// ActivityCodeId - The activity code id to filter the affected limit objects by. Required if timeOffLimitId is not specified
	ActivityCodeId *string `json:"activityCodeId,omitempty"`


	// DateRanges - The list of the date ranges to return time off limit, allocated and waitlisted minutes.
	DateRanges *[]Localdaterange `json:"dateRanges,omitempty"`

}

func (o *Querytimeofflimitvaluesrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Querytimeofflimitvaluesrequest
	
	return json.Marshal(&struct { 
		TimeOffLimitId *string `json:"timeOffLimitId,omitempty"`
		
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		DateRanges *[]Localdaterange `json:"dateRanges,omitempty"`
		*Alias
	}{ 
		TimeOffLimitId: o.TimeOffLimitId,
		
		ActivityCodeId: o.ActivityCodeId,
		
		DateRanges: o.DateRanges,
		Alias:    (*Alias)(o),
	})
}

func (o *Querytimeofflimitvaluesrequest) UnmarshalJSON(b []byte) error {
	var QuerytimeofflimitvaluesrequestMap map[string]interface{}
	err := json.Unmarshal(b, &QuerytimeofflimitvaluesrequestMap)
	if err != nil {
		return err
	}
	
	if TimeOffLimitId, ok := QuerytimeofflimitvaluesrequestMap["timeOffLimitId"].(string); ok {
		o.TimeOffLimitId = &TimeOffLimitId
	}
    
	if ActivityCodeId, ok := QuerytimeofflimitvaluesrequestMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
    
	if DateRanges, ok := QuerytimeofflimitvaluesrequestMap["dateRanges"].([]interface{}); ok {
		DateRangesString, _ := json.Marshal(DateRanges)
		json.Unmarshal(DateRangesString, &o.DateRanges)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Querytimeofflimitvaluesrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
