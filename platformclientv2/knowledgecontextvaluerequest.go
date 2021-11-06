package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgecontextvaluerequest
type Knowledgecontextvaluerequest struct { 
	// Name - Context value name.
	Name *string `json:"name,omitempty"`


	// Description - Context value description.
	Description *string `json:"description,omitempty"`

}

func (o *Knowledgecontextvaluerequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgecontextvaluerequest
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Description: o.Description,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgecontextvaluerequest) UnmarshalJSON(b []byte) error {
	var KnowledgecontextvaluerequestMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgecontextvaluerequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := KnowledgecontextvaluerequestMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Description, ok := KnowledgecontextvaluerequestMap["description"].(string); ok {
		o.Description = &Description
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgecontextvaluerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
