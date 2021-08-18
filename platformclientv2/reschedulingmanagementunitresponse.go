package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Reschedulingmanagementunitresponse
type Reschedulingmanagementunitresponse struct { 
	// ManagementUnit - The management unit
	ManagementUnit *Managementunitreference `json:"managementUnit,omitempty"`


	// Applied - Whether the rescheduling run is applied for the given management unit
	Applied *bool `json:"applied,omitempty"`

}

func (u *Reschedulingmanagementunitresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Reschedulingmanagementunitresponse

	

	return json.Marshal(&struct { 
		ManagementUnit *Managementunitreference `json:"managementUnit,omitempty"`
		
		Applied *bool `json:"applied,omitempty"`
		*Alias
	}{ 
		ManagementUnit: u.ManagementUnit,
		
		Applied: u.Applied,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Reschedulingmanagementunitresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
