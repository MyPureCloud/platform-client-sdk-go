package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Prompt
type Prompt struct { 
	// Id - The prompt identifier
	Id *string `json:"id,omitempty"`


	// Name - The prompt name.
	Name *string `json:"name,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// Resources - List of resources associated with this prompt
	Resources *[]Promptasset `json:"resources,omitempty"`


	// CurrentOperation - Current prompt operation status
	CurrentOperation *Operation `json:"currentOperation,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Prompt) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Prompt
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Resources *[]Promptasset `json:"resources,omitempty"`
		
		CurrentOperation *Operation `json:"currentOperation,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		Resources: o.Resources,
		
		CurrentOperation: o.CurrentOperation,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Prompt) UnmarshalJSON(b []byte) error {
	var PromptMap map[string]interface{}
	err := json.Unmarshal(b, &PromptMap)
	if err != nil {
		return err
	}
	
	if Id, ok := PromptMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := PromptMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Description, ok := PromptMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if Resources, ok := PromptMap["resources"].([]interface{}); ok {
		ResourcesString, _ := json.Marshal(Resources)
		json.Unmarshal(ResourcesString, &o.Resources)
	}
	
	if CurrentOperation, ok := PromptMap["currentOperation"].(map[string]interface{}); ok {
		CurrentOperationString, _ := json.Marshal(CurrentOperation)
		json.Unmarshal(CurrentOperationString, &o.CurrentOperation)
	}
	
	if SelfUri, ok := PromptMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Prompt) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
