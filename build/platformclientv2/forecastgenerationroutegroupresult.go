package platformclientv2
import (
	"encoding/json"
)

// Forecastgenerationroutegroupresult
type Forecastgenerationroutegroupresult struct { 
	// RouteGroup - The route group this result represents
	RouteGroup *Routegroupattributes `json:"routeGroup,omitempty"`


	// MetricResults - The generation results for the associated route group
	MetricResults *[]Forecasttimeseriesresult `json:"metricResults,omitempty"`

}

// String returns a JSON representation of the model
func (o *Forecastgenerationroutegroupresult) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
