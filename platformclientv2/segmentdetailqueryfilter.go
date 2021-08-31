package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Segmentdetailqueryfilter
type Segmentdetailqueryfilter struct { 
	// VarType - Boolean operation to apply to the provided predicates and clauses
	VarType *string `json:"type,omitempty"`


	// Clauses - Boolean 'and/or' logic with up to two-levels of nesting
	Clauses *[]Segmentdetailqueryclause `json:"clauses,omitempty"`


	// Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
	Predicates *[]Segmentdetailquerypredicate `json:"predicates,omitempty"`

}

func (o *Segmentdetailqueryfilter) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Segmentdetailqueryfilter
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Clauses *[]Segmentdetailqueryclause `json:"clauses,omitempty"`
		
		Predicates *[]Segmentdetailquerypredicate `json:"predicates,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Clauses: o.Clauses,
		
		Predicates: o.Predicates,
		Alias:    (*Alias)(o),
	})
}

func (o *Segmentdetailqueryfilter) UnmarshalJSON(b []byte) error {
	var SegmentdetailqueryfilterMap map[string]interface{}
	err := json.Unmarshal(b, &SegmentdetailqueryfilterMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := SegmentdetailqueryfilterMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Clauses, ok := SegmentdetailqueryfilterMap["clauses"].([]interface{}); ok {
		ClausesString, _ := json.Marshal(Clauses)
		json.Unmarshal(ClausesString, &o.Clauses)
	}
	
	if Predicates, ok := SegmentdetailqueryfilterMap["predicates"].([]interface{}); ok {
		PredicatesString, _ := json.Marshal(Predicates)
		json.Unmarshal(PredicatesString, &o.Predicates)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Segmentdetailqueryfilter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
