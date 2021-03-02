package platformclientv2
import (
	"encoding/json"
)

// Contactsort
type Contactsort struct { 
	// FieldName
	FieldName *string `json:"fieldName,omitempty"`


	// Direction - The direction in which to sort contacts.
	Direction *string `json:"direction,omitempty"`


	// Numeric - Whether or not the column contains numeric data.
	Numeric *bool `json:"numeric,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contactsort) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
