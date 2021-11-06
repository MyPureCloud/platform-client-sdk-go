package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Sentimentfeedback
type Sentimentfeedback struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Phrase - The phrase for which sentiment feedback is provided
	Phrase *string `json:"phrase,omitempty"`


	// Dialect - The dialect for the given phrase, dialect format is {language}-{country} where language follows ISO 639-1 standard and country follows ISO 3166-1 alpha 2 standard
	Dialect *string `json:"dialect,omitempty"`


	// FeedbackValue - The sentiment feedback value for the given phrase
	FeedbackValue *string `json:"feedbackValue,omitempty"`


	// DateCreated - The Timestamp when sentiment feedback created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// CreatedBy - The Id of user who created the sentiment feedback
	CreatedBy *Addressableentityref `json:"createdBy,omitempty"`

}

func (o *Sentimentfeedback) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Sentimentfeedback
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Phrase *string `json:"phrase,omitempty"`
		
		Dialect *string `json:"dialect,omitempty"`
		
		FeedbackValue *string `json:"feedbackValue,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		CreatedBy *Addressableentityref `json:"createdBy,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Phrase: o.Phrase,
		
		Dialect: o.Dialect,
		
		FeedbackValue: o.FeedbackValue,
		
		DateCreated: DateCreated,
		
		CreatedBy: o.CreatedBy,
		Alias:    (*Alias)(o),
	})
}

func (o *Sentimentfeedback) UnmarshalJSON(b []byte) error {
	var SentimentfeedbackMap map[string]interface{}
	err := json.Unmarshal(b, &SentimentfeedbackMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SentimentfeedbackMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Phrase, ok := SentimentfeedbackMap["phrase"].(string); ok {
		o.Phrase = &Phrase
	}
	
	if Dialect, ok := SentimentfeedbackMap["dialect"].(string); ok {
		o.Dialect = &Dialect
	}
	
	if FeedbackValue, ok := SentimentfeedbackMap["feedbackValue"].(string); ok {
		o.FeedbackValue = &FeedbackValue
	}
	
	if dateCreatedString, ok := SentimentfeedbackMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if CreatedBy, ok := SentimentfeedbackMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Sentimentfeedback) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
