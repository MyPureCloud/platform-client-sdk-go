package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Contactsort) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contactsort

	

	return json.Marshal(&struct { 
		FieldName *string `json:"fieldName,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		Numeric *bool `json:"numeric,omitempty"`
		*Alias
	}{ 
		FieldName: u.FieldName,
		
		Direction: u.Direction,
		
		Numeric: u.Numeric,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Contactsort) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
