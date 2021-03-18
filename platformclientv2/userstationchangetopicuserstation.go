package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Userstationchangetopicuserstation
type Userstationchangetopicuserstation struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// AssociatedUser
	AssociatedUser *Userstationchangetopicuser `json:"associatedUser,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userstationchangetopicuserstation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
