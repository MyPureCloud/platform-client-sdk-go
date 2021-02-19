package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Requestmapping) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
