package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Surveyaggregatequeryfilter
type Surveyaggregatequeryfilter struct { 
	// VarType - Boolean operation to apply to the provided predicates and clauses
	VarType *string `json:"type,omitempty"`


	// Clauses - Boolean 'and/or' logic with up to two-levels of nesting
	Clauses *[]Surveyaggregatequeryclause `json:"clauses,omitempty"`


	// Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
	Predicates *[]Surveyaggregatequerypredicate `json:"predicates,omitempty"`

}

func (o *Surveyaggregatequeryfilter) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Surveyaggregatequeryfilter
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Clauses *[]Surveyaggregatequeryclause `json:"clauses,omitempty"`
		
		Predicates *[]Surveyaggregatequerypredicate `json:"predicates,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Clauses: o.Clauses,
		
		Predicates: o.Predicates,
		Alias:    (*Alias)(o),
	})
}

func (o *Surveyaggregatequeryfilter) UnmarshalJSON(b []byte) error {
	var SurveyaggregatequeryfilterMap map[string]interface{}
	err := json.Unmarshal(b, &SurveyaggregatequeryfilterMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := SurveyaggregatequeryfilterMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Clauses, ok := SurveyaggregatequeryfilterMap["clauses"].([]interface{}); ok {
		ClausesString, _ := json.Marshal(Clauses)
		json.Unmarshal(ClausesString, &o.Clauses)
	}
	
	if Predicates, ok := SurveyaggregatequeryfilterMap["predicates"].([]interface{}); ok {
		PredicatesString, _ := json.Marshal(Predicates)
		json.Unmarshal(PredicatesString, &o.Predicates)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Surveyaggregatequeryfilter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
