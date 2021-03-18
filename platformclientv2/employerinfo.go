package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Employerinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
