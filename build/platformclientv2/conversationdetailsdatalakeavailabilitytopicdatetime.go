package platformclientv2
import (
	"encoding/json"
)

// Conversationdetailsdatalakeavailabilitytopicdatetime
type Conversationdetailsdatalakeavailabilitytopicdatetime struct { 
	// IMillis
	IMillis *int `json:"iMillis,omitempty"`


	// BeforeNow
	BeforeNow *bool `json:"beforeNow,omitempty"`


	// EqualNow
	EqualNow *bool `json:"equalNow,omitempty"`


	// AfterNow
	AfterNow *bool `json:"afterNow,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationdetailsdatalakeavailabilitytopicdatetime) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
