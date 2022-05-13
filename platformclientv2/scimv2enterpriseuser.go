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

func (o *Scimv2enterpriseuser) MarshalJSON() ([]byte, error) {
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
		Division: o.Division,
		
		Department: o.Department,
		
		Manager: o.Manager,
		
		EmployeeNumber: o.EmployeeNumber,
		Alias:    (*Alias)(o),
	})
}

func (o *Scimv2enterpriseuser) UnmarshalJSON(b []byte) error {
	var Scimv2enterpriseuserMap map[string]interface{}
	err := json.Unmarshal(b, &Scimv2enterpriseuserMap)
	if err != nil {
		return err
	}
	
	if Division, ok := Scimv2enterpriseuserMap["division"].(string); ok {
		o.Division = &Division
	}
    
	if Department, ok := Scimv2enterpriseuserMap["department"].(string); ok {
		o.Department = &Department
	}
    
	if Manager, ok := Scimv2enterpriseuserMap["manager"].(map[string]interface{}); ok {
		ManagerString, _ := json.Marshal(Manager)
		json.Unmarshal(ManagerString, &o.Manager)
	}
	
	if EmployeeNumber, ok := Scimv2enterpriseuserMap["employeeNumber"].(string); ok {
		o.EmployeeNumber = &EmployeeNumber
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Scimv2enterpriseuser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
