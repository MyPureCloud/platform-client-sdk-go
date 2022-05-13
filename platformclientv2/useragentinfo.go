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

func (o *Useragentinfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Useragentinfo
	
	return json.Marshal(&struct { 
		FirmwareVersion *string `json:"firmwareVersion,omitempty"`
		
		Manufacturer *string `json:"manufacturer,omitempty"`
		
		Model *string `json:"model,omitempty"`
		*Alias
	}{ 
		FirmwareVersion: o.FirmwareVersion,
		
		Manufacturer: o.Manufacturer,
		
		Model: o.Model,
		Alias:    (*Alias)(o),
	})
}

func (o *Useragentinfo) UnmarshalJSON(b []byte) error {
	var UseragentinfoMap map[string]interface{}
	err := json.Unmarshal(b, &UseragentinfoMap)
	if err != nil {
		return err
	}
	
	if FirmwareVersion, ok := UseragentinfoMap["firmwareVersion"].(string); ok {
		o.FirmwareVersion = &FirmwareVersion
	}
    
	if Manufacturer, ok := UseragentinfoMap["manufacturer"].(string); ok {
		o.Manufacturer = &Manufacturer
	}
    
	if Model, ok := UseragentinfoMap["model"].(string); ok {
		o.Model = &Model
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Useragentinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
