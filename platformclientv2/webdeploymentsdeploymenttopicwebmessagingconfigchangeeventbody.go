package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webdeploymentsdeploymenttopicwebmessagingconfigchangeeventbody
type Webdeploymentsdeploymenttopicwebmessagingconfigchangeeventbody struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Version
	Version *string `json:"version,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`

}

func (u *Webdeploymentsdeploymenttopicwebmessagingconfigchangeeventbody) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webdeploymentsdeploymenttopicwebmessagingconfigchangeeventbody

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Version *string `json:"version,omitempty"`
		
		Status *string `json:"status,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Version: u.Version,
		
		Status: u.Status,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Webdeploymentsdeploymenttopicwebmessagingconfigchangeeventbody) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
