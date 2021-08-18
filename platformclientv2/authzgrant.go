package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Authzgrant
type Authzgrant struct { 
	// SubjectId
	SubjectId *string `json:"subjectId,omitempty"`


	// Division
	Division *Authzdivision `json:"division,omitempty"`


	// Role
	Role *Authzgrantrole `json:"role,omitempty"`


	// GrantMadeAt - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	GrantMadeAt *time.Time `json:"grantMadeAt,omitempty"`

}

func (u *Authzgrant) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Authzgrant

	
	GrantMadeAt := new(string)
	if u.GrantMadeAt != nil {
		
		*GrantMadeAt = timeutil.Strftime(u.GrantMadeAt, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		GrantMadeAt = nil
	}
	

	return json.Marshal(&struct { 
		SubjectId *string `json:"subjectId,omitempty"`
		
		Division *Authzdivision `json:"division,omitempty"`
		
		Role *Authzgrantrole `json:"role,omitempty"`
		
		GrantMadeAt *string `json:"grantMadeAt,omitempty"`
		*Alias
	}{ 
		SubjectId: u.SubjectId,
		
		Division: u.Division,
		
		Role: u.Role,
		
		GrantMadeAt: GrantMadeAt,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Authzgrant) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
