package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Calltarget
type Calltarget struct { 
	// VarType - The type of call
	VarType *string `json:"type,omitempty"`


	// Value - The id of the station or an E.164 formatted phone number
	Value *string `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Calltarget) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
