package platformclientv2
import (
	"encoding/json"
)

// Aggregateviewdata
type Aggregateviewdata struct { 
	// Name
	Name *string `json:"name,omitempty"`


	// Stats
	Stats *Statisticalsummary `json:"stats,omitempty"`

}

// String returns a JSON representation of the model
func (o *Aggregateviewdata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
