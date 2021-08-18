package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Lexslot
type Lexslot struct { 
	// Name - The slot name
	Name *string `json:"name,omitempty"`


	// Description - The slot description
	Description *string `json:"description,omitempty"`


	// VarType - The slot type
	VarType *string `json:"type,omitempty"`


	// Priority - The priority of the slot
	Priority *int `json:"priority,omitempty"`

}

func (u *Lexslot) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Lexslot

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		Description: u.Description,
		
		VarType: u.VarType,
		
		Priority: u.Priority,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Lexslot) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
