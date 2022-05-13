package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Associatedvaluefield
type Associatedvaluefield struct { 
	// DataType - The data type of the value field.
	DataType *string `json:"dataType,omitempty"`


	// Name - The field name for extracting value from event.
	Name *string `json:"name,omitempty"`

}

func (o *Associatedvaluefield) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Associatedvaluefield
	
	return json.Marshal(&struct { 
		DataType *string `json:"dataType,omitempty"`
		
		Name *string `json:"name,omitempty"`
		*Alias
	}{ 
		DataType: o.DataType,
		
		Name: o.Name,
		Alias:    (*Alias)(o),
	})
}

func (o *Associatedvaluefield) UnmarshalJSON(b []byte) error {
	var AssociatedvaluefieldMap map[string]interface{}
	err := json.Unmarshal(b, &AssociatedvaluefieldMap)
	if err != nil {
		return err
	}
	
	if DataType, ok := AssociatedvaluefieldMap["dataType"].(string); ok {
		o.DataType = &DataType
	}
    
	if Name, ok := AssociatedvaluefieldMap["name"].(string); ok {
		o.Name = &Name
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Associatedvaluefield) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
