package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Updateplanninggrouprequest
type Updateplanninggrouprequest struct { 
	// Name - The name of the planning group
	Name *string `json:"name,omitempty"`


	// RoutePaths - Set of route paths to associate with the planning group
	RoutePaths *Setwrapperroutepathrequest `json:"routePaths,omitempty"`


	// ServiceGoalTemplateId - The ID of the service goal template to associate with this planning group
	ServiceGoalTemplateId *string `json:"serviceGoalTemplateId,omitempty"`


	// Metadata - Version metadata for the planning group
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

}

func (o *Updateplanninggrouprequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Updateplanninggrouprequest
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		RoutePaths *Setwrapperroutepathrequest `json:"routePaths,omitempty"`
		
		ServiceGoalTemplateId *string `json:"serviceGoalTemplateId,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		RoutePaths: o.RoutePaths,
		
		ServiceGoalTemplateId: o.ServiceGoalTemplateId,
		
		Metadata: o.Metadata,
		Alias:    (*Alias)(o),
	})
}

func (o *Updateplanninggrouprequest) UnmarshalJSON(b []byte) error {
	var UpdateplanninggrouprequestMap map[string]interface{}
	err := json.Unmarshal(b, &UpdateplanninggrouprequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := UpdateplanninggrouprequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if RoutePaths, ok := UpdateplanninggrouprequestMap["routePaths"].(map[string]interface{}); ok {
		RoutePathsString, _ := json.Marshal(RoutePaths)
		json.Unmarshal(RoutePathsString, &o.RoutePaths)
	}
	
	if ServiceGoalTemplateId, ok := UpdateplanninggrouprequestMap["serviceGoalTemplateId"].(string); ok {
		o.ServiceGoalTemplateId = &ServiceGoalTemplateId
	}
    
	if Metadata, ok := UpdateplanninggrouprequestMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Updateplanninggrouprequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
