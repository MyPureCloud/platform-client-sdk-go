package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Requestmapping) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
