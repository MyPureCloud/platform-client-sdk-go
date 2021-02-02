package platformclientv2
import (
	"encoding/json"
)

// Trunkinstancetopictrunkerrorinfo
type Trunkinstancetopictrunkerrorinfo struct { 
	// Text
	Text *string `json:"text,omitempty"`


	// Code
	Code *string `json:"code,omitempty"`


	// Details
	Details *Trunkinstancetopictrunkerrorinfodetails `json:"details,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trunkinstancetopictrunkerrorinfo) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
