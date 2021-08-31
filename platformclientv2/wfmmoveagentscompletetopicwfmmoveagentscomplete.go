package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmmoveagentscompletetopicwfmmoveagentscomplete
type Wfmmoveagentscompletetopicwfmmoveagentscomplete struct { 
	// RequestingUser
	RequestingUser *Wfmmoveagentscompletetopicuserreference `json:"requestingUser,omitempty"`


	// DestinationManagementUnit
	DestinationManagementUnit *Wfmmoveagentscompletetopicmanagementunit `json:"destinationManagementUnit,omitempty"`


	// Results
	Results *[]Wfmmoveagentscompletetopicwfmmoveagentdata `json:"results,omitempty"`

}

func (o *Wfmmoveagentscompletetopicwfmmoveagentscomplete) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmmoveagentscompletetopicwfmmoveagentscomplete
	
	return json.Marshal(&struct { 
		RequestingUser *Wfmmoveagentscompletetopicuserreference `json:"requestingUser,omitempty"`
		
		DestinationManagementUnit *Wfmmoveagentscompletetopicmanagementunit `json:"destinationManagementUnit,omitempty"`
		
		Results *[]Wfmmoveagentscompletetopicwfmmoveagentdata `json:"results,omitempty"`
		*Alias
	}{ 
		RequestingUser: o.RequestingUser,
		
		DestinationManagementUnit: o.DestinationManagementUnit,
		
		Results: o.Results,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmmoveagentscompletetopicwfmmoveagentscomplete) UnmarshalJSON(b []byte) error {
	var WfmmoveagentscompletetopicwfmmoveagentscompleteMap map[string]interface{}
	err := json.Unmarshal(b, &WfmmoveagentscompletetopicwfmmoveagentscompleteMap)
	if err != nil {
		return err
	}
	
	if RequestingUser, ok := WfmmoveagentscompletetopicwfmmoveagentscompleteMap["requestingUser"].(map[string]interface{}); ok {
		RequestingUserString, _ := json.Marshal(RequestingUser)
		json.Unmarshal(RequestingUserString, &o.RequestingUser)
	}
	
	if DestinationManagementUnit, ok := WfmmoveagentscompletetopicwfmmoveagentscompleteMap["destinationManagementUnit"].(map[string]interface{}); ok {
		DestinationManagementUnitString, _ := json.Marshal(DestinationManagementUnit)
		json.Unmarshal(DestinationManagementUnitString, &o.DestinationManagementUnit)
	}
	
	if Results, ok := WfmmoveagentscompletetopicwfmmoveagentscompleteMap["results"].([]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmmoveagentscompletetopicwfmmoveagentscomplete) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
