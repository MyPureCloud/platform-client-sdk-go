package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Record
type Record struct { 
	// Name - The name of the record.
	Name *string `json:"name,omitempty"`


	// VarType - The type of the record. (Example values:  MX, TXT, CNAME)
	VarType *string `json:"type,omitempty"`


	// Value - The value of the record.
	Value *string `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Record) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
