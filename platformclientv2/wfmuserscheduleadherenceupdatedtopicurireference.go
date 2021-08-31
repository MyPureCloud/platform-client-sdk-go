package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmuserscheduleadherenceupdatedtopicurireference
type Wfmuserscheduleadherenceupdatedtopicurireference struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`

}

func (o *Wfmuserscheduleadherenceupdatedtopicurireference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmuserscheduleadherenceupdatedtopicurireference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmuserscheduleadherenceupdatedtopicurireference) UnmarshalJSON(b []byte) error {
	var WfmuserscheduleadherenceupdatedtopicurireferenceMap map[string]interface{}
	err := json.Unmarshal(b, &WfmuserscheduleadherenceupdatedtopicurireferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmuserscheduleadherenceupdatedtopicurireferenceMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := WfmuserscheduleadherenceupdatedtopicurireferenceMap["name"].(string); ok {
		o.Name = &Name
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmuserscheduleadherenceupdatedtopicurireference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
