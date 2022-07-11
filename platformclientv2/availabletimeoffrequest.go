package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Availabletimeoffrequest
type Availabletimeoffrequest struct { 
	// ActivityCodeId - The ID for activity code to query available time off minutes
	ActivityCodeId *string `json:"activityCodeId,omitempty"`


	// DateRanges - A list of date ranges of available time off minutes. A maximum number of date ranges is 30. The maximum total number of days in all ranges is 366. If no ranges are specified, then only the presence of the associated time off limit object will be checked. In such case, if the association exists, then the response will contain a list with of a single element filled with timeOffLimitId only.
	DateRanges *[]Localdaterange `json:"dateRanges,omitempty"`

}

func (o *Availabletimeoffrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Availabletimeoffrequest
	
	return json.Marshal(&struct { 
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		DateRanges *[]Localdaterange `json:"dateRanges,omitempty"`
		*Alias
	}{ 
		ActivityCodeId: o.ActivityCodeId,
		
		DateRanges: o.DateRanges,
		Alias:    (*Alias)(o),
	})
}

func (o *Availabletimeoffrequest) UnmarshalJSON(b []byte) error {
	var AvailabletimeoffrequestMap map[string]interface{}
	err := json.Unmarshal(b, &AvailabletimeoffrequestMap)
	if err != nil {
		return err
	}
	
	if ActivityCodeId, ok := AvailabletimeoffrequestMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
    
	if DateRanges, ok := AvailabletimeoffrequestMap["dateRanges"].([]interface{}); ok {
		DateRangesString, _ := json.Marshal(DateRanges)
		json.Unmarshal(DateRangesString, &o.DateRanges)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Availabletimeoffrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
