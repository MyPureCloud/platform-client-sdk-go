package platformclientv2
import (
	"encoding/json"
)

// Shiftstartvariance - Variance in minutes among start times of shifts in work plan
type Shiftstartvariance struct { 
	// ApplicableDays - Days for which shift start variance is configured
	ApplicableDays *[]string `json:"applicableDays,omitempty"`


	// MaxShiftStartVarianceMinutes - Maximum variance in minutes across shift starts
	MaxShiftStartVarianceMinutes *int32 `json:"maxShiftStartVarianceMinutes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Shiftstartvariance) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
