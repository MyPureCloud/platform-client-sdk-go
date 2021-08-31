package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Requestmapping
type Requestmapping struct { 
	// Name - Name of the Integration Action Attribute to supply the value for
	Name *string `json:"name,omitempty"`


	// AttributeType - Type of the value supplied
	AttributeType *string `json:"attributeType,omitempty"`


	// MappingType - Method of finding value to use with Attribute
	MappingType *string `json:"mappingType,omitempty"`


	// Value - Value to supply for the specified Attribute
	Value *string `json:"value,omitempty"`

}

func (o *Requestmapping) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Requestmapping
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		AttributeType *string `json:"attributeType,omitempty"`
		
		MappingType *string `json:"mappingType,omitempty"`
		
		Value *string `json:"value,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		AttributeType: o.AttributeType,
		
		MappingType: o.MappingType,
		
		Value: o.Value,
		Alias:    (*Alias)(o),
	})
}

func (o *Requestmapping) UnmarshalJSON(b []byte) error {
	var RequestmappingMap map[string]interface{}
	err := json.Unmarshal(b, &RequestmappingMap)
	if err != nil {
		return err
	}
	
	if Name, ok := RequestmappingMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if AttributeType, ok := RequestmappingMap["attributeType"].(string); ok {
		o.AttributeType = &AttributeType
	}
	
	if MappingType, ok := RequestmappingMap["mappingType"].(string); ok {
		o.MappingType = &MappingType
	}
	
	if Value, ok := RequestmappingMap["value"].(string); ok {
		o.Value = &Value
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Requestmapping) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
