package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Intent
type Intent struct { 
	// Name
	Name *string `json:"name,omitempty"`

}

// String returns a JSON representation of the model
func (o *Intent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
