package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkerrordetail
type Bulkerrordetail struct { 
	// FieldName
	FieldName *string `json:"fieldName,omitempty"`


	// Value
	Value *string `json:"value,omitempty"`


	// Message
	Message *string `json:"message,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bulkerrordetail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
