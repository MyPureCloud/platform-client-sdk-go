package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Patchuser
type Patchuser struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// AcdAutoAnswer - The value that denotes if acdAutoAnswer is set on the user
	AcdAutoAnswer *bool `json:"acdAutoAnswer,omitempty"`


	// Certifications
	Certifications *[]string `json:"certifications,omitempty"`


	// Biography
	Biography *Biography `json:"biography,omitempty"`


	// EmployerInfo
	EmployerInfo *Employerinfo `json:"employerInfo,omitempty"`

}

// String returns a JSON representation of the model
func (o *Patchuser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
