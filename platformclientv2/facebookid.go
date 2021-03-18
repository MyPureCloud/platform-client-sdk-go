package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Facebookid - User information for a Facebook user interacting with a page or app
type Facebookid struct { 
	// Ids - The set of scopedIds that this person has. Each scopedId is specific to a page or app that the user interacts with.
	Ids *[]Facebookscopedid `json:"ids,omitempty"`


	// DisplayName - The displayName of this person's Facebook account. Roughly translates to user.first_name + ' ' + user.last_name in the Facebook API.
	DisplayName *string `json:"displayName,omitempty"`

}

// String returns a JSON representation of the model
func (o *Facebookid) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
