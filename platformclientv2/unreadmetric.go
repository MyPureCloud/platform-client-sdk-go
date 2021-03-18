package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Unreadmetric
type Unreadmetric struct { 
	// Count - The count of unread alerts for a specific rule type.
	Count *int `json:"count,omitempty"`

}

// String returns a JSON representation of the model
func (o *Unreadmetric) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
