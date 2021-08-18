package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Coachingappointmentstatusrequest
type Coachingappointmentstatusrequest struct { 
	// Status - The status of the coaching appointment
	Status *string `json:"status,omitempty"`

}

func (u *Coachingappointmentstatusrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Coachingappointmentstatusrequest

	

	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		*Alias
	}{ 
		Status: u.Status,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Coachingappointmentstatusrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
