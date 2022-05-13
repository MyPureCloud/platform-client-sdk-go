package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createplanninggrouprequest
type Createplanninggrouprequest struct { 
	// Name - The name of the planning group
	Name *string `json:"name,omitempty"`


	// RoutePaths - Set of route paths to associate with the planning group
	RoutePaths *[]Routepathrequest `json:"routePaths,omitempty"`


	// ServiceGoalTemplateId - The ID of the service goal template to associate with this planning group
	ServiceGoalTemplateId *string `json:"serviceGoalTemplateId,omitempty"`

}

func (o *Createplanninggrouprequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createplanninggrouprequest
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		RoutePaths *[]Routepathrequest `json:"routePaths,omitempty"`
		
		ServiceGoalTemplateId *string `json:"serviceGoalTemplateId,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		RoutePaths: o.RoutePaths,
		
		ServiceGoalTemplateId: o.ServiceGoalTemplateId,
		Alias:    (*Alias)(o),
	})
}

func (o *Createplanninggrouprequest) UnmarshalJSON(b []byte) error {
	var CreateplanninggrouprequestMap map[string]interface{}
	err := json.Unmarshal(b, &CreateplanninggrouprequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := CreateplanninggrouprequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if RoutePaths, ok := CreateplanninggrouprequestMap["routePaths"].([]interface{}); ok {
		RoutePathsString, _ := json.Marshal(RoutePaths)
		json.Unmarshal(RoutePathsString, &o.RoutePaths)
	}
	
	if ServiceGoalTemplateId, ok := CreateplanninggrouprequestMap["serviceGoalTemplateId"].(string); ok {
		o.ServiceGoalTemplateId = &ServiceGoalTemplateId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Createplanninggrouprequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
