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

func (u *Updateplanninggrouprequest) MarshalJSON() ([]byte, error) {
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
		Name: u.Name,
		
		RoutePaths: u.RoutePaths,
		
		ServiceGoalTemplateId: u.ServiceGoalTemplateId,
		
		Metadata: u.Metadata,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Updateplanninggrouprequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
