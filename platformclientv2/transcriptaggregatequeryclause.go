package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Transcriptaggregatequeryclause
type Transcriptaggregatequeryclause struct { 
	// VarType - Boolean operation to apply to the provided predicates
	VarType *string `json:"type,omitempty"`


	// Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
	Predicates *[]Transcriptaggregatequerypredicate `json:"predicates,omitempty"`

}

func (o *Transcriptaggregatequeryclause) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Transcriptaggregatequeryclause
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Predicates *[]Transcriptaggregatequerypredicate `json:"predicates,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Predicates: o.Predicates,
		Alias:    (*Alias)(o),
	})
}

func (o *Transcriptaggregatequeryclause) UnmarshalJSON(b []byte) error {
	var TranscriptaggregatequeryclauseMap map[string]interface{}
	err := json.Unmarshal(b, &TranscriptaggregatequeryclauseMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := TranscriptaggregatequeryclauseMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Predicates, ok := TranscriptaggregatequeryclauseMap["predicates"].([]interface{}); ok {
		PredicatesString, _ := json.Marshal(Predicates)
		json.Unmarshal(PredicatesString, &o.Predicates)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Transcriptaggregatequeryclause) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
