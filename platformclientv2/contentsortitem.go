package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Contentsortitem
type Contentsortitem struct { 
	// Name
	Name *string `json:"name,omitempty"`


	// Ascending
	Ascending *bool `json:"ascending,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contentsortitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
