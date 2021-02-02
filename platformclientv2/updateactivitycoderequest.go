package platformclientv2
import (
	"encoding/json"
)

// Updateactivitycoderequest - Activity Code
type Updateactivitycoderequest struct { 
	// Name - The name of the activity code
	Name *string `json:"name,omitempty"`


	// Category - The activity code's category. Attempting to change the category of a default activity code will return an error
	Category *string `json:"category,omitempty"`


	// LengthInMinutes - The default length of the activity in minutes
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`


	// CountsAsPaidTime - Whether an agent is paid while performing this activity
	CountsAsPaidTime *bool `json:"countsAsPaidTime,omitempty"`


	// CountsAsWorkTime - Indicates whether or not the activity should be counted as work time
	CountsAsWorkTime *bool `json:"countsAsWorkTime,omitempty"`


	// AgentTimeOffSelectable - Whether an agent can select this activity code when creating or editing a time off request
	AgentTimeOffSelectable *bool `json:"agentTimeOffSelectable,omitempty"`


	// Metadata - Version metadata for the associated business unit's list of activity codes
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

}

// String returns a JSON representation of the model
func (o *Updateactivitycoderequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
