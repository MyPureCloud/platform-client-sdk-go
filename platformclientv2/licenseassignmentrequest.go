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

func (u *Licenseassignmentrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Licenseassignmentrequest

	

	return json.Marshal(&struct { 
		LicenseId *string `json:"licenseId,omitempty"`
		
		UserIdsAdd *[]string `json:"userIdsAdd,omitempty"`
		
		UserIdsRemove *[]string `json:"userIdsRemove,omitempty"`
		*Alias
	}{ 
		LicenseId: u.LicenseId,
		
		UserIdsAdd: u.UserIdsAdd,
		
		UserIdsRemove: u.UserIdsRemove,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Licenseassignmentrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
