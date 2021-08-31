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

func (o *Twitterid) MarshalJSON() ([]byte, error) {
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
		Id: o.Id,
		
		Name: o.Name,
		
		ScreenName: o.ScreenName,
		
		Verified: o.Verified,
		
		ProfileUrl: o.ProfileUrl,
		Alias:    (*Alias)(o),
	})
}

func (o *Twitterid) UnmarshalJSON(b []byte) error {
	var TwitteridMap map[string]interface{}
	err := json.Unmarshal(b, &TwitteridMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TwitteridMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := TwitteridMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if ScreenName, ok := TwitteridMap["screenName"].(string); ok {
		o.ScreenName = &ScreenName
	}
	
	if Verified, ok := TwitteridMap["verified"].(bool); ok {
		o.Verified = &Verified
	}
	
	if ProfileUrl, ok := TwitteridMap["profileUrl"].(string); ok {
		o.ProfileUrl = &ProfileUrl
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Twitterid) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
