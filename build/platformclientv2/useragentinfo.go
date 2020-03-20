package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Useragentinfo) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
