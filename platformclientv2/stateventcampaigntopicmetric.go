package platformclientv2
import (
	"encoding/json"
)

// Stateventcampaigntopicmetric
type Stateventcampaigntopicmetric struct { 
	// Metric
	Metric *string `json:"metric,omitempty"`


	// Qualifier
	Qualifier *string `json:"qualifier,omitempty"`


	// Stats
	Stats *map[string]float32 `json:"stats,omitempty"`

}

// String returns a JSON representation of the model
func (o *Stateventcampaigntopicmetric) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
