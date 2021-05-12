package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Oauthlasttokenissued
type Oauthlasttokenissued struct { 
	// DateIssued - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateIssued *time.Time `json:"dateIssued,omitempty"`

}

// String returns a JSON representation of the model
func (o *Oauthlasttokenissued) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
