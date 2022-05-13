package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Planninggroup
type Planninggroup struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ServiceGoalTemplate - The ID of the service goal template associated with this planning group
	ServiceGoalTemplate *Servicegoaltemplatereference `json:"serviceGoalTemplate,omitempty"`


	// RoutePaths - Set of route paths associated with the planning group
	RoutePaths *[]Routepathresponse `json:"routePaths,omitempty"`


	// Metadata - Version metadata for the planning group
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Planninggroup) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Planninggroup
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ServiceGoalTemplate *Servicegoaltemplatereference `json:"serviceGoalTemplate,omitempty"`
		
		RoutePaths *[]Routepathresponse `json:"routePaths,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		ServiceGoalTemplate: o.ServiceGoalTemplate,
		
		RoutePaths: o.RoutePaths,
		
		Metadata: o.Metadata,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Planninggroup) UnmarshalJSON(b []byte) error {
	var PlanninggroupMap map[string]interface{}
	err := json.Unmarshal(b, &PlanninggroupMap)
	if err != nil {
		return err
	}
	
	if Id, ok := PlanninggroupMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := PlanninggroupMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ServiceGoalTemplate, ok := PlanninggroupMap["serviceGoalTemplate"].(map[string]interface{}); ok {
		ServiceGoalTemplateString, _ := json.Marshal(ServiceGoalTemplate)
		json.Unmarshal(ServiceGoalTemplateString, &o.ServiceGoalTemplate)
	}
	
	if RoutePaths, ok := PlanninggroupMap["routePaths"].([]interface{}); ok {
		RoutePathsString, _ := json.Marshal(RoutePaths)
		json.Unmarshal(RoutePathsString, &o.RoutePaths)
	}
	
	if Metadata, ok := PlanninggroupMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if SelfUri, ok := PlanninggroupMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Planninggroup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
