package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Chatmessageuser
type Chatmessageuser struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DisplayName
	DisplayName *string `json:"displayName,omitempty"`


	// Username
	Username *string `json:"username,omitempty"`


	// Images
	Images *[]Userimage `json:"images,omitempty"`

}

func (u *Chatmessageuser) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Chatmessageuser

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DisplayName *string `json:"displayName,omitempty"`
		
		Username *string `json:"username,omitempty"`
		
		Images *[]Userimage `json:"images,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		DisplayName: u.DisplayName,
		
		Username: u.Username,
		
		Images: u.Images,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Chatmessageuser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
