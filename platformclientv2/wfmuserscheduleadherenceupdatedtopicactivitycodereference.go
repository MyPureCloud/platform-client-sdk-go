package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmuserscheduleadherenceupdatedtopicactivitycodereference
type Wfmuserscheduleadherenceupdatedtopicactivitycodereference struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// SecondaryPresences
	SecondaryPresences *[]Wfmuserscheduleadherenceupdatedtopicsecondarypresencereference `json:"secondaryPresences,omitempty"`

}

func (o *Wfmuserscheduleadherenceupdatedtopicactivitycodereference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmuserscheduleadherenceupdatedtopicactivitycodereference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		SecondaryPresences *[]Wfmuserscheduleadherenceupdatedtopicsecondarypresencereference `json:"secondaryPresences,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		SecondaryPresences: o.SecondaryPresences,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmuserscheduleadherenceupdatedtopicactivitycodereference) UnmarshalJSON(b []byte) error {
	var WfmuserscheduleadherenceupdatedtopicactivitycodereferenceMap map[string]interface{}
	err := json.Unmarshal(b, &WfmuserscheduleadherenceupdatedtopicactivitycodereferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmuserscheduleadherenceupdatedtopicactivitycodereferenceMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if SecondaryPresences, ok := WfmuserscheduleadherenceupdatedtopicactivitycodereferenceMap["secondaryPresences"].([]interface{}); ok {
		SecondaryPresencesString, _ := json.Marshal(SecondaryPresences)
		json.Unmarshal(SecondaryPresencesString, &o.SecondaryPresences)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmuserscheduleadherenceupdatedtopicactivitycodereference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
