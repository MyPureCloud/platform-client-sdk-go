package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Realtimeadherenceexplanation
type Realtimeadherenceexplanation struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// StartDate - The start timestamp of the adherence explanation in ISO-8601 format
	StartDate *time.Time `json:"startDate,omitempty"`


	// LengthMinutes - The length of the adherence explanation in minutes
	LengthMinutes *int `json:"lengthMinutes,omitempty"`


	// Status - The status of the adherence explanation
	Status *string `json:"status,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Realtimeadherenceexplanation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Realtimeadherenceexplanation
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		LengthMinutes *int `json:"lengthMinutes,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		StartDate: StartDate,
		
		LengthMinutes: o.LengthMinutes,
		
		Status: o.Status,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Realtimeadherenceexplanation) UnmarshalJSON(b []byte) error {
	var RealtimeadherenceexplanationMap map[string]interface{}
	err := json.Unmarshal(b, &RealtimeadherenceexplanationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := RealtimeadherenceexplanationMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if startDateString, ok := RealtimeadherenceexplanationMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if LengthMinutes, ok := RealtimeadherenceexplanationMap["lengthMinutes"].(float64); ok {
		LengthMinutesInt := int(LengthMinutes)
		o.LengthMinutes = &LengthMinutesInt
	}
	
	if Status, ok := RealtimeadherenceexplanationMap["status"].(string); ok {
		o.Status = &Status
	}
	
	if SelfUri, ok := RealtimeadherenceexplanationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Realtimeadherenceexplanation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
