package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Objective
type Objective struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// TemplateId - The id of this objective's base template
	TemplateId *string `json:"templateId,omitempty"`


	// Zones - Objective zone specifies min,max points and values for the associated metric
	Zones *[]Objectivezone `json:"zones,omitempty"`


	// Enabled - A flag for whether this objective is enabled for the related metric
	Enabled *bool `json:"enabled,omitempty"`


	// DateStart - start date of the objective. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateStart *time.Time `json:"dateStart,omitempty"`

}

func (o *Objective) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Objective
	
	DateStart := new(string)
	if o.DateStart != nil {
		*DateStart = timeutil.Strftime(o.DateStart, "%Y-%m-%d")
	} else {
		DateStart = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		TemplateId *string `json:"templateId,omitempty"`
		
		Zones *[]Objectivezone `json:"zones,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		DateStart *string `json:"dateStart,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		TemplateId: o.TemplateId,
		
		Zones: o.Zones,
		
		Enabled: o.Enabled,
		
		DateStart: DateStart,
		Alias:    (*Alias)(o),
	})
}

func (o *Objective) UnmarshalJSON(b []byte) error {
	var ObjectiveMap map[string]interface{}
	err := json.Unmarshal(b, &ObjectiveMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ObjectiveMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if TemplateId, ok := ObjectiveMap["templateId"].(string); ok {
		o.TemplateId = &TemplateId
	}
	
	if Zones, ok := ObjectiveMap["zones"].([]interface{}); ok {
		ZonesString, _ := json.Marshal(Zones)
		json.Unmarshal(ZonesString, &o.Zones)
	}
	
	if Enabled, ok := ObjectiveMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
	
	if dateStartString, ok := ObjectiveMap["dateStart"].(string); ok {
		DateStart, _ := time.Parse("2006-01-02", dateStartString)
		o.DateStart = &DateStart
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Objective) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
