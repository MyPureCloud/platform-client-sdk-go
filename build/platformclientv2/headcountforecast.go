package platformclientv2
import (
	"encoding/json"
)

// Headcountforecast - Headcount interval information for schedule
type Headcountforecast struct { 
	// Required - Headcount information with shrinkage
	Required *[]Headcountinterval `json:"required,omitempty"`


	// RequiredWithoutShrinkage - Headcount information without shrinkage
	RequiredWithoutShrinkage *[]Headcountinterval `json:"requiredWithoutShrinkage,omitempty"`

}

// String returns a JSON representation of the model
func (o *Headcountforecast) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
