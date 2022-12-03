package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Predictormodelbrief
type Predictormodelbrief struct { 
	// MediaType - The media type of the model.
	MediaType *string `json:"mediaType,omitempty"`


	// DateModified - The date the model was modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// RetrainingErrors - The model's retraining errors.
	RetrainingErrors *[]Predictormodelretrainingerror `json:"retrainingErrors,omitempty"`


	// State - The state of the model
	State *string `json:"state,omitempty"`

}

func (o *Predictormodelbrief) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Predictormodelbrief
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		MediaType *string `json:"mediaType,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		RetrainingErrors *[]Predictormodelretrainingerror `json:"retrainingErrors,omitempty"`
		
		State *string `json:"state,omitempty"`
		*Alias
	}{ 
		MediaType: o.MediaType,
		
		DateModified: DateModified,
		
		RetrainingErrors: o.RetrainingErrors,
		
		State: o.State,
		Alias:    (*Alias)(o),
	})
}

func (o *Predictormodelbrief) UnmarshalJSON(b []byte) error {
	var PredictormodelbriefMap map[string]interface{}
	err := json.Unmarshal(b, &PredictormodelbriefMap)
	if err != nil {
		return err
	}
	
	if MediaType, ok := PredictormodelbriefMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if dateModifiedString, ok := PredictormodelbriefMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if RetrainingErrors, ok := PredictormodelbriefMap["retrainingErrors"].([]interface{}); ok {
		RetrainingErrorsString, _ := json.Marshal(RetrainingErrors)
		json.Unmarshal(RetrainingErrorsString, &o.RetrainingErrors)
	}
	
	if State, ok := PredictormodelbriefMap["state"].(string); ok {
		o.State = &State
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Predictormodelbrief) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
