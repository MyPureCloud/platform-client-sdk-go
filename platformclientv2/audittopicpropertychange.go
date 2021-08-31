package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Audittopicpropertychange
type Audittopicpropertychange struct { 
	// Property
	Property *string `json:"property,omitempty"`


	// OldValues
	OldValues *[]string `json:"oldValues,omitempty"`


	// NewValues
	NewValues *[]string `json:"newValues,omitempty"`

}

func (o *Audittopicpropertychange) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Audittopicpropertychange
	
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

func (o *Audittopicpropertychange) UnmarshalJSON(b []byte) error {
	var AudittopicpropertychangeMap map[string]interface{}
	err := json.Unmarshal(b, &AudittopicpropertychangeMap)
	if err != nil {
		return err
	}
	
	if Property, ok := AudittopicpropertychangeMap["property"].(string); ok {
		o.Property = &Property
	}
	
	if OldValues, ok := AudittopicpropertychangeMap["oldValues"].([]interface{}); ok {
		OldValuesString, _ := json.Marshal(OldValues)
		json.Unmarshal(OldValuesString, &o.OldValues)
	}
	
	if NewValues, ok := AudittopicpropertychangeMap["newValues"].([]interface{}); ok {
		NewValuesString, _ := json.Marshal(NewValues)
		json.Unmarshal(NewValuesString, &o.NewValues)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Audittopicpropertychange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
