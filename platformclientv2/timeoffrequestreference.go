package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Timeoffrequestreference
type Timeoffrequestreference struct { 
	// Id - The id of the time off request
	Id *string `json:"id,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Timeoffrequestreference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Timeoffrequestreference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Timeoffrequestreference) UnmarshalJSON(b []byte) error {
	var TimeoffrequestreferenceMap map[string]interface{}
	err := json.Unmarshal(b, &TimeoffrequestreferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TimeoffrequestreferenceMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if SelfUri, ok := TimeoffrequestreferenceMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Timeoffrequestreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
