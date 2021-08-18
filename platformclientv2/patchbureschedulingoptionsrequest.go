package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Patchbureschedulingoptionsrequest
type Patchbureschedulingoptionsrequest struct { 
	// ManagementUnits - Per-management unit rescheduling options to update
	ManagementUnits *[]Patchbureschedulingoptionsmanagementunitrequest `json:"managementUnits,omitempty"`

}

func (u *Patchbureschedulingoptionsrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Patchbureschedulingoptionsrequest

	

	return json.Marshal(&struct { 
		ManagementUnits *[]Patchbureschedulingoptionsmanagementunitrequest `json:"managementUnits,omitempty"`
		*Alias
	}{ 
		ManagementUnits: u.ManagementUnits,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Patchbureschedulingoptionsrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
