package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Callrecord
type Callrecord struct { 
	// LastAttempt - Timestamp of the last attempt to reach this number. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	LastAttempt *time.Time `json:"lastAttempt,omitempty"`


	// LastResult - Result of the last attempt to reach this number
	LastResult *string `json:"lastResult,omitempty"`

}

// String returns a JSON representation of the model
func (o *Callrecord) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
