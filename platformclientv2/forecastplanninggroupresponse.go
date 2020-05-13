package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Forecastplanninggroupresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
