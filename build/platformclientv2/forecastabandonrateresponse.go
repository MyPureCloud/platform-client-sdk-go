package platformclientv2
import (
	"encoding/json"
)

// Forecastabandonrateresponse
type Forecastabandonrateresponse struct { 
	// Percent - The target percent abandon rate goal
	Percent *int `json:"percent,omitempty"`

}

// String returns a JSON representation of the model
func (o *Forecastabandonrateresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
