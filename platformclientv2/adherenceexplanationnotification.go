package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Adherenceexplanationnotification
type Adherenceexplanationnotification struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Agent - The agent for whom the adherence explanation applies
	Agent *Userreference `json:"agent,omitempty"`


	// ManagementUnit - The management unit to which the agent belonged at the time the adherence explanation was submitted
	ManagementUnit *Managementunitreference `json:"managementUnit,omitempty"`


	// BusinessUnit - The business unit to which the agent belonged at the time the adherence explanation was submitted
	BusinessUnit *Businessunitreference `json:"businessUnit,omitempty"`


	// StartDate - The start date of the adherence explanation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartDate *time.Time `json:"startDate,omitempty"`


	// LengthMinutes - The length of the adherence explanation in minutes
	LengthMinutes *int `json:"lengthMinutes,omitempty"`


	// Status - The status of the adherence explanation
	Status *string `json:"status,omitempty"`


	// VarType - The type of the adherence explanation
	VarType *string `json:"type,omitempty"`


	// Notes - Notes about the adherence explanation
	Notes *string `json:"notes,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Adherenceexplanationnotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Adherenceexplanationnotification
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Agent *Userreference `json:"agent,omitempty"`
		
		ManagementUnit *Managementunitreference `json:"managementUnit,omitempty"`
		
		BusinessUnit *Businessunitreference `json:"businessUnit,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		LengthMinutes *int `json:"lengthMinutes,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Agent: o.Agent,
		
		ManagementUnit: o.ManagementUnit,
		
		BusinessUnit: o.BusinessUnit,
		
		StartDate: StartDate,
		
		LengthMinutes: o.LengthMinutes,
		
		Status: o.Status,
		
		VarType: o.VarType,
		
		Notes: o.Notes,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Adherenceexplanationnotification) UnmarshalJSON(b []byte) error {
	var AdherenceexplanationnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &AdherenceexplanationnotificationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AdherenceexplanationnotificationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Agent, ok := AdherenceexplanationnotificationMap["agent"].(map[string]interface{}); ok {
		AgentString, _ := json.Marshal(Agent)
		json.Unmarshal(AgentString, &o.Agent)
	}
	
	if ManagementUnit, ok := AdherenceexplanationnotificationMap["managementUnit"].(map[string]interface{}); ok {
		ManagementUnitString, _ := json.Marshal(ManagementUnit)
		json.Unmarshal(ManagementUnitString, &o.ManagementUnit)
	}
	
	if BusinessUnit, ok := AdherenceexplanationnotificationMap["businessUnit"].(map[string]interface{}); ok {
		BusinessUnitString, _ := json.Marshal(BusinessUnit)
		json.Unmarshal(BusinessUnitString, &o.BusinessUnit)
	}
	
	if startDateString, ok := AdherenceexplanationnotificationMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if LengthMinutes, ok := AdherenceexplanationnotificationMap["lengthMinutes"].(float64); ok {
		LengthMinutesInt := int(LengthMinutes)
		o.LengthMinutes = &LengthMinutesInt
	}
	
	if Status, ok := AdherenceexplanationnotificationMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if VarType, ok := AdherenceexplanationnotificationMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Notes, ok := AdherenceexplanationnotificationMap["notes"].(string); ok {
		o.Notes = &Notes
	}
    
	if SelfUri, ok := AdherenceexplanationnotificationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Adherenceexplanationnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
