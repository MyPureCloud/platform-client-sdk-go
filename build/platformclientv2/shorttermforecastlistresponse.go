package platformclientv2
import (
	"encoding/json"
)

// Shorttermforecastlistresponse
type Shorttermforecastlistresponse struct { 
	// Entities
	Entities *[]Shorttermforecastlistitemresponse `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Shorttermforecastlistresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
