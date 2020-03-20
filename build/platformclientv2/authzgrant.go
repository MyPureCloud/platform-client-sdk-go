package platformclientv2
import (
	"time"
	"encoding/json"
)

// Authzgrant
type Authzgrant struct { 
	// SubjectId
	SubjectId *string `json:"subjectId,omitempty"`


	// Division
	Division *Authzdivision `json:"division,omitempty"`


	// Role
	Role *Authzgrantrole `json:"role,omitempty"`


	// GrantMadeAt - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	GrantMadeAt *time.Time `json:"grantMadeAt,omitempty"`

}

// String returns a JSON representation of the model
func (o *Authzgrant) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
