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

func (o *Groupcontact) MarshalJSON() ([]byte, error) {
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
		Address: o.Address,
		
		Extension: o.Extension,
		
		Display: o.Display,
		
		VarType: o.VarType,
		
		MediaType: o.MediaType,
		Alias:    (*Alias)(o),
	})
}

func (o *Groupcontact) UnmarshalJSON(b []byte) error {
	var GroupcontactMap map[string]interface{}
	err := json.Unmarshal(b, &GroupcontactMap)
	if err != nil {
		return err
	}
	
	if Address, ok := GroupcontactMap["address"].(string); ok {
		o.Address = &Address
	}
    
	if Extension, ok := GroupcontactMap["extension"].(string); ok {
		o.Extension = &Extension
	}
    
	if Display, ok := GroupcontactMap["display"].(string); ok {
		o.Display = &Display
	}
    
	if VarType, ok := GroupcontactMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if MediaType, ok := GroupcontactMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Groupcontact) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
