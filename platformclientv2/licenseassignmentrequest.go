package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Licenseassignmentrequest
type Licenseassignmentrequest struct { 
	// LicenseId - The id of the license to assign/unassign.
	LicenseId *string `json:"licenseId,omitempty"`


	// UserIdsAdd - The ids of users to assign this license to.
	UserIdsAdd *[]string `json:"userIdsAdd,omitempty"`


	// UserIdsRemove - The ids of users to unassign this license from.
	UserIdsRemove *[]string `json:"userIdsRemove,omitempty"`

}

func (o *Licenseassignmentrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Licenseassignmentrequest
	
	return json.Marshal(&struct { 
		LicenseId *string `json:"licenseId,omitempty"`
		
		UserIdsAdd *[]string `json:"userIdsAdd,omitempty"`
		
		UserIdsRemove *[]string `json:"userIdsRemove,omitempty"`
		*Alias
	}{ 
		LicenseId: o.LicenseId,
		
		UserIdsAdd: o.UserIdsAdd,
		
		UserIdsRemove: o.UserIdsRemove,
		Alias:    (*Alias)(o),
	})
}

func (o *Licenseassignmentrequest) UnmarshalJSON(b []byte) error {
	var LicenseassignmentrequestMap map[string]interface{}
	err := json.Unmarshal(b, &LicenseassignmentrequestMap)
	if err != nil {
		return err
	}
	
	if LicenseId, ok := LicenseassignmentrequestMap["licenseId"].(string); ok {
		o.LicenseId = &LicenseId
	}
	
	if UserIdsAdd, ok := LicenseassignmentrequestMap["userIdsAdd"].([]interface{}); ok {
		UserIdsAddString, _ := json.Marshal(UserIdsAdd)
		json.Unmarshal(UserIdsAddString, &o.UserIdsAdd)
	}
	
	if UserIdsRemove, ok := LicenseassignmentrequestMap["userIdsRemove"].([]interface{}); ok {
		UserIdsRemoveString, _ := json.Marshal(UserIdsRemove)
		json.Unmarshal(UserIdsRemoveString, &o.UserIdsRemove)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Licenseassignmentrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
