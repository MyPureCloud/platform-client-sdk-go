package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Channel
type Channel struct { 
	// ConnectUri
	ConnectUri *string `json:"connectUri,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`


	// Expires - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Expires *time.Time `json:"expires,omitempty"`

}

func (u *Channel) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Channel

	
	Expires := new(string)
	if u.Expires != nil {
		
		*Expires = timeutil.Strftime(u.Expires, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Expires = nil
	}
	

	return json.Marshal(&struct { 
		ConnectUri *string `json:"connectUri,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Expires *string `json:"expires,omitempty"`
		*Alias
	}{ 
		ConnectUri: u.ConnectUri,
		
		Id: u.Id,
		
		Expires: Expires,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Channel) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
