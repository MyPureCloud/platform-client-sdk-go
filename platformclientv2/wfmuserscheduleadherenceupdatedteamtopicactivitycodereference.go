package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmuserscheduleadherenceupdatedteamtopicactivitycodereference
type Wfmuserscheduleadherenceupdatedteamtopicactivitycodereference struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// SecondaryPresences
	SecondaryPresences *[]Wfmuserscheduleadherenceupdatedteamtopicsecondarypresencereference `json:"secondaryPresences,omitempty"`

}

func (o *Wfmuserscheduleadherenceupdatedteamtopicactivitycodereference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmuserscheduleadherenceupdatedteamtopicactivitycodereference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		SecondaryPresences *[]Wfmuserscheduleadherenceupdatedteamtopicsecondarypresencereference `json:"secondaryPresences,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		SecondaryPresences: o.SecondaryPresences,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmuserscheduleadherenceupdatedteamtopicactivitycodereference) UnmarshalJSON(b []byte) error {
	var WfmuserscheduleadherenceupdatedteamtopicactivitycodereferenceMap map[string]interface{}
	err := json.Unmarshal(b, &WfmuserscheduleadherenceupdatedteamtopicactivitycodereferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmuserscheduleadherenceupdatedteamtopicactivitycodereferenceMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if SecondaryPresences, ok := WfmuserscheduleadherenceupdatedteamtopicactivitycodereferenceMap["secondaryPresences"].([]interface{}); ok {
		SecondaryPresencesString, _ := json.Marshal(SecondaryPresences)
		json.Unmarshal(SecondaryPresencesString, &o.SecondaryPresences)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmuserscheduleadherenceupdatedteamtopicactivitycodereference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
