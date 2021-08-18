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

func (u *Objective) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Objective

	
	DateStart := new(string)
	if u.DateStart != nil {
		*DateStart = timeutil.Strftime(u.DateStart, "%Y-%m-%d")
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
		Id: u.Id,
		
		TemplateId: u.TemplateId,
		
		Zones: u.Zones,
		
		Enabled: u.Enabled,
		
		DateStart: DateStart,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Objective) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
