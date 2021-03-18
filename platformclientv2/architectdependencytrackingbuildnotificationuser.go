package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Architectdependencytrackingbuildnotificationuser
type Architectdependencytrackingbuildnotificationuser struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// HomeOrg
	HomeOrg *Architectdependencytrackingbuildnotificationhomeorganization `json:"homeOrg,omitempty"`

}

// String returns a JSON representation of the model
func (o *Architectdependencytrackingbuildnotificationuser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
