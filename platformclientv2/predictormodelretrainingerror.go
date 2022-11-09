package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Predictormodelretrainingerror
type Predictormodelretrainingerror struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// ErrorCode - Error code describing model training failure.
	ErrorCode *string `json:"errorCode,omitempty"`


	// DateOfFirstOccurrence - Date when the first retraining failure happened. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateOfFirstOccurrence *time.Time `json:"dateOfFirstOccurrence,omitempty"`

}

func (o *Predictormodelretrainingerror) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Predictormodelretrainingerror
	
	DateOfFirstOccurrence := new(string)
	if o.DateOfFirstOccurrence != nil {
		
		*DateOfFirstOccurrence = timeutil.Strftime(o.DateOfFirstOccurrence, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateOfFirstOccurrence = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ErrorCode *string `json:"errorCode,omitempty"`
		
		DateOfFirstOccurrence *string `json:"dateOfFirstOccurrence,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		ErrorCode: o.ErrorCode,
		
		DateOfFirstOccurrence: DateOfFirstOccurrence,
		Alias:    (*Alias)(o),
	})
}

func (o *Predictormodelretrainingerror) UnmarshalJSON(b []byte) error {
	var PredictormodelretrainingerrorMap map[string]interface{}
	err := json.Unmarshal(b, &PredictormodelretrainingerrorMap)
	if err != nil {
		return err
	}
	
	if Id, ok := PredictormodelretrainingerrorMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if ErrorCode, ok := PredictormodelretrainingerrorMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
    
	if dateOfFirstOccurrenceString, ok := PredictormodelretrainingerrorMap["dateOfFirstOccurrence"].(string); ok {
		DateOfFirstOccurrence, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateOfFirstOccurrenceString)
		o.DateOfFirstOccurrence = &DateOfFirstOccurrence
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Predictormodelretrainingerror) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
