package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Adherenceexplanationresponse
type Adherenceexplanationresponse struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Agent - The agent to whom this adherence explanation applies
	Agent *Userreference `json:"agent,omitempty"`


	// ManagementUnit - The management unit to which the agent belonged at the time the adherence explanation was submitted
	ManagementUnit *Managementunitreference `json:"managementUnit,omitempty"`


	// BusinessUnit - The business unit to which the agent belonged at the time the adherence explanation was submitted
	BusinessUnit *Businessunitreference `json:"businessUnit,omitempty"`


	// VarType - The type of the adherence explanation
	VarType *string `json:"type,omitempty"`


	// Status - The status of the adherence explanation
	Status *string `json:"status,omitempty"`


	// StartDate - The start timestamp of the adherence explanation in ISO-8601 format
	StartDate *time.Time `json:"startDate,omitempty"`


	// LengthMinutes - The length of the adherence explanation in minutes
	LengthMinutes *int `json:"lengthMinutes,omitempty"`


	// Notes - Notes about the adherence explanation
	Notes *string `json:"notes,omitempty"`


	// ReviewedBy - The user who reviewed the adherence explanation, if applicable
	ReviewedBy *Userreference `json:"reviewedBy,omitempty"`


	// ReviewedDate - The timestamp for when the adherence explanation was reviewed, if applicable. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ReviewedDate *time.Time `json:"reviewedDate,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Adherenceexplanationresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Adherenceexplanationresponse
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	ReviewedDate := new(string)
	if o.ReviewedDate != nil {
		
		*ReviewedDate = timeutil.Strftime(o.ReviewedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ReviewedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Agent *Userreference `json:"agent,omitempty"`
		
		ManagementUnit *Managementunitreference `json:"managementUnit,omitempty"`
		
		BusinessUnit *Businessunitreference `json:"businessUnit,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		LengthMinutes *int `json:"lengthMinutes,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		
		ReviewedBy *Userreference `json:"reviewedBy,omitempty"`
		
		ReviewedDate *string `json:"reviewedDate,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Agent: o.Agent,
		
		ManagementUnit: o.ManagementUnit,
		
		BusinessUnit: o.BusinessUnit,
		
		VarType: o.VarType,
		
		Status: o.Status,
		
		StartDate: StartDate,
		
		LengthMinutes: o.LengthMinutes,
		
		Notes: o.Notes,
		
		ReviewedBy: o.ReviewedBy,
		
		ReviewedDate: ReviewedDate,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Adherenceexplanationresponse) UnmarshalJSON(b []byte) error {
	var AdherenceexplanationresponseMap map[string]interface{}
	err := json.Unmarshal(b, &AdherenceexplanationresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AdherenceexplanationresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Agent, ok := AdherenceexplanationresponseMap["agent"].(map[string]interface{}); ok {
		AgentString, _ := json.Marshal(Agent)
		json.Unmarshal(AgentString, &o.Agent)
	}
	
	if ManagementUnit, ok := AdherenceexplanationresponseMap["managementUnit"].(map[string]interface{}); ok {
		ManagementUnitString, _ := json.Marshal(ManagementUnit)
		json.Unmarshal(ManagementUnitString, &o.ManagementUnit)
	}
	
	if BusinessUnit, ok := AdherenceexplanationresponseMap["businessUnit"].(map[string]interface{}); ok {
		BusinessUnitString, _ := json.Marshal(BusinessUnit)
		json.Unmarshal(BusinessUnitString, &o.BusinessUnit)
	}
	
	if VarType, ok := AdherenceexplanationresponseMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Status, ok := AdherenceexplanationresponseMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if startDateString, ok := AdherenceexplanationresponseMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if LengthMinutes, ok := AdherenceexplanationresponseMap["lengthMinutes"].(float64); ok {
		LengthMinutesInt := int(LengthMinutes)
		o.LengthMinutes = &LengthMinutesInt
	}
	
	if Notes, ok := AdherenceexplanationresponseMap["notes"].(string); ok {
		o.Notes = &Notes
	}
    
	if ReviewedBy, ok := AdherenceexplanationresponseMap["reviewedBy"].(map[string]interface{}); ok {
		ReviewedByString, _ := json.Marshal(ReviewedBy)
		json.Unmarshal(ReviewedByString, &o.ReviewedBy)
	}
	
	if reviewedDateString, ok := AdherenceexplanationresponseMap["reviewedDate"].(string); ok {
		ReviewedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", reviewedDateString)
		o.ReviewedDate = &ReviewedDate
	}
	
	if SelfUri, ok := AdherenceexplanationresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Adherenceexplanationresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
