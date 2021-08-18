package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningmodulesummary - Learning module summary data
type Learningmodulesummary struct { 
	// AssignedCount - The total number of assignments assigned for a learning module
	AssignedCount *int `json:"assignedCount,omitempty"`


	// CompletedCount - The number of assignments completed for a learning module
	CompletedCount *int `json:"completedCount,omitempty"`


	// PassedCount - The number of assignments passed for a learning module
	PassedCount *int `json:"passedCount,omitempty"`


	// CompletedSum - The sum of assignment scores for a learning module
	CompletedSum *float32 `json:"completedSum,omitempty"`

}

func (u *Learningmodulesummary) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningmodulesummary

	

	return json.Marshal(&struct { 
		AssignedCount *int `json:"assignedCount,omitempty"`
		
		CompletedCount *int `json:"completedCount,omitempty"`
		
		PassedCount *int `json:"passedCount,omitempty"`
		
		CompletedSum *float32 `json:"completedSum,omitempty"`
		*Alias
	}{ 
		AssignedCount: u.AssignedCount,
		
		CompletedCount: u.CompletedCount,
		
		PassedCount: u.PassedCount,
		
		CompletedSum: u.CompletedSum,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Learningmodulesummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
