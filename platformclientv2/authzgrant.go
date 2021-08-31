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

func (o *Authzgrant) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Authzgrant
	
	GrantMadeAt := new(string)
	if o.GrantMadeAt != nil {
		
		*GrantMadeAt = timeutil.Strftime(o.GrantMadeAt, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		SubjectId: o.SubjectId,
		
		Division: o.Division,
		
		Role: o.Role,
		
		GrantMadeAt: GrantMadeAt,
		Alias:    (*Alias)(o),
	})
}

func (o *Authzgrant) UnmarshalJSON(b []byte) error {
	var AuthzgrantMap map[string]interface{}
	err := json.Unmarshal(b, &AuthzgrantMap)
	if err != nil {
		return err
	}
	
	if SubjectId, ok := AuthzgrantMap["subjectId"].(string); ok {
		o.SubjectId = &SubjectId
	}
	
	if Division, ok := AuthzgrantMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Role, ok := AuthzgrantMap["role"].(map[string]interface{}); ok {
		RoleString, _ := json.Marshal(Role)
		json.Unmarshal(RoleString, &o.Role)
	}
	
	if grantMadeAtString, ok := AuthzgrantMap["grantMadeAt"].(string); ok {
		GrantMadeAt, _ := time.Parse("2006-01-02T15:04:05.999999Z", grantMadeAtString)
		o.GrantMadeAt = &GrantMadeAt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Authzgrant) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
