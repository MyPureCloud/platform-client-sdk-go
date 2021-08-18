package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Groupcontact
type Groupcontact struct { 
	// Address - Phone number for this contact type
	Address *string `json:"address,omitempty"`


	// Extension - Extension is set if the number is e164 valid
	Extension *string `json:"extension,omitempty"`


	// Display - Formatted version of the address property
	Display *string `json:"display,omitempty"`


	// VarType - Contact type of the address
	VarType *string `json:"type,omitempty"`


	// MediaType - Media type of the address
	MediaType *string `json:"mediaType,omitempty"`

}

func (u *Groupcontact) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Groupcontact

	

	return json.Marshal(&struct { 
		Address *string `json:"address,omitempty"`
		
		Extension *string `json:"extension,omitempty"`
		
		Display *string `json:"display,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		*Alias
	}{ 
		Address: u.Address,
		
		Extension: u.Extension,
		
		Display: u.Display,
		
		VarType: u.VarType,
		
		MediaType: u.MediaType,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Groupcontact) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
