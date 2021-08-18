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

func (u *Forecastplanninggroupresponse) MarshalJSON() ([]byte, error) {
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
		Id: u.Id,
		
		Name: u.Name,
		
		RoutePaths: u.RoutePaths,
		
		ServiceGoalTemplate: u.ServiceGoalTemplate,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Forecastplanninggroupresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
