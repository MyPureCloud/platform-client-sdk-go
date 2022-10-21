package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Timeoffbalancerequest
type Timeoffbalancerequest struct { 
	// ActivityCodeIds - The set of activity code IDs for which to query available time off balances
	ActivityCodeIds *[]string `json:"activityCodeIds,omitempty"`


	// DateRanges - The list of date ranges for which to query time off balance
	DateRanges *[]Localdaterange `json:"dateRanges,omitempty"`

}

func (o *Timeoffbalancerequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Timeoffbalancerequest
	
	return json.Marshal(&struct { 
		ActivityCodeIds *[]string `json:"activityCodeIds,omitempty"`
		
		DateRanges *[]Localdaterange `json:"dateRanges,omitempty"`
		*Alias
	}{ 
		ActivityCodeIds: o.ActivityCodeIds,
		
		DateRanges: o.DateRanges,
		Alias:    (*Alias)(o),
	})
}

func (o *Timeoffbalancerequest) UnmarshalJSON(b []byte) error {
	var TimeoffbalancerequestMap map[string]interface{}
	err := json.Unmarshal(b, &TimeoffbalancerequestMap)
	if err != nil {
		return err
	}
	
	if ActivityCodeIds, ok := TimeoffbalancerequestMap["activityCodeIds"].([]interface{}); ok {
		ActivityCodeIdsString, _ := json.Marshal(ActivityCodeIds)
		json.Unmarshal(ActivityCodeIdsString, &o.ActivityCodeIds)
	}
	
	if DateRanges, ok := TimeoffbalancerequestMap["dateRanges"].([]interface{}); ok {
		DateRangesString, _ := json.Marshal(DateRanges)
		json.Unmarshal(DateRangesString, &o.DateRanges)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Timeoffbalancerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
