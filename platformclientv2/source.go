package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Source
type Source struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the source
	Name *string `json:"name,omitempty"`


	// Description - The description of the source
	Description *string `json:"description,omitempty"`


	// VarType - The type of source
	VarType *string `json:"type,omitempty"`


	// Deactivated
	Deactivated *bool `json:"deactivated,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Source) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Source
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Deactivated *bool `json:"deactivated,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		VarType: o.VarType,
		
		Deactivated: o.Deactivated,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Source) UnmarshalJSON(b []byte) error {
	var SourceMap map[string]interface{}
	err := json.Unmarshal(b, &SourceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SourceMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := SourceMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := SourceMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if VarType, ok := SourceMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Deactivated, ok := SourceMap["deactivated"].(bool); ok {
		o.Deactivated = &Deactivated
	}
    
	if SelfUri, ok := SourceMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Source) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
