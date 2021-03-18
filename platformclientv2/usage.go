package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Usage
type Usage struct { 
	// Types
	Types *[]Usageitem `json:"types,omitempty"`

}

// String returns a JSON representation of the model
func (o *Usage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
