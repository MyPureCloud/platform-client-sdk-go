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

func (o *Headcountforecast) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Headcountforecast
	
	return json.Marshal(&struct { 
		Required *[]Headcountinterval `json:"required,omitempty"`
		
		RequiredWithoutShrinkage *[]Headcountinterval `json:"requiredWithoutShrinkage,omitempty"`
		*Alias
	}{ 
		Required: o.Required,
		
		RequiredWithoutShrinkage: o.RequiredWithoutShrinkage,
		Alias:    (*Alias)(o),
	})
}

func (o *Headcountforecast) UnmarshalJSON(b []byte) error {
	var HeadcountforecastMap map[string]interface{}
	err := json.Unmarshal(b, &HeadcountforecastMap)
	if err != nil {
		return err
	}
	
	if Required, ok := HeadcountforecastMap["required"].([]interface{}); ok {
		RequiredString, _ := json.Marshal(Required)
		json.Unmarshal(RequiredString, &o.Required)
	}
	
	if RequiredWithoutShrinkage, ok := HeadcountforecastMap["requiredWithoutShrinkage"].([]interface{}); ok {
		RequiredWithoutShrinkageString, _ := json.Marshal(RequiredWithoutShrinkage)
		json.Unmarshal(RequiredWithoutShrinkageString, &o.RequiredWithoutShrinkage)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Headcountforecast) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
