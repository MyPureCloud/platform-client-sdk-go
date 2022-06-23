package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningmodulejobresponse - Learning module job response
type Learningmodulejobresponse struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Status - The status of learning module job
	Status *string `json:"status,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Learningmodulejobresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningmodulejobresponse
	
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

func (o *Learningmodulejobresponse) UnmarshalJSON(b []byte) error {
	var LearningmodulejobresponseMap map[string]interface{}
	err := json.Unmarshal(b, &LearningmodulejobresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := LearningmodulejobresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Status, ok := LearningmodulejobresponseMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if SelfUri, ok := LearningmodulejobresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Learningmodulejobresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
