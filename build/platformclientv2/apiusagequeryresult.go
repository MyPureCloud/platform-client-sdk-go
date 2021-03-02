package platformclientv2
import (
	"encoding/json"
)

// Apiusagequeryresult
type Apiusagequeryresult struct { 
	// Results - Query results
	Results *[]Apiusagerow `json:"results,omitempty"`


	// QueryStatus - Query status
	QueryStatus *string `json:"queryStatus,omitempty"`

}

// String returns a JSON representation of the model
func (o *Apiusagequeryresult) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
