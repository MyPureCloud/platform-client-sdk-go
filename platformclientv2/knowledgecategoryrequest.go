package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgecategoryrequest
type Knowledgecategoryrequest struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - Category name
	Name *string `json:"name,omitempty"`


	// Description - Category description
	Description *string `json:"description,omitempty"`


	// Parent - Category parent
	Parent *Documentcategoryinput `json:"parent,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Knowledgecategoryrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgecategoryrequest
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Parent *Documentcategoryinput `json:"parent,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		Parent: o.Parent,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgecategoryrequest) UnmarshalJSON(b []byte) error {
	var KnowledgecategoryrequestMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgecategoryrequestMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KnowledgecategoryrequestMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := KnowledgecategoryrequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := KnowledgecategoryrequestMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Parent, ok := KnowledgecategoryrequestMap["parent"].(map[string]interface{}); ok {
		ParentString, _ := json.Marshal(Parent)
		json.Unmarshal(ParentString, &o.Parent)
	}
	
	if SelfUri, ok := KnowledgecategoryrequestMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgecategoryrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
