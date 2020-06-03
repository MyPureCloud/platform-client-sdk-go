package platformclientv2
import (
	"encoding/json"
)

// Maxlength
type Maxlength struct { 
	// Min - A non-negative integer for a text-based schema field denoting the minimum largest length string the field can contain for a schema instance.
	Min *int64 `json:"min,omitempty"`


	// Max - A non-negative integer for a text-based schema field denoting the maximum largest string the field can contain for a schema instance.
	Max *int64 `json:"max,omitempty"`

}

// String returns a JSON representation of the model
func (o *Maxlength) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
