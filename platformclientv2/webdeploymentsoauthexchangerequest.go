package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webdeploymentsoauthexchangerequest
type Webdeploymentsoauthexchangerequest struct { 
	// DeploymentId - The WebDeployment ID
	DeploymentId *string `json:"deploymentId,omitempty"`


	// JourneyContext - A Customer journey context.
	JourneyContext *Webdeploymentsjourneycontext `json:"journeyContext,omitempty"`


	// Oauth
	Oauth *Webdeploymentsoauthrequestparameters `json:"oauth,omitempty"`

}

func (o *Webdeploymentsoauthexchangerequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webdeploymentsoauthexchangerequest
	
	return json.Marshal(&struct { 
		DeploymentId *string `json:"deploymentId,omitempty"`
		
		JourneyContext *Webdeploymentsjourneycontext `json:"journeyContext,omitempty"`
		
		Oauth *Webdeploymentsoauthrequestparameters `json:"oauth,omitempty"`
		*Alias
	}{ 
		DeploymentId: o.DeploymentId,
		
		JourneyContext: o.JourneyContext,
		
		Oauth: o.Oauth,
		Alias:    (*Alias)(o),
	})
}

func (o *Webdeploymentsoauthexchangerequest) UnmarshalJSON(b []byte) error {
	var WebdeploymentsoauthexchangerequestMap map[string]interface{}
	err := json.Unmarshal(b, &WebdeploymentsoauthexchangerequestMap)
	if err != nil {
		return err
	}
	
	if DeploymentId, ok := WebdeploymentsoauthexchangerequestMap["deploymentId"].(string); ok {
		o.DeploymentId = &DeploymentId
	}
    
	if JourneyContext, ok := WebdeploymentsoauthexchangerequestMap["journeyContext"].(map[string]interface{}); ok {
		JourneyContextString, _ := json.Marshal(JourneyContext)
		json.Unmarshal(JourneyContextString, &o.JourneyContext)
	}
	
	if Oauth, ok := WebdeploymentsoauthexchangerequestMap["oauth"].(map[string]interface{}); ok {
		OauthString, _ := json.Marshal(Oauth)
		json.Unmarshal(OauthString, &o.Oauth)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Webdeploymentsoauthexchangerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
