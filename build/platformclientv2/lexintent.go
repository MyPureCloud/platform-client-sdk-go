package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Lexintent
type Lexintent struct { 
	// Name - The intent name
	Name *string `json:"name,omitempty"`


	// Description - A description of the intent
	Description *string `json:"description,omitempty"`


	// Slots - An object mapping slot names to Slot objects
	Slots *map[string]Lexslot `json:"slots,omitempty"`


	// Version - The intent version
	Version *string `json:"version,omitempty"`

}

// String returns a JSON representation of the model
func (o *Lexintent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
