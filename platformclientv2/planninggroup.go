package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Planninggroup - Planning Group
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

func (u *Planninggroup) MarshalJSON() ([]byte, error) {
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
		Id: u.Id,
		
		Name: u.Name,
		
		ServiceGoalTemplate: u.ServiceGoalTemplate,
		
		RoutePaths: u.RoutePaths,
		
		Metadata: u.Metadata,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Planninggroup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
