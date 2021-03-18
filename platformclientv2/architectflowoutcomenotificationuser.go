package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Architectflowoutcomenotificationuser
type Architectflowoutcomenotificationuser struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// HomeOrg
	HomeOrg *Architectflowoutcomenotificationhomeorganization `json:"homeOrg,omitempty"`

}

// String returns a JSON representation of the model
func (o *Architectflowoutcomenotificationuser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
