package platformclientv2
import (
	"encoding/json"
)

// Stateventcampaigntopicdatum
type Stateventcampaigntopicdatum struct { 
	// Interval
	Interval *string `json:"interval,omitempty"`


	// Metrics
	Metrics *[]Stateventcampaigntopicmetric `json:"metrics,omitempty"`

}

// String returns a JSON representation of the model
func (o *Stateventcampaigntopicdatum) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
