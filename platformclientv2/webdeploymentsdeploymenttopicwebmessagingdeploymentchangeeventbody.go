package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webdeploymentsdeploymenttopicwebmessagingdeploymentchangeeventbody
type Webdeploymentsdeploymenttopicwebmessagingdeploymentchangeeventbody struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Configuration
	Configuration *Webdeploymentsdeploymenttopicwebmessagingconfigchangeeventbody `json:"configuration,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`

}

func (u *Webdeploymentsdeploymenttopicwebmessagingdeploymentchangeeventbody) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webdeploymentsdeploymenttopicwebmessagingdeploymentchangeeventbody

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Configuration *Webdeploymentsdeploymenttopicwebmessagingconfigchangeeventbody `json:"configuration,omitempty"`
		
		Status *string `json:"status,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Configuration: u.Configuration,
		
		Status: u.Status,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Webdeploymentsdeploymenttopicwebmessagingdeploymentchangeeventbody) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
