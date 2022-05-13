package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeywebeventsnotificationeventaction
type Journeywebeventsnotificationeventaction struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// CreatedDate
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// State
	State *string `json:"state,omitempty"`


	// MediaType
	MediaType *string `json:"mediaType,omitempty"`


	// Prompt
	Prompt *string `json:"prompt,omitempty"`

}

func (o *Journeywebeventsnotificationeventaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeywebeventsnotificationeventaction
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		Prompt *string `json:"prompt,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		CreatedDate: CreatedDate,
		
		State: o.State,
		
		MediaType: o.MediaType,
		
		Prompt: o.Prompt,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeywebeventsnotificationeventaction) UnmarshalJSON(b []byte) error {
	var JourneywebeventsnotificationeventactionMap map[string]interface{}
	err := json.Unmarshal(b, &JourneywebeventsnotificationeventactionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := JourneywebeventsnotificationeventactionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if createdDateString, ok := JourneywebeventsnotificationeventactionMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if State, ok := JourneywebeventsnotificationeventactionMap["state"].(string); ok {
		o.State = &State
	}
    
	if MediaType, ok := JourneywebeventsnotificationeventactionMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if Prompt, ok := JourneywebeventsnotificationeventactionMap["prompt"].(string); ok {
		o.Prompt = &Prompt
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Journeywebeventsnotificationeventaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
