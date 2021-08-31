package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Surveydetailqueryclause
type Surveydetailqueryclause struct { 
	// VarType - Boolean operation to apply to the provided predicates
	VarType *string `json:"type,omitempty"`


	// Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
	Predicates *[]Surveydetailquerypredicate `json:"predicates,omitempty"`

}

func (o *Surveydetailqueryclause) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Surveydetailqueryclause
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Predicates *[]Surveydetailquerypredicate `json:"predicates,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Predicates: o.Predicates,
		Alias:    (*Alias)(o),
	})
}

func (o *Surveydetailqueryclause) UnmarshalJSON(b []byte) error {
	var SurveydetailqueryclauseMap map[string]interface{}
	err := json.Unmarshal(b, &SurveydetailqueryclauseMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := SurveydetailqueryclauseMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Predicates, ok := SurveydetailqueryclauseMap["predicates"].([]interface{}); ok {
		PredicatesString, _ := json.Marshal(Predicates)
		json.Unmarshal(PredicatesString, &o.Predicates)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Surveydetailqueryclause) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
