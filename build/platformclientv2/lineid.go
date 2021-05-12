package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Lineid - User information for a Line account
type Lineid struct { 
	// Ids - The set of Line userIds that this person has. Each userId is specific to the Line channel that the user interacts with.
	Ids *[]Lineuserid `json:"ids,omitempty"`


	// DisplayName - The displayName of this person's account in Line
	DisplayName *string `json:"displayName,omitempty"`

}

// String returns a JSON representation of the model
func (o *Lineid) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
