package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmagent - Workforce management agent data
type Wfmagent struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// User - The user associated with this data
	User *Userreference `json:"user,omitempty"`


	// WorkPlan - The work plan associated with this agent, if applicable
	WorkPlan *Workplanreference `json:"workPlan,omitempty"`


	// WorkPlanRotation - The work plan rotation associated with this agent, if applicable
	WorkPlanRotation *Workplanrotationreference `json:"workPlanRotation,omitempty"`


	// AcceptDirectShiftTrades - Whether the agent accepts direct shift trade requests
	AcceptDirectShiftTrades *bool `json:"acceptDirectShiftTrades,omitempty"`


	// Queues - List of queues to which this agent is capable of handling
	Queues *[]Queuereference `json:"queues,omitempty"`


	// Languages - The list of languages this agent is capable of handling
	Languages *[]Languagereference `json:"languages,omitempty"`


	// Skills - The list of skills this agent is capable of handling
	Skills *[]Routingskillreference `json:"skills,omitempty"`


	// Schedulable - Whether the agent has the permission to be included in schedule generation
	Schedulable *bool `json:"schedulable,omitempty"`


	// Metadata - Metadata for this agent
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Wfmagent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmagent
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		User *Userreference `json:"user,omitempty"`
		
		WorkPlan *Workplanreference `json:"workPlan,omitempty"`
		
		WorkPlanRotation *Workplanrotationreference `json:"workPlanRotation,omitempty"`
		
		AcceptDirectShiftTrades *bool `json:"acceptDirectShiftTrades,omitempty"`
		
		Queues *[]Queuereference `json:"queues,omitempty"`
		
		Languages *[]Languagereference `json:"languages,omitempty"`
		
		Skills *[]Routingskillreference `json:"skills,omitempty"`
		
		Schedulable *bool `json:"schedulable,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		User: o.User,
		
		WorkPlan: o.WorkPlan,
		
		WorkPlanRotation: o.WorkPlanRotation,
		
		AcceptDirectShiftTrades: o.AcceptDirectShiftTrades,
		
		Queues: o.Queues,
		
		Languages: o.Languages,
		
		Skills: o.Skills,
		
		Schedulable: o.Schedulable,
		
		Metadata: o.Metadata,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmagent) UnmarshalJSON(b []byte) error {
	var WfmagentMap map[string]interface{}
	err := json.Unmarshal(b, &WfmagentMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmagentMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if User, ok := WfmagentMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if WorkPlan, ok := WfmagentMap["workPlan"].(map[string]interface{}); ok {
		WorkPlanString, _ := json.Marshal(WorkPlan)
		json.Unmarshal(WorkPlanString, &o.WorkPlan)
	}
	
	if WorkPlanRotation, ok := WfmagentMap["workPlanRotation"].(map[string]interface{}); ok {
		WorkPlanRotationString, _ := json.Marshal(WorkPlanRotation)
		json.Unmarshal(WorkPlanRotationString, &o.WorkPlanRotation)
	}
	
	if AcceptDirectShiftTrades, ok := WfmagentMap["acceptDirectShiftTrades"].(bool); ok {
		o.AcceptDirectShiftTrades = &AcceptDirectShiftTrades
	}
	
	if Queues, ok := WfmagentMap["queues"].([]interface{}); ok {
		QueuesString, _ := json.Marshal(Queues)
		json.Unmarshal(QueuesString, &o.Queues)
	}
	
	if Languages, ok := WfmagentMap["languages"].([]interface{}); ok {
		LanguagesString, _ := json.Marshal(Languages)
		json.Unmarshal(LanguagesString, &o.Languages)
	}
	
	if Skills, ok := WfmagentMap["skills"].([]interface{}); ok {
		SkillsString, _ := json.Marshal(Skills)
		json.Unmarshal(SkillsString, &o.Skills)
	}
	
	if Schedulable, ok := WfmagentMap["schedulable"].(bool); ok {
		o.Schedulable = &Schedulable
	}
	
	if Metadata, ok := WfmagentMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if SelfUri, ok := WfmagentMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmagent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
