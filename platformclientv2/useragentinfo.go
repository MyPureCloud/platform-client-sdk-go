package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Useragentinfo
type Useragentinfo struct { 
	// FirmwareVersion - The firmware version of the phone.
	FirmwareVersion *string `json:"firmwareVersion,omitempty"`


	// Manufacturer - The manufacturer of the phone.
	Manufacturer *string `json:"manufacturer,omitempty"`


	// Model - The model of the phone.
	Model *string `json:"model,omitempty"`

}

func (u *Useragentinfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Useragentinfo

	

	return json.Marshal(&struct { 
		FirmwareVersion *string `json:"firmwareVersion,omitempty"`
		
		Manufacturer *string `json:"manufacturer,omitempty"`
		
		Model *string `json:"model,omitempty"`
		*Alias
	}{ 
		FirmwareVersion: u.FirmwareVersion,
		
		Manufacturer: u.Manufacturer,
		
		Model: u.Model,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Useragentinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
