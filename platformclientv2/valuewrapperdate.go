package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Valuewrapperdate - An object to provide context to nullable fields in PATCH requests
type Valuewrapperdate struct { 
	// Value - The value for the associated field. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Value *time.Time `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Valuewrapperdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
