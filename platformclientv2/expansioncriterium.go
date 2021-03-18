package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Expansioncriterium
type Expansioncriterium struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// Threshold
	Threshold *float64 `json:"threshold,omitempty"`

}

// String returns a JSON representation of the model
func (o *Expansioncriterium) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
