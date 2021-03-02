package platformclientv2
import (
	"encoding/json"
)

// Userroutingstatususerparam
type Userroutingstatususerparam struct { 
	// Key
	Key *string `json:"key,omitempty"`


	// Value
	Value *string `json:"value,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userroutingstatususerparam) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
