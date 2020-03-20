package platformclientv2
import (
	"encoding/json"
)

// Routegroup - Route group for calculated forecasts
type Routegroup struct { 
	// Attributes - The attributes that describe this route group
	Attributes *Routegroupattributes `json:"attributes,omitempty"`


	// OfferedPerInterval - Interactions offered per 15 minute interval
	OfferedPerInterval *[]float64 `json:"offeredPerInterval,omitempty"`


	// AverageTalkTimeSecondsPerInterval - Average talk time in seconds per 15 minute interval
	AverageTalkTimeSecondsPerInterval *[]float64 `json:"averageTalkTimeSecondsPerInterval,omitempty"`


	// AverageAfterCallWorkSecondsPerInterval - Average after call work in seconds per 15 minute interval
	AverageAfterCallWorkSecondsPerInterval *[]float64 `json:"averageAfterCallWorkSecondsPerInterval,omitempty"`


	// CompletedPerInterval - Interactions completed per 15 minute interval
	CompletedPerInterval *[]float64 `json:"completedPerInterval,omitempty"`


	// AbandonedPerInterval - Interactions abandoned per 15 minute interval
	AbandonedPerInterval *[]float64 `json:"abandonedPerInterval,omitempty"`

}

// String returns a JSON representation of the model
func (o *Routegroup) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
