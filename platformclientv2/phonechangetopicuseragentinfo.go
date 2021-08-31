package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Phonechangetopicuseragentinfo
type Phonechangetopicuseragentinfo struct { 
	// FirmwareVersion
	FirmwareVersion *string `json:"firmwareVersion,omitempty"`


	// Manufacturer
	Manufacturer *string `json:"manufacturer,omitempty"`


	// Model
	Model *string `json:"model,omitempty"`

}

func (o *Phonechangetopicuseragentinfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Phonechangetopicuseragentinfo
	
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

func (o *Phonechangetopicuseragentinfo) UnmarshalJSON(b []byte) error {
	var PhonechangetopicuseragentinfoMap map[string]interface{}
	err := json.Unmarshal(b, &PhonechangetopicuseragentinfoMap)
	if err != nil {
		return err
	}
	
	if FirmwareVersion, ok := PhonechangetopicuseragentinfoMap["firmwareVersion"].(string); ok {
		o.FirmwareVersion = &FirmwareVersion
	}
	
	if Manufacturer, ok := PhonechangetopicuseragentinfoMap["manufacturer"].(string); ok {
		o.Manufacturer = &Manufacturer
	}
	
	if Model, ok := PhonechangetopicuseragentinfoMap["model"].(string); ok {
		o.Model = &Model
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Phonechangetopicuseragentinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
