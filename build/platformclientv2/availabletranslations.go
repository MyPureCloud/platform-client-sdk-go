package platformclientv2
import (
	"encoding/json"
)

// Availabletranslations
type Availabletranslations struct { 
	// OrgSpecific
	OrgSpecific *[]string `json:"orgSpecific,omitempty"`


	// Builtin
	Builtin *[]string `json:"builtin,omitempty"`

}

// String returns a JSON representation of the model
func (o *Availabletranslations) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
