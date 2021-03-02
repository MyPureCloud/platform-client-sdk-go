package platformclientv2
import (
	"encoding/json"
)

// Forecastplanninggroupsresponse
type Forecastplanninggroupsresponse struct { 
	// Entities
	Entities *[]Forecastplanninggroupresponse `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Forecastplanninggroupsresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
