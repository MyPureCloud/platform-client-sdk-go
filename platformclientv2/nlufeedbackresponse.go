package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Nlufeedbackresponse
type Nlufeedbackresponse struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Text - The feedback text.
	Text *string `json:"text,omitempty"`


	// Intents - Detected intent of the utterance
	Intents *[]Intentfeedback `json:"intents,omitempty"`


	// Version - The domain version of the feedback.
	Version *Nludomainversion `json:"version,omitempty"`


	// DateCreated - The date when the feedback was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// Language - The language of the version to which feedback is linked, e.g. en-us, de-de
	Language *string `json:"language,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Nlufeedbackresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Nlufeedbackresponse
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		Intents *[]Intentfeedback `json:"intents,omitempty"`
		
		Version *Nludomainversion `json:"version,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Text: o.Text,
		
		Intents: o.Intents,
		
		Version: o.Version,
		
		DateCreated: DateCreated,
		
		Language: o.Language,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Nlufeedbackresponse) UnmarshalJSON(b []byte) error {
	var NlufeedbackresponseMap map[string]interface{}
	err := json.Unmarshal(b, &NlufeedbackresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := NlufeedbackresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Text, ok := NlufeedbackresponseMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Intents, ok := NlufeedbackresponseMap["intents"].([]interface{}); ok {
		IntentsString, _ := json.Marshal(Intents)
		json.Unmarshal(IntentsString, &o.Intents)
	}
	
	if Version, ok := NlufeedbackresponseMap["version"].(map[string]interface{}); ok {
		VersionString, _ := json.Marshal(Version)
		json.Unmarshal(VersionString, &o.Version)
	}
	
	if dateCreatedString, ok := NlufeedbackresponseMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if Language, ok := NlufeedbackresponseMap["language"].(string); ok {
		o.Language = &Language
	}
    
	if SelfUri, ok := NlufeedbackresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Nlufeedbackresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
