package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Screenrecordingsessionrequest
type Screenrecordingsessionrequest struct { 
	// State - The screen recording session's state.  Values can be: 'stopped'
	State *string `json:"state,omitempty"`


	// ArchiveDate - The screen recording session's archive date. Must be greater than 1 day from now if set. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ArchiveDate *time.Time `json:"archiveDate,omitempty"`


	// DeleteDate - The screen recording session's delete date. Must be greater than archiveDate if set, otherwise one day from now. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DeleteDate *time.Time `json:"deleteDate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Screenrecordingsessionrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
