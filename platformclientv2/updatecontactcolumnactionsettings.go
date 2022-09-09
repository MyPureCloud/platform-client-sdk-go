package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Updatecontactcolumnactionsettings
type Updatecontactcolumnactionsettings struct { 
	// Properties - A mapping of contact columns to their new values.
	Properties *map[string]string `json:"properties,omitempty"`


	// UpdateOption - The type of update to make to the specified contact column(s).
	UpdateOption *string `json:"updateOption,omitempty"`

}

func (o *Updatecontactcolumnactionsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Updatecontactcolumnactionsettings
	
	return json.Marshal(&struct { 
		Properties *map[string]string `json:"properties,omitempty"`
		
		UpdateOption *string `json:"updateOption,omitempty"`
		*Alias
	}{ 
		Properties: o.Properties,
		
		UpdateOption: o.UpdateOption,
		Alias:    (*Alias)(o),
	})
}

func (o *Updatecontactcolumnactionsettings) UnmarshalJSON(b []byte) error {
	var UpdatecontactcolumnactionsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &UpdatecontactcolumnactionsettingsMap)
	if err != nil {
		return err
	}
	
	if Properties, ok := UpdatecontactcolumnactionsettingsMap["properties"].(map[string]interface{}); ok {
		PropertiesString, _ := json.Marshal(Properties)
		json.Unmarshal(PropertiesString, &o.Properties)
	}
	
	if UpdateOption, ok := UpdatecontactcolumnactionsettingsMap["updateOption"].(string); ok {
		o.UpdateOption = &UpdateOption
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Updatecontactcolumnactionsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
