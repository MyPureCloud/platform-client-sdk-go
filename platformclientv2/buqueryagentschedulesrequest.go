package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buqueryagentschedulesrequest
type Buqueryagentschedulesrequest struct { 
	// ManagementUnitId - The ID of the management unit to query
	ManagementUnitId *string `json:"managementUnitId,omitempty"`


	// UserIds - The IDs of the users to query.  Omit to query all user schedules in the management unit. 
	UserIds *[]string `json:"userIds,omitempty"`

}

func (o *Buqueryagentschedulesrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buqueryagentschedulesrequest
	
	return json.Marshal(&struct { 
		ManagementUnitId *string `json:"managementUnitId,omitempty"`
		
		UserIds *[]string `json:"userIds,omitempty"`
		*Alias
	}{ 
		ManagementUnitId: o.ManagementUnitId,
		
		UserIds: o.UserIds,
		Alias:    (*Alias)(o),
	})
}

func (o *Buqueryagentschedulesrequest) UnmarshalJSON(b []byte) error {
	var BuqueryagentschedulesrequestMap map[string]interface{}
	err := json.Unmarshal(b, &BuqueryagentschedulesrequestMap)
	if err != nil {
		return err
	}
	
	if ManagementUnitId, ok := BuqueryagentschedulesrequestMap["managementUnitId"].(string); ok {
		o.ManagementUnitId = &ManagementUnitId
	}
    
	if UserIds, ok := BuqueryagentschedulesrequestMap["userIds"].([]interface{}); ok {
		UserIdsString, _ := json.Marshal(UserIds)
		json.Unmarshal(UserIdsString, &o.UserIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buqueryagentschedulesrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
