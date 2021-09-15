package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Assignedsegment
type Assignedsegment struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Segment - The ID of the segment assigned.
	Segment *Addressableentityref `json:"segment,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Assignedsegment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Assignedsegment
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Segment *Addressableentityref `json:"segment,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Segment: o.Segment,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Assignedsegment) UnmarshalJSON(b []byte) error {
	var AssignedsegmentMap map[string]interface{}
	err := json.Unmarshal(b, &AssignedsegmentMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AssignedsegmentMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Segment, ok := AssignedsegmentMap["segment"].(map[string]interface{}); ok {
		SegmentString, _ := json.Marshal(Segment)
		json.Unmarshal(SegmentString, &o.Segment)
	}
	
	if SelfUri, ok := AssignedsegmentMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Assignedsegment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
