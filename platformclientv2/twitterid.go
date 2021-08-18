package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Twitterid) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Twitterid

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ScreenName *string `json:"screenName,omitempty"`
		
		Verified *bool `json:"verified,omitempty"`
		
		ProfileUrl *string `json:"profileUrl,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		ScreenName: u.ScreenName,
		
		Verified: u.Verified,
		
		ProfileUrl: u.ProfileUrl,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Twitterid) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
