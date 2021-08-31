package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialeraction
type Dialeraction struct { 
	// VarType - The type of this DialerAction.
	VarType *string `json:"type,omitempty"`


	// ActionTypeName - Additional type specification for this DialerAction.
	ActionTypeName *string `json:"actionTypeName,omitempty"`


	// UpdateOption - Specifies how a contact attribute should be updated. Required for MODIFY_CONTACT_ATTRIBUTE.
	UpdateOption *string `json:"updateOption,omitempty"`


	// Properties - A map of key-value pairs pertinent to the DialerAction. Different types of DialerActions require different properties. MODIFY_CONTACT_ATTRIBUTE with an updateOption of SET takes a contact column as the key and accepts any value. SCHEDULE_CALLBACK takes a key 'callbackOffset' that specifies how far in the future the callback should be scheduled, in minutes. SET_CALLER_ID takes two keys: 'callerAddress', which should be the caller id phone number, and 'callerName'. For either key, you can also specify a column on the contact to get the value from. To do this, specify 'contact.Column', where 'Column' is the name of the contact column from which to get the value. SET_SKILLS takes a key 'skills' with an array of skill ids wrapped into a string (Example: {'skills': '['skillIdHere']'} ).
	Properties *map[string]string `json:"properties,omitempty"`

}

func (o *Dialeraction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialeraction
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		ActionTypeName *string `json:"actionTypeName,omitempty"`
		
		UpdateOption *string `json:"updateOption,omitempty"`
		
		Properties *map[string]string `json:"properties,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		ActionTypeName: o.ActionTypeName,
		
		UpdateOption: o.UpdateOption,
		
		Properties: o.Properties,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialeraction) UnmarshalJSON(b []byte) error {
	var DialeractionMap map[string]interface{}
	err := json.Unmarshal(b, &DialeractionMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := DialeractionMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if ActionTypeName, ok := DialeractionMap["actionTypeName"].(string); ok {
		o.ActionTypeName = &ActionTypeName
	}
	
	if UpdateOption, ok := DialeractionMap["updateOption"].(string); ok {
		o.UpdateOption = &UpdateOption
	}
	
	if Properties, ok := DialeractionMap["properties"].(map[string]interface{}); ok {
		PropertiesString, _ := json.Marshal(Properties)
		json.Unmarshal(PropertiesString, &o.Properties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialeraction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
