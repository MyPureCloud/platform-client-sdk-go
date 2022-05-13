package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningmodulecoverartresponse - Learning module cover art response
type Learningmodulecoverartresponse struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`


	// Url - The URL for the cover art
	Url *string `json:"url,omitempty"`

}

func (o *Learningmodulecoverartresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningmodulecoverartresponse
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		Url *string `json:"url,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		SelfUri: o.SelfUri,
		
		Url: o.Url,
		Alias:    (*Alias)(o),
	})
}

func (o *Learningmodulecoverartresponse) UnmarshalJSON(b []byte) error {
	var LearningmodulecoverartresponseMap map[string]interface{}
	err := json.Unmarshal(b, &LearningmodulecoverartresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := LearningmodulecoverartresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if SelfUri, ok := LearningmodulecoverartresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if Url, ok := LearningmodulecoverartresponseMap["url"].(string); ok {
		o.Url = &Url
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Learningmodulecoverartresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
