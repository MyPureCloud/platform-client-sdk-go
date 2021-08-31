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

func (o *Patchbureschedulingoptionsmanagementunitrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Patchbureschedulingoptionsmanagementunitrequest
	
	return json.Marshal(&struct { 
		ManagementUnitId *string `json:"managementUnitId,omitempty"`
		
		Applied *bool `json:"applied,omitempty"`
		*Alias
	}{ 
		ManagementUnitId: o.ManagementUnitId,
		
		Applied: o.Applied,
		Alias:    (*Alias)(o),
	})
}

func (o *Patchbureschedulingoptionsmanagementunitrequest) UnmarshalJSON(b []byte) error {
	var PatchbureschedulingoptionsmanagementunitrequestMap map[string]interface{}
	err := json.Unmarshal(b, &PatchbureschedulingoptionsmanagementunitrequestMap)
	if err != nil {
		return err
	}
	
	if ManagementUnitId, ok := PatchbureschedulingoptionsmanagementunitrequestMap["managementUnitId"].(string); ok {
		o.ManagementUnitId = &ManagementUnitId
	}
	
	if Applied, ok := PatchbureschedulingoptionsmanagementunitrequestMap["applied"].(bool); ok {
		o.Applied = &Applied
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Patchbureschedulingoptionsmanagementunitrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
