package platformclientv2
import (
	"encoding/json"
)

// Documentattribute
type Documentattribute struct { 
	// Attribute
	Attribute *Attribute `json:"attribute,omitempty"`


	// Values
	Values *[]string `json:"values,omitempty"`

}

// String returns a JSON representation of the model
func (o *Documentattribute) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
