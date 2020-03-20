package platformclientv2
import (
	"encoding/json"
)

// Timeoffrequestsettings - Time Off Request Settings
type Timeoffrequestsettings struct { 
	// SubmissionRangeEnforced - Whether to enforce a submission range for agent time off requests
	SubmissionRangeEnforced *bool `json:"submissionRangeEnforced,omitempty"`


	// SubmissionEarliestDaysFromNow - The earliest number of days from now for which an agent can submit a time off request.  Use negative numbers to indicate days in the past
	SubmissionEarliestDaysFromNow *int32 `json:"submissionEarliestDaysFromNow,omitempty"`


	// SubmissionLatestDaysFromNow - The latest number of days from now for which an agent can submit a time off request
	SubmissionLatestDaysFromNow *int32 `json:"submissionLatestDaysFromNow,omitempty"`

}

// String returns a JSON representation of the model
func (o *Timeoffrequestsettings) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
