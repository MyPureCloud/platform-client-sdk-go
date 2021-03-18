package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Replacerequest
type Replacerequest struct { 
	// ChangeNumber
	ChangeNumber *int `json:"changeNumber,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// AuthToken
	AuthToken *string `json:"authToken,omitempty"`

}

// String returns a JSON representation of the model
func (o *Replacerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
