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

func (o *Callrecord) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Callrecord
	
	LastAttempt := new(string)
	if o.LastAttempt != nil {
		
		*LastAttempt = timeutil.Strftime(o.LastAttempt, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		LastAttempt = nil
	}
	
	return json.Marshal(&struct { 
		LastAttempt *string `json:"lastAttempt,omitempty"`
		
		LastResult *string `json:"lastResult,omitempty"`
		*Alias
	}{ 
		LastAttempt: LastAttempt,
		
		LastResult: o.LastResult,
		Alias:    (*Alias)(o),
	})
}

func (o *Callrecord) UnmarshalJSON(b []byte) error {
	var CallrecordMap map[string]interface{}
	err := json.Unmarshal(b, &CallrecordMap)
	if err != nil {
		return err
	}
	
	if lastAttemptString, ok := CallrecordMap["lastAttempt"].(string); ok {
		LastAttempt, _ := time.Parse("2006-01-02T15:04:05.999999Z", lastAttemptString)
		o.LastAttempt = &LastAttempt
	}
	
	if LastResult, ok := CallrecordMap["lastResult"].(string); ok {
		o.LastResult = &LastResult
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Callrecord) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
