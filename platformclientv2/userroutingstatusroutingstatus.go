package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userroutingstatusroutingstatus
type Userroutingstatusroutingstatus struct { 
	// Status
	Status *string `json:"status,omitempty"`


	// StartTime
	StartTime *time.Time `json:"startTime,omitempty"`

}

func (u *Userroutingstatusroutingstatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userroutingstatusroutingstatus

	
	StartTime := new(string)
	if u.StartTime != nil {
		
		*StartTime = timeutil.Strftime(u.StartTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartTime = nil
	}
	

	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		
		StartTime *string `json:"startTime,omitempty"`
		*Alias
	}{ 
		Status: u.Status,
		
		StartTime: StartTime,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Userroutingstatusroutingstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
