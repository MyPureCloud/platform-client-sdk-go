package platformclientv2
import (
	"encoding/json"
)

// Buqueryagentschedulesrequest
type Buqueryagentschedulesrequest struct { 
	// ManagementUnitId - The ID of the management unit to query
	ManagementUnitId *string `json:"managementUnitId,omitempty"`


	// UserIds - The IDs of the users to query.  Omit to query all user schedules in the management unit. Note: Only one of [teamIds, userIds] can be requested
	UserIds *[]string `json:"userIds,omitempty"`


	// TeamIds - The teamIds to report on. If null or not set, results will be queried for requested users if applicable or otherwise all users in the management unit. Note: Only one of [teamIds, userIds] can be requested
	TeamIds *[]string `json:"teamIds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Buqueryagentschedulesrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
