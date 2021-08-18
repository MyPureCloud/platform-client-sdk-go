package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentcreate
type Learningassignmentcreate struct { 
	// ModuleId - The Learning module Id associated with this assignment
	ModuleId *string `json:"moduleId,omitempty"`


	// UserId - The User for whom the assignment is assigned
	UserId *string `json:"userId,omitempty"`


	// RecommendedCompletionDate - The recommended completion date of assignment. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	RecommendedCompletionDate *time.Time `json:"recommendedCompletionDate,omitempty"`

}

func (u *Learningassignmentcreate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassignmentcreate

	
	RecommendedCompletionDate := new(string)
	if u.RecommendedCompletionDate != nil {
		
		*RecommendedCompletionDate = timeutil.Strftime(u.RecommendedCompletionDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		RecommendedCompletionDate = nil
	}
	

	return json.Marshal(&struct { 
		ModuleId *string `json:"moduleId,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		RecommendedCompletionDate *string `json:"recommendedCompletionDate,omitempty"`
		*Alias
	}{ 
		ModuleId: u.ModuleId,
		
		UserId: u.UserId,
		
		RecommendedCompletionDate: RecommendedCompletionDate,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Learningassignmentcreate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
