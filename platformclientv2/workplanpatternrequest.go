package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Workplanpatternrequest
type Workplanpatternrequest struct { 
	// WorkPlanIds - List of work plan IDs in order of rotation on a weekly basis. Values in the list cannot be null or empty
	WorkPlanIds *[]string `json:"workPlanIds,omitempty"`

}

func (o *Workplanpatternrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Workplanpatternrequest
	
	return json.Marshal(&struct { 
		WorkPlanIds *[]string `json:"workPlanIds,omitempty"`
		*Alias
	}{ 
		WorkPlanIds: o.WorkPlanIds,
		Alias:    (*Alias)(o),
	})
}

func (o *Workplanpatternrequest) UnmarshalJSON(b []byte) error {
	var WorkplanpatternrequestMap map[string]interface{}
	err := json.Unmarshal(b, &WorkplanpatternrequestMap)
	if err != nil {
		return err
	}
	
	if WorkPlanIds, ok := WorkplanpatternrequestMap["workPlanIds"].([]interface{}); ok {
		WorkPlanIdsString, _ := json.Marshal(WorkPlanIds)
		json.Unmarshal(WorkPlanIdsString, &o.WorkPlanIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Workplanpatternrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
