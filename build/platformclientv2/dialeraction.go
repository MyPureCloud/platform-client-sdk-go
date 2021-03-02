package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Dialeraction) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
