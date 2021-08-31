package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Timeoffrequestsettings - Time Off Request Settings
type Timeoffrequestsettings struct { 
	// SubmissionRangeEnforced - Whether to enforce a submission range for agent time off requests
	SubmissionRangeEnforced *bool `json:"submissionRangeEnforced,omitempty"`


	// SubmissionEarliestDaysFromNow - The earliest number of days from now for which an agent can submit a time off request.  Use negative numbers to indicate days in the past
	SubmissionEarliestDaysFromNow *int `json:"submissionEarliestDaysFromNow,omitempty"`


	// SubmissionLatestDaysFromNow - The latest number of days from now for which an agent can submit a time off request
	SubmissionLatestDaysFromNow *int `json:"submissionLatestDaysFromNow,omitempty"`

}

func (o *Timeoffrequestsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Timeoffrequestsettings
	
	return json.Marshal(&struct { 
		SubmissionRangeEnforced *bool `json:"submissionRangeEnforced,omitempty"`
		
		SubmissionEarliestDaysFromNow *int `json:"submissionEarliestDaysFromNow,omitempty"`
		
		SubmissionLatestDaysFromNow *int `json:"submissionLatestDaysFromNow,omitempty"`
		*Alias
	}{ 
		SubmissionRangeEnforced: o.SubmissionRangeEnforced,
		
		SubmissionEarliestDaysFromNow: o.SubmissionEarliestDaysFromNow,
		
		SubmissionLatestDaysFromNow: o.SubmissionLatestDaysFromNow,
		Alias:    (*Alias)(o),
	})
}

func (o *Timeoffrequestsettings) UnmarshalJSON(b []byte) error {
	var TimeoffrequestsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &TimeoffrequestsettingsMap)
	if err != nil {
		return err
	}
	
	if SubmissionRangeEnforced, ok := TimeoffrequestsettingsMap["submissionRangeEnforced"].(bool); ok {
		o.SubmissionRangeEnforced = &SubmissionRangeEnforced
	}
	
	if SubmissionEarliestDaysFromNow, ok := TimeoffrequestsettingsMap["submissionEarliestDaysFromNow"].(float64); ok {
		SubmissionEarliestDaysFromNowInt := int(SubmissionEarliestDaysFromNow)
		o.SubmissionEarliestDaysFromNow = &SubmissionEarliestDaysFromNowInt
	}
	
	if SubmissionLatestDaysFromNow, ok := TimeoffrequestsettingsMap["submissionLatestDaysFromNow"].(float64); ok {
		SubmissionLatestDaysFromNowInt := int(SubmissionLatestDaysFromNow)
		o.SubmissionLatestDaysFromNow = &SubmissionLatestDaysFromNowInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Timeoffrequestsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
