package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Whatsappavailablephonenumberdetails
type Whatsappavailablephonenumberdetails struct { 
	// Name - The verified name associated with this phone number.
	Name *string `json:"name,omitempty"`


	// PhoneNumber - The display name associated with this phone number. It's typically the E.164 representation of the number.
	PhoneNumber *string `json:"phoneNumber,omitempty"`


	// Status - The status of this phone number.
	Status *string `json:"status,omitempty"`

}

func (o *Whatsappavailablephonenumberdetails) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Whatsappavailablephonenumberdetails
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		PhoneNumber *string `json:"phoneNumber,omitempty"`
		
		Status *string `json:"status,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		PhoneNumber: o.PhoneNumber,
		
		Status: o.Status,
		Alias:    (*Alias)(o),
	})
}

func (o *Whatsappavailablephonenumberdetails) UnmarshalJSON(b []byte) error {
	var WhatsappavailablephonenumberdetailsMap map[string]interface{}
	err := json.Unmarshal(b, &WhatsappavailablephonenumberdetailsMap)
	if err != nil {
		return err
	}
	
	if Name, ok := WhatsappavailablephonenumberdetailsMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if PhoneNumber, ok := WhatsappavailablephonenumberdetailsMap["phoneNumber"].(string); ok {
		o.PhoneNumber = &PhoneNumber
	}
    
	if Status, ok := WhatsappavailablephonenumberdetailsMap["status"].(string); ok {
		o.Status = &Status
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Whatsappavailablephonenumberdetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
