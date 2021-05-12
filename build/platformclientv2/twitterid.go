package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Twitterid - User information for a twitter account
type Twitterid struct { 
	// Id - twitter user.id_str
	Id *string `json:"id,omitempty"`


	// Name - twitter user.name
	Name *string `json:"name,omitempty"`


	// ScreenName - twitter user.screen_name
	ScreenName *string `json:"screenName,omitempty"`


	// Verified - whether this data has been verified using the twitter API
	Verified *bool `json:"verified,omitempty"`


	// ProfileUrl - url of user's twitter profile
	ProfileUrl *string `json:"profileUrl,omitempty"`

}

// String returns a JSON representation of the model
func (o *Twitterid) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
