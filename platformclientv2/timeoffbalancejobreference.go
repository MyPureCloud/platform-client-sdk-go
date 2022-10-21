package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Timeoffbalancejobreference
type Timeoffbalancejobreference struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Status - The status of the job
	Status *string `json:"status,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Timeoffbalancejobreference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Timeoffbalancejobreference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Status: o.Status,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Timeoffbalancejobreference) UnmarshalJSON(b []byte) error {
	var TimeoffbalancejobreferenceMap map[string]interface{}
	err := json.Unmarshal(b, &TimeoffbalancejobreferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TimeoffbalancejobreferenceMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Status, ok := TimeoffbalancejobreferenceMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if SelfUri, ok := TimeoffbalancejobreferenceMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Timeoffbalancejobreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
