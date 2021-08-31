package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialogflowagentsummary
type Dialogflowagentsummary struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Project - The project this Dialogflow agent belongs to
	Project *Dialogflowproject `json:"project,omitempty"`


	// Description - A description of the Dialogflow agent
	Description *string `json:"description,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Dialogflowagentsummary) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialogflowagentsummary
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Project *Dialogflowproject `json:"project,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Project: o.Project,
		
		Description: o.Description,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialogflowagentsummary) UnmarshalJSON(b []byte) error {
	var DialogflowagentsummaryMap map[string]interface{}
	err := json.Unmarshal(b, &DialogflowagentsummaryMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DialogflowagentsummaryMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := DialogflowagentsummaryMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Project, ok := DialogflowagentsummaryMap["project"].(map[string]interface{}); ok {
		ProjectString, _ := json.Marshal(Project)
		json.Unmarshal(ProjectString, &o.Project)
	}
	
	if Description, ok := DialogflowagentsummaryMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if SelfUri, ok := DialogflowagentsummaryMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialogflowagentsummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
