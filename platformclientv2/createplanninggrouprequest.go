package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Createplanninggrouprequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
