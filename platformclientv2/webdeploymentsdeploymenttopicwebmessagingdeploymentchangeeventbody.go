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

func (o *Webdeploymentsdeploymenttopicwebmessagingdeploymentchangeeventbody) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webdeploymentsdeploymenttopicwebmessagingdeploymentchangeeventbody
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Configuration *Webdeploymentsdeploymenttopicwebmessagingconfigchangeeventbody `json:"configuration,omitempty"`
		
		Status *string `json:"status,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Configuration: o.Configuration,
		
		Status: o.Status,
		Alias:    (*Alias)(o),
	})
}

func (o *Webdeploymentsdeploymenttopicwebmessagingdeploymentchangeeventbody) UnmarshalJSON(b []byte) error {
	var WebdeploymentsdeploymenttopicwebmessagingdeploymentchangeeventbodyMap map[string]interface{}
	err := json.Unmarshal(b, &WebdeploymentsdeploymenttopicwebmessagingdeploymentchangeeventbodyMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WebdeploymentsdeploymenttopicwebmessagingdeploymentchangeeventbodyMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Configuration, ok := WebdeploymentsdeploymenttopicwebmessagingdeploymentchangeeventbodyMap["configuration"].(map[string]interface{}); ok {
		ConfigurationString, _ := json.Marshal(Configuration)
		json.Unmarshal(ConfigurationString, &o.Configuration)
	}
	
	if Status, ok := WebdeploymentsdeploymenttopicwebmessagingdeploymentchangeeventbodyMap["status"].(string); ok {
		o.Status = &Status
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Webdeploymentsdeploymenttopicwebmessagingdeploymentchangeeventbody) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
