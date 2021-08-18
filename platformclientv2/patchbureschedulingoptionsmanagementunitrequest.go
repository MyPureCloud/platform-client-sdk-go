package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Patchbureschedulingoptionsmanagementunitrequest
type Patchbureschedulingoptionsmanagementunitrequest struct { 
	// ManagementUnitId - The management unit portion of the rescheduling run to update
	ManagementUnitId *string `json:"managementUnitId,omitempty"`


	// Applied - Whether to mark the run as applied.  Only applies to reschedule runs.  Once applied, a run cannot be un-marked as applied
	Applied *bool `json:"applied,omitempty"`

}

func (u *Patchbureschedulingoptionsmanagementunitrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Patchbureschedulingoptionsmanagementunitrequest

	

	return json.Marshal(&struct { 
		ManagementUnitId *string `json:"managementUnitId,omitempty"`
		
		Applied *bool `json:"applied,omitempty"`
		*Alias
	}{ 
		ManagementUnitId: u.ManagementUnitId,
		
		Applied: u.Applied,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Patchbureschedulingoptionsmanagementunitrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
