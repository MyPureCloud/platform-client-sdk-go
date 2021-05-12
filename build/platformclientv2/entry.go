package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Entry
type Entry struct { 
	// Value - A value included in this facet.
	Value *string `json:"value,omitempty"`


	// Count - The number of results with this value.
	Count *int `json:"count,omitempty"`

}

// String returns a JSON representation of the model
func (o *Entry) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
