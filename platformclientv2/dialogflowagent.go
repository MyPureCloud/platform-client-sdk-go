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

func (u *Dialogflowagent) MarshalJSON() ([]byte, error) {
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
		Id: u.Id,
		
		Name: u.Name,
		
		Project: u.Project,
		
		Languages: u.Languages,
		
		Intents: u.Intents,
		
		Environments: u.Environments,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dialogflowagent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
