package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgeguestsessioncontext
type Knowledgeguestsessioncontext struct { 
	// Id - The context id associated with the session.
	Id *string `json:"id,omitempty"`


	// Values - The list of knowledge context values associated with the session.
	Values *[]Entity `json:"values,omitempty"`

}

func (o *Knowledgeguestsessioncontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgeguestsessioncontext
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Values *[]Entity `json:"values,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Values: o.Values,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgeguestsessioncontext) UnmarshalJSON(b []byte) error {
	var KnowledgeguestsessioncontextMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgeguestsessioncontextMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KnowledgeguestsessioncontextMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Values, ok := KnowledgeguestsessioncontextMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgeguestsessioncontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
