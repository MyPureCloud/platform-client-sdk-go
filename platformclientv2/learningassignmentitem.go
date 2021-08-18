package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentitem
type Learningassignmentitem struct { 
	// ModuleId - The Learning Module ID associated with this assignment
	ModuleId *string `json:"moduleId,omitempty"`


	// UserId - The User ID associated with this assignment
	UserId *string `json:"userId,omitempty"`

}

func (u *Learningassignmentitem) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassignmentitem

	

	return json.Marshal(&struct { 
		ModuleId *string `json:"moduleId,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		*Alias
	}{ 
		ModuleId: u.ModuleId,
		
		UserId: u.UserId,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Learningassignmentitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
