package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Documentvariationcontext
type Documentvariationcontext struct { 
	// Context - The knowledge context associated with the variation.
	Context *Knowledgecontextreference `json:"context,omitempty"`


	// Values - The list of knowledge context values associated with the variation.
	Values *[]Knowledgecontextvaluereference `json:"values,omitempty"`

}

func (o *Documentvariationcontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Documentvariationcontext
	
	return json.Marshal(&struct { 
		Context *Knowledgecontextreference `json:"context,omitempty"`
		
		Values *[]Knowledgecontextvaluereference `json:"values,omitempty"`
		*Alias
	}{ 
		Context: o.Context,
		
		Values: o.Values,
		Alias:    (*Alias)(o),
	})
}

func (o *Documentvariationcontext) UnmarshalJSON(b []byte) error {
	var DocumentvariationcontextMap map[string]interface{}
	err := json.Unmarshal(b, &DocumentvariationcontextMap)
	if err != nil {
		return err
	}
	
	if Context, ok := DocumentvariationcontextMap["context"].(map[string]interface{}); ok {
		ContextString, _ := json.Marshal(Context)
		json.Unmarshal(ContextString, &o.Context)
	}
	
	if Values, ok := DocumentvariationcontextMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Documentvariationcontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
