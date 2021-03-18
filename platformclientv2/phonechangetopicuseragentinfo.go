package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Phonechangetopicuseragentinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
