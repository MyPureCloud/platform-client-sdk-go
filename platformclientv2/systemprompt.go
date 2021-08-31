package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Systemprompt
type Systemprompt struct { 
	// Id - The system prompt identifier
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// Resources
	Resources *[]Systempromptasset `json:"resources,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Systemprompt) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Systemprompt
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Resources *[]Systempromptasset `json:"resources,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		Resources: o.Resources,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Systemprompt) UnmarshalJSON(b []byte) error {
	var SystempromptMap map[string]interface{}
	err := json.Unmarshal(b, &SystempromptMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SystempromptMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := SystempromptMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Description, ok := SystempromptMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if Resources, ok := SystempromptMap["resources"].([]interface{}); ok {
		ResourcesString, _ := json.Marshal(Resources)
		json.Unmarshal(ResourcesString, &o.Resources)
	}
	
	if SelfUri, ok := SystempromptMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Systemprompt) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
