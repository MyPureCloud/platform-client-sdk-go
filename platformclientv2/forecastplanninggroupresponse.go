package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Forecastplanninggroupresponse
type Forecastplanninggroupresponse struct { 
	// Id - The ID of the planning group
	Id *string `json:"id,omitempty"`


	// Name - The name of the planning group
	Name *string `json:"name,omitempty"`


	// RoutePaths - Route path configuration for this planning group
	RoutePaths *[]Routepathresponse `json:"routePaths,omitempty"`


	// ServiceGoalTemplate - Service goals for this planning group
	ServiceGoalTemplate *Forecastservicegoaltemplateresponse `json:"serviceGoalTemplate,omitempty"`

}

func (o *Forecastplanninggroupresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Forecastplanninggroupresponse
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		RoutePaths *[]Routepathresponse `json:"routePaths,omitempty"`
		
		ServiceGoalTemplate *Forecastservicegoaltemplateresponse `json:"serviceGoalTemplate,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		RoutePaths: o.RoutePaths,
		
		ServiceGoalTemplate: o.ServiceGoalTemplate,
		Alias:    (*Alias)(o),
	})
}

func (o *Forecastplanninggroupresponse) UnmarshalJSON(b []byte) error {
	var ForecastplanninggroupresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ForecastplanninggroupresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ForecastplanninggroupresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ForecastplanninggroupresponseMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if RoutePaths, ok := ForecastplanninggroupresponseMap["routePaths"].([]interface{}); ok {
		RoutePathsString, _ := json.Marshal(RoutePaths)
		json.Unmarshal(RoutePathsString, &o.RoutePaths)
	}
	
	if ServiceGoalTemplate, ok := ForecastplanninggroupresponseMap["serviceGoalTemplate"].(map[string]interface{}); ok {
		ServiceGoalTemplateString, _ := json.Marshal(ServiceGoalTemplate)
		json.Unmarshal(ServiceGoalTemplateString, &o.ServiceGoalTemplate)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Forecastplanninggroupresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
