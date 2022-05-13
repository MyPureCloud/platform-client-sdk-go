package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Namedentitytypeitem
type Namedentitytypeitem struct { 
	// Value - A value for an named entity type definition.
	Value *string `json:"value,omitempty"`


	// Synonyms - Synonyms for the given named entity value.
	Synonyms *[]string `json:"synonyms,omitempty"`

}

func (o *Namedentitytypeitem) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Namedentitytypeitem
	
	return json.Marshal(&struct { 
		Value *string `json:"value,omitempty"`
		
		Synonyms *[]string `json:"synonyms,omitempty"`
		*Alias
	}{ 
		Value: o.Value,
		
		Synonyms: o.Synonyms,
		Alias:    (*Alias)(o),
	})
}

func (o *Namedentitytypeitem) UnmarshalJSON(b []byte) error {
	var NamedentitytypeitemMap map[string]interface{}
	err := json.Unmarshal(b, &NamedentitytypeitemMap)
	if err != nil {
		return err
	}
	
	if Value, ok := NamedentitytypeitemMap["value"].(string); ok {
		o.Value = &Value
	}
    
	if Synonyms, ok := NamedentitytypeitemMap["synonyms"].([]interface{}); ok {
		SynonymsString, _ := json.Marshal(Synonyms)
		json.Unmarshal(SynonymsString, &o.Synonyms)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Namedentitytypeitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
