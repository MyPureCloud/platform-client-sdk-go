package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Headcountforecast - Headcount interval information for schedule
type Headcountforecast struct { 
	// Required - Headcount information with shrinkage
	Required *[]Headcountinterval `json:"required,omitempty"`


	// RequiredWithoutShrinkage - Headcount information without shrinkage
	RequiredWithoutShrinkage *[]Headcountinterval `json:"requiredWithoutShrinkage,omitempty"`

}

func (u *Headcountforecast) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Headcountforecast

	

	return json.Marshal(&struct { 
		Required *[]Headcountinterval `json:"required,omitempty"`
		
		RequiredWithoutShrinkage *[]Headcountinterval `json:"requiredWithoutShrinkage,omitempty"`
		*Alias
	}{ 
		Required: u.Required,
		
		RequiredWithoutShrinkage: u.RequiredWithoutShrinkage,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Headcountforecast) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
