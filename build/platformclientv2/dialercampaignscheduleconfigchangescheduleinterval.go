package platformclientv2
import (
	"encoding/json"
)

// Dialercampaignscheduleconfigchangescheduleinterval
type Dialercampaignscheduleconfigchangescheduleinterval struct { 
	// Start
	Start *string `json:"start,omitempty"`


	// End
	End *string `json:"end,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialercampaignscheduleconfigchangescheduleinterval) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
