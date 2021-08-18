package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Entitytypecriteria
type Entitytypecriteria struct { 
	// Key - The criteria key.
	Key *string `json:"key,omitempty"`


	// Values - The criteria values.
	Values *[]string `json:"values,omitempty"`


	// ShouldIgnoreCase - Should criteria be case insensitive.
	ShouldIgnoreCase *bool `json:"shouldIgnoreCase,omitempty"`


	// Operator - The comparison operator.
	Operator *string `json:"operator,omitempty"`


	// EntityType - The entity to match the pattern against.
	EntityType *string `json:"entityType,omitempty"`

}

func (u *Entitytypecriteria) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Entitytypecriteria

	

	return json.Marshal(&struct { 
		Key *string `json:"key,omitempty"`
		
		Values *[]string `json:"values,omitempty"`
		
		ShouldIgnoreCase *bool `json:"shouldIgnoreCase,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		EntityType *string `json:"entityType,omitempty"`
		*Alias
	}{ 
		Key: u.Key,
		
		Values: u.Values,
		
		ShouldIgnoreCase: u.ShouldIgnoreCase,
		
		Operator: u.Operator,
		
		EntityType: u.EntityType,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Entitytypecriteria) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
