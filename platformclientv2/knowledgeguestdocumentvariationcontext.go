package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgeguestdocumentvariationcontext
type Knowledgeguestdocumentvariationcontext struct { 
	// Context - The knowledge context associated with the variation.
	Context *Addressableentityref `json:"context,omitempty"`


	// Values - The list of knowledge context values associated with the variation.
	Values *[]Addressableentityref `json:"values,omitempty"`

}

func (o *Knowledgeguestdocumentvariationcontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgeguestdocumentvariationcontext
	
	return json.Marshal(&struct { 
		Context *Addressableentityref `json:"context,omitempty"`
		
		Values *[]Addressableentityref `json:"values,omitempty"`
		*Alias
	}{ 
		Context: o.Context,
		
		Values: o.Values,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgeguestdocumentvariationcontext) UnmarshalJSON(b []byte) error {
	var KnowledgeguestdocumentvariationcontextMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgeguestdocumentvariationcontextMap)
	if err != nil {
		return err
	}
	
	if Context, ok := KnowledgeguestdocumentvariationcontextMap["context"].(map[string]interface{}); ok {
		ContextString, _ := json.Marshal(Context)
		json.Unmarshal(ContextString, &o.Context)
	}
	
	if Values, ok := KnowledgeguestdocumentvariationcontextMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgeguestdocumentvariationcontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
