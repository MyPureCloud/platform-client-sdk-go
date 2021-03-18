package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Architectflowoutcomenotificationhomeorganization
type Architectflowoutcomenotificationhomeorganization struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ThirdPartyOrgName
	ThirdPartyOrgName *string `json:"thirdPartyOrgName,omitempty"`

}

// String returns a JSON representation of the model
func (o *Architectflowoutcomenotificationhomeorganization) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
