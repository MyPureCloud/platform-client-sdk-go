package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Scimv2enterpriseuser - Defines a SCIM enterprise user.
type Scimv2enterpriseuser struct { 
	// Division - The division that the user belongs to.
	Division *string `json:"division,omitempty"`


	// Department - The department that the user belongs to.
	Department *string `json:"department,omitempty"`


	// Manager - The user's manager.
	Manager *Manager `json:"manager,omitempty"`


	// EmployeeNumber - The user's employee number.
	EmployeeNumber *string `json:"employeeNumber,omitempty"`

}

func (u *Scimv2enterpriseuser) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Scimv2enterpriseuser

	

	return json.Marshal(&struct { 
		Division *string `json:"division,omitempty"`
		
		Department *string `json:"department,omitempty"`
		
		Manager *Manager `json:"manager,omitempty"`
		
		EmployeeNumber *string `json:"employeeNumber,omitempty"`
		*Alias
	}{ 
		Division: u.Division,
		
		Department: u.Department,
		
		Manager: u.Manager,
		
		EmployeeNumber: u.EmployeeNumber,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Scimv2enterpriseuser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
