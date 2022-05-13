package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmusernotificationtopicadherenceexplanationnotification
type Wfmusernotificationtopicadherenceexplanationnotification struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Agent
	Agent *Wfmusernotificationtopicuserreference `json:"agent,omitempty"`


	// ManagementUnit
	ManagementUnit *Wfmusernotificationtopicmanagementunit `json:"managementUnit,omitempty"`


	// BusinessUnit
	BusinessUnit *Wfmusernotificationtopicbusinessunit `json:"businessUnit,omitempty"`


	// StartDate
	StartDate *time.Time `json:"startDate,omitempty"`


	// LengthMinutes
	LengthMinutes *int `json:"lengthMinutes,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// Notes
	Notes *string `json:"notes,omitempty"`

}

func (o *Wfmusernotificationtopicadherenceexplanationnotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmusernotificationtopicadherenceexplanationnotification
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Agent *Wfmusernotificationtopicuserreference `json:"agent,omitempty"`
		
		ManagementUnit *Wfmusernotificationtopicmanagementunit `json:"managementUnit,omitempty"`
		
		BusinessUnit *Wfmusernotificationtopicbusinessunit `json:"businessUnit,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		LengthMinutes *int `json:"lengthMinutes,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
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
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmusernotificationtopicadherenceexplanationnotification) UnmarshalJSON(b []byte) error {
	var WfmusernotificationtopicadherenceexplanationnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &WfmusernotificationtopicadherenceexplanationnotificationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmusernotificationtopicadherenceexplanationnotificationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Agent, ok := WfmusernotificationtopicadherenceexplanationnotificationMap["agent"].(map[string]interface{}); ok {
		AgentString, _ := json.Marshal(Agent)
		json.Unmarshal(AgentString, &o.Agent)
	}
	
	if ManagementUnit, ok := WfmusernotificationtopicadherenceexplanationnotificationMap["managementUnit"].(map[string]interface{}); ok {
		ManagementUnitString, _ := json.Marshal(ManagementUnit)
		json.Unmarshal(ManagementUnitString, &o.ManagementUnit)
	}
	
	if BusinessUnit, ok := WfmusernotificationtopicadherenceexplanationnotificationMap["businessUnit"].(map[string]interface{}); ok {
		BusinessUnitString, _ := json.Marshal(BusinessUnit)
		json.Unmarshal(BusinessUnitString, &o.BusinessUnit)
	}
	
	if startDateString, ok := WfmusernotificationtopicadherenceexplanationnotificationMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if LengthMinutes, ok := WfmusernotificationtopicadherenceexplanationnotificationMap["lengthMinutes"].(float64); ok {
		LengthMinutesInt := int(LengthMinutes)
		o.LengthMinutes = &LengthMinutesInt
	}
	
	if Status, ok := WfmusernotificationtopicadherenceexplanationnotificationMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if VarType, ok := WfmusernotificationtopicadherenceexplanationnotificationMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Notes, ok := WfmusernotificationtopicadherenceexplanationnotificationMap["notes"].(string); ok {
		o.Notes = &Notes
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmusernotificationtopicadherenceexplanationnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
