package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Employerinfo
type Employerinfo struct { 
	// OfficialName
	OfficialName *string `json:"officialName,omitempty"`


	// EmployeeId
	EmployeeId *string `json:"employeeId,omitempty"`


	// EmployeeType
	EmployeeType *string `json:"employeeType,omitempty"`


	// DateHire
	DateHire *string `json:"dateHire,omitempty"`

}

func (o *Employerinfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Employerinfo
	
	return json.Marshal(&struct { 
		OfficialName *string `json:"officialName,omitempty"`
		
		EmployeeId *string `json:"employeeId,omitempty"`
		
		EmployeeType *string `json:"employeeType,omitempty"`
		
		DateHire *string `json:"dateHire,omitempty"`
		*Alias
	}{ 
		OfficialName: o.OfficialName,
		
		EmployeeId: o.EmployeeId,
		
		EmployeeType: o.EmployeeType,
		
		DateHire: o.DateHire,
		Alias:    (*Alias)(o),
	})
}

func (o *Employerinfo) UnmarshalJSON(b []byte) error {
	var EmployerinfoMap map[string]interface{}
	err := json.Unmarshal(b, &EmployerinfoMap)
	if err != nil {
		return err
	}
	
	if OfficialName, ok := EmployerinfoMap["officialName"].(string); ok {
		o.OfficialName = &OfficialName
	}
	
	if EmployeeId, ok := EmployerinfoMap["employeeId"].(string); ok {
		o.EmployeeId = &EmployeeId
	}
	
	if EmployeeType, ok := EmployerinfoMap["employeeType"].(string); ok {
		o.EmployeeType = &EmployeeType
	}
	
	if DateHire, ok := EmployerinfoMap["dateHire"].(string); ok {
		o.DateHire = &DateHire
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Employerinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
