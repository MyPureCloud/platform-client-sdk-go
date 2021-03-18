package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Permissions
type Permissions struct { 
	// Ids - List of permission ids.
	Ids *[]string `json:"ids,omitempty"`

}

// String returns a JSON representation of the model
func (o *Permissions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
