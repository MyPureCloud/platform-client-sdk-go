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

func (o *Wfmmovemanagementunittopicmovemanagementunitnotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmmovemanagementunittopicmovemanagementunitnotification
	
	return json.Marshal(&struct { 
		BusinessUnit *Wfmmovemanagementunittopicbusinessunit `json:"businessUnit,omitempty"`
		
		Status *string `json:"status,omitempty"`
		*Alias
	}{ 
		BusinessUnit: o.BusinessUnit,
		
		Status: o.Status,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmmovemanagementunittopicmovemanagementunitnotification) UnmarshalJSON(b []byte) error {
	var WfmmovemanagementunittopicmovemanagementunitnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &WfmmovemanagementunittopicmovemanagementunitnotificationMap)
	if err != nil {
		return err
	}
	
	if BusinessUnit, ok := WfmmovemanagementunittopicmovemanagementunitnotificationMap["businessUnit"].(map[string]interface{}); ok {
		BusinessUnitString, _ := json.Marshal(BusinessUnit)
		json.Unmarshal(BusinessUnitString, &o.BusinessUnit)
	}
	
	if Status, ok := WfmmovemanagementunittopicmovemanagementunitnotificationMap["status"].(string); ok {
		o.Status = &Status
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmmovemanagementunittopicmovemanagementunitnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
