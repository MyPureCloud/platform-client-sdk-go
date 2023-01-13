package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webdeploymentsrefreshjwtrequest
type Webdeploymentsrefreshjwtrequest struct { 
	// RefreshToken - Refresh token used to issue a new JWT.
	RefreshToken *string `json:"refreshToken,omitempty"`


	// DeploymentId - The WebDeployment ID
	DeploymentId *string `json:"deploymentId,omitempty"`

}

func (o *Webdeploymentsrefreshjwtrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webdeploymentsrefreshjwtrequest
	
	return json.Marshal(&struct { 
		RefreshToken *string `json:"refreshToken,omitempty"`
		
		DeploymentId *string `json:"deploymentId,omitempty"`
		*Alias
	}{ 
		RefreshToken: o.RefreshToken,
		
		DeploymentId: o.DeploymentId,
		Alias:    (*Alias)(o),
	})
}

func (o *Webdeploymentsrefreshjwtrequest) UnmarshalJSON(b []byte) error {
	var WebdeploymentsrefreshjwtrequestMap map[string]interface{}
	err := json.Unmarshal(b, &WebdeploymentsrefreshjwtrequestMap)
	if err != nil {
		return err
	}
	
	if RefreshToken, ok := WebdeploymentsrefreshjwtrequestMap["refreshToken"].(string); ok {
		o.RefreshToken = &RefreshToken
	}
    
	if DeploymentId, ok := WebdeploymentsrefreshjwtrequestMap["deploymentId"].(string); ok {
		o.DeploymentId = &DeploymentId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Webdeploymentsrefreshjwtrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
