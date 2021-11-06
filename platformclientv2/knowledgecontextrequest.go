package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgecontextrequest
type Knowledgecontextrequest struct { 
	// Name - Context name.
	Name *string `json:"name,omitempty"`


	// Description - Context description.
	Description *string `json:"description,omitempty"`

}

func (o *Knowledgecontextrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgecontextrequest
	
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

func (o *Knowledgecontextrequest) UnmarshalJSON(b []byte) error {
	var KnowledgecontextrequestMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgecontextrequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := KnowledgecontextrequestMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Description, ok := KnowledgecontextrequestMap["description"].(string); ok {
		o.Description = &Description
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgecontextrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
