package platformclientv2
import (
	"encoding/json"
)

// Schedulingsettingsresponse - Scheduling Settings
type Schedulingsettingsresponse struct { 
	// MaxOccupancyPercentForDeferredWork - Max occupancy percent for deferred work
	MaxOccupancyPercentForDeferredWork *int32 `json:"maxOccupancyPercentForDeferredWork,omitempty"`


	// DefaultShrinkagePercent - Default shrinkage percent for scheduling
	DefaultShrinkagePercent *float64 `json:"defaultShrinkagePercent,omitempty"`


	// ShrinkageOverrides - Shrinkage overrides for scheduling
	ShrinkageOverrides *Shrinkageoverrides `json:"shrinkageOverrides,omitempty"`

}

// String returns a JSON representation of the model
func (o *Schedulingsettingsresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
