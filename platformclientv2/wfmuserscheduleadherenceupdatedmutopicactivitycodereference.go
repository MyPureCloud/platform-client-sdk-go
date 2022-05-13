package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmuserscheduleadherenceupdatedmutopicactivitycodereference
type Wfmuserscheduleadherenceupdatedmutopicactivitycodereference struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// SecondaryPresences
	SecondaryPresences *[]Wfmuserscheduleadherenceupdatedmutopicsecondarypresencereference `json:"secondaryPresences,omitempty"`

}

func (o *Wfmuserscheduleadherenceupdatedmutopicactivitycodereference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmuserscheduleadherenceupdatedmutopicactivitycodereference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		SecondaryPresences *[]Wfmuserscheduleadherenceupdatedmutopicsecondarypresencereference `json:"secondaryPresences,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		SecondaryPresences: o.SecondaryPresences,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmuserscheduleadherenceupdatedmutopicactivitycodereference) UnmarshalJSON(b []byte) error {
	var WfmuserscheduleadherenceupdatedmutopicactivitycodereferenceMap map[string]interface{}
	err := json.Unmarshal(b, &WfmuserscheduleadherenceupdatedmutopicactivitycodereferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmuserscheduleadherenceupdatedmutopicactivitycodereferenceMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if SecondaryPresences, ok := WfmuserscheduleadherenceupdatedmutopicactivitycodereferenceMap["secondaryPresences"].([]interface{}); ok {
		SecondaryPresencesString, _ := json.Marshal(SecondaryPresences)
		json.Unmarshal(SecondaryPresencesString, &o.SecondaryPresences)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmuserscheduleadherenceupdatedmutopicactivitycodereference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
