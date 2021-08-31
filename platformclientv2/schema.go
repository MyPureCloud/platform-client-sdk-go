package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Schema
type Schema struct { 
	// Title - A core type's title
	Title *string `json:"title,omitempty"`


	// Description - A core type's description
	Description *string `json:"description,omitempty"`


	// VarType - An array of fundamental JSON Schema primitive types on which the core type is based
	VarType *[]string `json:"type,omitempty"`


	// Items - Denotes the type and pattern of the items in an enum core type
	Items *Items `json:"items,omitempty"`


	// Pattern - For the \"date\" and \"datetime\" core types, denotes the regex prescribing the allowable date/datetime format
	Pattern *string `json:"pattern,omitempty"`

}

func (o *Schema) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Schema
	
	return json.Marshal(&struct { 
		Title *string `json:"title,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		VarType *[]string `json:"type,omitempty"`
		
		Items *Items `json:"items,omitempty"`
		
		Pattern *string `json:"pattern,omitempty"`
		*Alias
	}{ 
		Title: o.Title,
		
		Description: o.Description,
		
		VarType: o.VarType,
		
		Items: o.Items,
		
		Pattern: o.Pattern,
		Alias:    (*Alias)(o),
	})
}

func (o *Schema) UnmarshalJSON(b []byte) error {
	var SchemaMap map[string]interface{}
	err := json.Unmarshal(b, &SchemaMap)
	if err != nil {
		return err
	}
	
	if Title, ok := SchemaMap["title"].(string); ok {
		o.Title = &Title
	}
	
	if Description, ok := SchemaMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if VarType, ok := SchemaMap["type"].([]interface{}); ok {
		VarTypeString, _ := json.Marshal(VarType)
		json.Unmarshal(VarTypeString, &o.VarType)
	}
	
	if Items, ok := SchemaMap["items"].(map[string]interface{}); ok {
		ItemsString, _ := json.Marshal(Items)
		json.Unmarshal(ItemsString, &o.Items)
	}
	
	if Pattern, ok := SchemaMap["pattern"].(string); ok {
		o.Pattern = &Pattern
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Schema) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
