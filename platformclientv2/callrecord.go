package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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

func (u *Callrecord) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Callrecord

	
	LastAttempt := new(string)
	if u.LastAttempt != nil {
		
		*LastAttempt = timeutil.Strftime(u.LastAttempt, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		LastAttempt = nil
	}
	

	return json.Marshal(&struct { 
		LastAttempt *string `json:"lastAttempt,omitempty"`
		
		LastResult *string `json:"lastResult,omitempty"`
		*Alias
	}{ 
		LastAttempt: LastAttempt,
		
		LastResult: u.LastResult,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Callrecord) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
