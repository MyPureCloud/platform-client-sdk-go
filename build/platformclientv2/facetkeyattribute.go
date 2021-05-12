package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Facetkeyattribute
type Facetkeyattribute struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Count
	Count *int `json:"count,omitempty"`

}

// String returns a JSON representation of the model
func (o *Facetkeyattribute) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
