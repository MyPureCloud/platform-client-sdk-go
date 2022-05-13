package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialogflowagent
type Dialogflowagent struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Project - The project this Dialogflow agent belongs to
	Project *Dialogflowproject `json:"project,omitempty"`


	// Languages - The supported languages of the Dialogflow agent
	Languages *[]string `json:"languages,omitempty"`


	// Intents - An array of Intents associated with this agent
	Intents *[]Dialogflowintent `json:"intents,omitempty"`


	// Environments - Available environments for this agent
	Environments *[]string `json:"environments,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Dialogflowagent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialogflowagent
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Project *Dialogflowproject `json:"project,omitempty"`
		
		Languages *[]string `json:"languages,omitempty"`
		
		Intents *[]Dialogflowintent `json:"intents,omitempty"`
		
		Environments *[]string `json:"environments,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Project: o.Project,
		
		Languages: o.Languages,
		
		Intents: o.Intents,
		
		Environments: o.Environments,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialogflowagent) UnmarshalJSON(b []byte) error {
	var DialogflowagentMap map[string]interface{}
	err := json.Unmarshal(b, &DialogflowagentMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DialogflowagentMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := DialogflowagentMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Project, ok := DialogflowagentMap["project"].(map[string]interface{}); ok {
		ProjectString, _ := json.Marshal(Project)
		json.Unmarshal(ProjectString, &o.Project)
	}
	
	if Languages, ok := DialogflowagentMap["languages"].([]interface{}); ok {
		LanguagesString, _ := json.Marshal(Languages)
		json.Unmarshal(LanguagesString, &o.Languages)
	}
	
	if Intents, ok := DialogflowagentMap["intents"].([]interface{}); ok {
		IntentsString, _ := json.Marshal(Intents)
		json.Unmarshal(IntentsString, &o.Intents)
	}
	
	if Environments, ok := DialogflowagentMap["environments"].([]interface{}); ok {
		EnvironmentsString, _ := json.Marshal(Environments)
		json.Unmarshal(EnvironmentsString, &o.Environments)
	}
	
	if SelfUri, ok := DialogflowagentMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Dialogflowagent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
