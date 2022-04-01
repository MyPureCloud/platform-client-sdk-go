package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Processscheduleupdateuploadrequest
type Processscheduleupdateuploadrequest struct { 
	// UploadKey - The uploadKey provided by the request to get an upload URL
	UploadKey *string `json:"uploadKey,omitempty"`


	// TeamIds - The list of teams to which the users being modified belong. Only required if the requesting user has conditional permission to wfm:schedule:edit
	TeamIds *[]string `json:"teamIds,omitempty"`


	// ManagementUnitIdsForAddedTeamUsers - The set of muIds to which agents belong if agents are being newly added to the schedule, if the requesting user has conditional permission to wfm:schedule:edit
	ManagementUnitIdsForAddedTeamUsers *[]string `json:"managementUnitIdsForAddedTeamUsers,omitempty"`

}

func (o *Processscheduleupdateuploadrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Processscheduleupdateuploadrequest
	
	return json.Marshal(&struct { 
		UploadKey *string `json:"uploadKey,omitempty"`
		
		TeamIds *[]string `json:"teamIds,omitempty"`
		
		ManagementUnitIdsForAddedTeamUsers *[]string `json:"managementUnitIdsForAddedTeamUsers,omitempty"`
		*Alias
	}{ 
		UploadKey: o.UploadKey,
		
		TeamIds: o.TeamIds,
		
		ManagementUnitIdsForAddedTeamUsers: o.ManagementUnitIdsForAddedTeamUsers,
		Alias:    (*Alias)(o),
	})
}

func (o *Processscheduleupdateuploadrequest) UnmarshalJSON(b []byte) error {
	var ProcessscheduleupdateuploadrequestMap map[string]interface{}
	err := json.Unmarshal(b, &ProcessscheduleupdateuploadrequestMap)
	if err != nil {
		return err
	}
	
	if UploadKey, ok := ProcessscheduleupdateuploadrequestMap["uploadKey"].(string); ok {
		o.UploadKey = &UploadKey
	}
	
	if TeamIds, ok := ProcessscheduleupdateuploadrequestMap["teamIds"].([]interface{}); ok {
		TeamIdsString, _ := json.Marshal(TeamIds)
		json.Unmarshal(TeamIdsString, &o.TeamIds)
	}
	
	if ManagementUnitIdsForAddedTeamUsers, ok := ProcessscheduleupdateuploadrequestMap["managementUnitIdsForAddedTeamUsers"].([]interface{}); ok {
		ManagementUnitIdsForAddedTeamUsersString, _ := json.Marshal(ManagementUnitIdsForAddedTeamUsers)
		json.Unmarshal(ManagementUnitIdsForAddedTeamUsersString, &o.ManagementUnitIdsForAddedTeamUsers)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Processscheduleupdateuploadrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
