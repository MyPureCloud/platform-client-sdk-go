package platformclientv2
import (
	"encoding/json"
)

// Forecastgenerationresult
type Forecastgenerationresult struct { 
	// RouteGroupResults - Generation results, broken down by route group
	RouteGroupResults *[]Forecastgenerationroutegroupresult `json:"routeGroupResults,omitempty"`

}

// String returns a JSON representation of the model
func (o *Forecastgenerationresult) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
