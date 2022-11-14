package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Addadherenceexplanationadminrequest
type Addadherenceexplanationadminrequest struct { 
	// VarType - The type of the adherence explanation
	VarType *string `json:"type,omitempty"`


	// StartDate - The start timestamp of the adherence explanation in ISO-8601 format
	StartDate *time.Time `json:"startDate,omitempty"`


	// LengthMinutes - The length of the adherence explanation in minutes
	LengthMinutes *int `json:"lengthMinutes,omitempty"`


	// Notes - Notes about the adherence explanation
	Notes *string `json:"notes,omitempty"`


	// Status - The status of the adherence explanation
	Status *string `json:"status,omitempty"`

}

func (o *Addadherenceexplanationadminrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Addadherenceexplanationadminrequest
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		LengthMinutes *int `json:"lengthMinutes,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		
		Status *string `json:"status,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		StartDate: StartDate,
		
		LengthMinutes: o.LengthMinutes,
		
		Notes: o.Notes,
		
		Status: o.Status,
		Alias:    (*Alias)(o),
	})
}

func (o *Addadherenceexplanationadminrequest) UnmarshalJSON(b []byte) error {
	var AddadherenceexplanationadminrequestMap map[string]interface{}
	err := json.Unmarshal(b, &AddadherenceexplanationadminrequestMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := AddadherenceexplanationadminrequestMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if startDateString, ok := AddadherenceexplanationadminrequestMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if LengthMinutes, ok := AddadherenceexplanationadminrequestMap["lengthMinutes"].(float64); ok {
		LengthMinutesInt := int(LengthMinutes)
		o.LengthMinutes = &LengthMinutesInt
	}
	
	if Notes, ok := AddadherenceexplanationadminrequestMap["notes"].(string); ok {
		o.Notes = &Notes
	}
    
	if Status, ok := AddadherenceexplanationadminrequestMap["status"].(string); ok {
		o.Status = &Status
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Addadherenceexplanationadminrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
