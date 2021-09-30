package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotflowmilestone
type Textbotflowmilestone struct { 
	// Id - The Milestone's ID.
	Id *string `json:"id,omitempty"`


	// DateReached - The timestamp of when the milestone was reached. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateReached *time.Time `json:"dateReached,omitempty"`


	// Sequence - The sequence number of the milestone.
	Sequence *int `json:"sequence,omitempty"`

}

func (o *Textbotflowmilestone) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Textbotflowmilestone
	
	DateReached := new(string)
	if o.DateReached != nil {
		
		*DateReached = timeutil.Strftime(o.DateReached, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateReached = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		DateReached *string `json:"dateReached,omitempty"`
		
		Sequence *int `json:"sequence,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		DateReached: DateReached,
		
		Sequence: o.Sequence,
		Alias:    (*Alias)(o),
	})
}

func (o *Textbotflowmilestone) UnmarshalJSON(b []byte) error {
	var TextbotflowmilestoneMap map[string]interface{}
	err := json.Unmarshal(b, &TextbotflowmilestoneMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TextbotflowmilestoneMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if dateReachedString, ok := TextbotflowmilestoneMap["dateReached"].(string); ok {
		DateReached, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateReachedString)
		o.DateReached = &DateReached
	}
	
	if Sequence, ok := TextbotflowmilestoneMap["sequence"].(float64); ok {
		SequenceInt := int(Sequence)
		o.Sequence = &SequenceInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Textbotflowmilestone) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
