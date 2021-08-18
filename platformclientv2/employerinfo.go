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

func (u *Employerinfo) MarshalJSON() ([]byte, error) {
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
		OfficialName: u.OfficialName,
		
		EmployeeId: u.EmployeeId,
		
		EmployeeType: u.EmployeeType,
		
		DateHire: u.DateHire,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Employerinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
