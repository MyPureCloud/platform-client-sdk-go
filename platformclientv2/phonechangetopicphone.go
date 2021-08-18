package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Phonechangetopicphone
type Phonechangetopicphone struct { 
	// UserAgentInfo
	UserAgentInfo *Phonechangetopicuseragentinfo `json:"userAgentInfo,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`


	// Status
	Status *Phonechangetopicphonestatus `json:"status,omitempty"`


	// SecondaryStatus
	SecondaryStatus *Phonechangetopicphonestatus `json:"secondaryStatus,omitempty"`

}

func (u *Phonechangetopicphone) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Phonechangetopicphone

	

	return json.Marshal(&struct { 
		UserAgentInfo *Phonechangetopicuseragentinfo `json:"userAgentInfo,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Status *Phonechangetopicphonestatus `json:"status,omitempty"`
		
		SecondaryStatus *Phonechangetopicphonestatus `json:"secondaryStatus,omitempty"`
		*Alias
	}{ 
		UserAgentInfo: u.UserAgentInfo,
		
		Id: u.Id,
		
		Status: u.Status,
		
		SecondaryStatus: u.SecondaryStatus,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Phonechangetopicphone) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
