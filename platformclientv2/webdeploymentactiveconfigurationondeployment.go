package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webdeploymentactiveconfigurationondeployment - Details about the active configuration on a deployment
type Webdeploymentactiveconfigurationondeployment struct { 
	// ConfigurationVersion - The active configuration on a deployment
	ConfigurationVersion *Webdeploymentconfigurationversion `json:"configurationVersion,omitempty"`


	// Deployment - The web deployment associated with the active configuration
	Deployment *Webdeployment `json:"deployment,omitempty"`

}

func (o *Webdeploymentactiveconfigurationondeployment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webdeploymentactiveconfigurationondeployment
	
	return json.Marshal(&struct { 
		ConfigurationVersion *Webdeploymentconfigurationversion `json:"configurationVersion,omitempty"`
		
		Deployment *Webdeployment `json:"deployment,omitempty"`
		*Alias
	}{ 
		ConfigurationVersion: o.ConfigurationVersion,
		
		Deployment: o.Deployment,
		Alias:    (*Alias)(o),
	})
}

func (o *Webdeploymentactiveconfigurationondeployment) UnmarshalJSON(b []byte) error {
	var WebdeploymentactiveconfigurationondeploymentMap map[string]interface{}
	err := json.Unmarshal(b, &WebdeploymentactiveconfigurationondeploymentMap)
	if err != nil {
		return err
	}
	
	if ConfigurationVersion, ok := WebdeploymentactiveconfigurationondeploymentMap["configurationVersion"].(map[string]interface{}); ok {
		ConfigurationVersionString, _ := json.Marshal(ConfigurationVersion)
		json.Unmarshal(ConfigurationVersionString, &o.ConfigurationVersion)
	}
	
	if Deployment, ok := WebdeploymentactiveconfigurationondeploymentMap["deployment"].(map[string]interface{}); ok {
		DeploymentString, _ := json.Marshal(Deployment)
		json.Unmarshal(DeploymentString, &o.Deployment)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Webdeploymentactiveconfigurationondeployment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
