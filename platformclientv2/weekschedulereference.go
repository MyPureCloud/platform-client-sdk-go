package platformclientv2
import (
	"encoding/json"
)

// Weekschedulereference
type Weekschedulereference struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`


	// WeekDate - First day of this week schedule in yyyy-MM-dd format
	WeekDate *string `json:"weekDate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Weekschedulereference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
