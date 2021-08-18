package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Provisioninfo
type Provisioninfo struct { 
	// Time - The time at which this phone was provisioned. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Time *time.Time `json:"time,omitempty"`


	// Source - The source of the provisioning
	Source *string `json:"source,omitempty"`


	// ErrorInfo - The error information from the provision process, if any
	ErrorInfo *string `json:"errorInfo,omitempty"`

}

func (u *Provisioninfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Provisioninfo

	
	Time := new(string)
	if u.Time != nil {
		
		*Time = timeutil.Strftime(u.Time, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Time = nil
	}
	

	return json.Marshal(&struct { 
		Time *string `json:"time,omitempty"`
		
		Source *string `json:"source,omitempty"`
		
		ErrorInfo *string `json:"errorInfo,omitempty"`
		*Alias
	}{ 
		Time: Time,
		
		Source: u.Source,
		
		ErrorInfo: u.ErrorInfo,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Provisioninfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
