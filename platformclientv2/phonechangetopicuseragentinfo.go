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

func (u *Phonechangetopicuseragentinfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Phonechangetopicuseragentinfo

	

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
func (o *Phonechangetopicuseragentinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
