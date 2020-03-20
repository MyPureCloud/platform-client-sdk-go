package platformclientv2
import (
	"encoding/json"
)

// Stateventcampaigntopicstatsnotification
type Stateventcampaigntopicstatsnotification struct { 
	// Group
	Group *map[string]string `json:"group,omitempty"`


	// Data
	Data *[]Stateventcampaigntopicdatum `json:"data,omitempty"`

}

// String returns a JSON representation of the model
func (o *Stateventcampaigntopicstatsnotification) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
