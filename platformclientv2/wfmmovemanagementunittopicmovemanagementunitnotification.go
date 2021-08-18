package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmmovemanagementunittopicmovemanagementunitnotification
type Wfmmovemanagementunittopicmovemanagementunitnotification struct { 
	// BusinessUnit
	BusinessUnit *Wfmmovemanagementunittopicbusinessunit `json:"businessUnit,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`

}

func (u *Wfmmovemanagementunittopicmovemanagementunitnotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmmovemanagementunittopicmovemanagementunitnotification

	

	return json.Marshal(&struct { 
		BusinessUnit *Wfmmovemanagementunittopicbusinessunit `json:"businessUnit,omitempty"`
		
		Status *string `json:"status,omitempty"`
		*Alias
	}{ 
		BusinessUnit: u.BusinessUnit,
		
		Status: u.Status,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmmovemanagementunittopicmovemanagementunitnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
