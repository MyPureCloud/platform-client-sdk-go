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

func (o *Channel) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Channel
	
	Expires := new(string)
	if o.Expires != nil {
		
		*Expires = timeutil.Strftime(o.Expires, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Expires = nil
	}
	
	return json.Marshal(&struct { 
		ConnectUri *string `json:"connectUri,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Expires *string `json:"expires,omitempty"`
		*Alias
	}{ 
		ConnectUri: o.ConnectUri,
		
		Id: o.Id,
		
		Expires: Expires,
		Alias:    (*Alias)(o),
	})
}

func (o *Channel) UnmarshalJSON(b []byte) error {
	var ChannelMap map[string]interface{}
	err := json.Unmarshal(b, &ChannelMap)
	if err != nil {
		return err
	}
	
	if ConnectUri, ok := ChannelMap["connectUri"].(string); ok {
		o.ConnectUri = &ConnectUri
	}
    
	if Id, ok := ChannelMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if expiresString, ok := ChannelMap["expires"].(string); ok {
		Expires, _ := time.Parse("2006-01-02T15:04:05.999999Z", expiresString)
		o.Expires = &Expires
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Channel) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
