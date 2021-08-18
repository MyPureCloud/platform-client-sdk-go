package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Disallowedentitylearningassignmentreference
type Disallowedentitylearningassignmentreference struct { 
	// ErrorCode - The error code associated with this disallowed entity
	ErrorCode *string `json:"errorCode,omitempty"`


	// Entity - The entity that was disallowed
	Entity *Learningassignmentreference `json:"entity,omitempty"`

}

func (u *Disallowedentitylearningassignmentreference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Disallowedentitylearningassignmentreference

	

	return json.Marshal(&struct { 
		ErrorCode *string `json:"errorCode,omitempty"`
		
		Entity *Learningassignmentreference `json:"entity,omitempty"`
		*Alias
	}{ 
		ErrorCode: u.ErrorCode,
		
		Entity: u.Entity,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Disallowedentitylearningassignmentreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
