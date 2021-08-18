package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Lexintent
type Lexintent struct { 
	// Name - The intent name
	Name *string `json:"name,omitempty"`


	// Description - A description of the intent
	Description *string `json:"description,omitempty"`


	// Slots - An object mapping slot names to Slot objects
	Slots *map[string]Lexslot `json:"slots,omitempty"`


	// Version - The intent version
	Version *string `json:"version,omitempty"`

}

func (u *Lexintent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Lexintent

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Slots *map[string]Lexslot `json:"slots,omitempty"`
		
		Version *string `json:"version,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		Description: u.Description,
		
		Slots: u.Slots,
		
		Version: u.Version,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Lexintent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
