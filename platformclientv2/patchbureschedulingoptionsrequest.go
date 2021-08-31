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

func (o *Patchbureschedulingoptionsrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Patchbureschedulingoptionsrequest
	
	return json.Marshal(&struct { 
		ManagementUnits *[]Patchbureschedulingoptionsmanagementunitrequest `json:"managementUnits,omitempty"`
		*Alias
	}{ 
		ManagementUnits: o.ManagementUnits,
		Alias:    (*Alias)(o),
	})
}

func (o *Patchbureschedulingoptionsrequest) UnmarshalJSON(b []byte) error {
	var PatchbureschedulingoptionsrequestMap map[string]interface{}
	err := json.Unmarshal(b, &PatchbureschedulingoptionsrequestMap)
	if err != nil {
		return err
	}
	
	if ManagementUnits, ok := PatchbureschedulingoptionsrequestMap["managementUnits"].([]interface{}); ok {
		ManagementUnitsString, _ := json.Marshal(ManagementUnits)
		json.Unmarshal(ManagementUnitsString, &o.ManagementUnits)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Patchbureschedulingoptionsrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
