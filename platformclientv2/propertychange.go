package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Propertychange
type Propertychange struct { 
	// Property - The property that was changed
	Property *string `json:"property,omitempty"`


	// OldValues - Previous values for the property.
	OldValues *[]string `json:"oldValues,omitempty"`


	// NewValues - New values for the property.
	NewValues *[]string `json:"newValues,omitempty"`

}

func (o *Propertychange) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Propertychange
	
	return json.Marshal(&struct { 
		Property *string `json:"property,omitempty"`
		
		OldValues *[]string `json:"oldValues,omitempty"`
		
		NewValues *[]string `json:"newValues,omitempty"`
		*Alias
	}{ 
		Property: o.Property,
		
		OldValues: o.OldValues,
		
		NewValues: o.NewValues,
		Alias:    (*Alias)(o),
	})
}

func (o *Propertychange) UnmarshalJSON(b []byte) error {
	var PropertychangeMap map[string]interface{}
	err := json.Unmarshal(b, &PropertychangeMap)
	if err != nil {
		return err
	}
	
	if Property, ok := PropertychangeMap["property"].(string); ok {
		o.Property = &Property
	}
	
	if OldValues, ok := PropertychangeMap["oldValues"].([]interface{}); ok {
		OldValuesString, _ := json.Marshal(OldValues)
		json.Unmarshal(OldValuesString, &o.OldValues)
	}
	
	if NewValues, ok := PropertychangeMap["newValues"].([]interface{}); ok {
		NewValuesString, _ := json.Marshal(NewValues)
		json.Unmarshal(NewValuesString, &o.NewValues)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Propertychange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
