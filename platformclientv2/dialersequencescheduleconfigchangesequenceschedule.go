package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialersequencescheduleconfigchangesequenceschedule
type Dialersequencescheduleconfigchangesequenceschedule struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DateCreated
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version
	Version *int `json:"version,omitempty"`


	// Intervals
	Intervals *[]Dialersequencescheduleconfigchangescheduleinterval `json:"intervals,omitempty"`


	// TimeZone
	TimeZone *string `json:"timeZone,omitempty"`


	// Sequence
	Sequence *Dialersequencescheduleconfigchangeurireference `json:"sequence,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (o *Dialersequencescheduleconfigchangesequenceschedule) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialersequencescheduleconfigchangesequenceschedule
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		Intervals *[]Dialersequencescheduleconfigchangescheduleinterval `json:"intervals,omitempty"`
		
		TimeZone *string `json:"timeZone,omitempty"`
		
		Sequence *Dialersequencescheduleconfigchangeurireference `json:"sequence,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		
		Intervals: o.Intervals,
		
		TimeZone: o.TimeZone,
		
		Sequence: o.Sequence,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialersequencescheduleconfigchangesequenceschedule) UnmarshalJSON(b []byte) error {
	var DialersequencescheduleconfigchangesequencescheduleMap map[string]interface{}
	err := json.Unmarshal(b, &DialersequencescheduleconfigchangesequencescheduleMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DialersequencescheduleconfigchangesequencescheduleMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := DialersequencescheduleconfigchangesequencescheduleMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if dateCreatedString, ok := DialersequencescheduleconfigchangesequencescheduleMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := DialersequencescheduleconfigchangesequencescheduleMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := DialersequencescheduleconfigchangesequencescheduleMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if Intervals, ok := DialersequencescheduleconfigchangesequencescheduleMap["intervals"].([]interface{}); ok {
		IntervalsString, _ := json.Marshal(Intervals)
		json.Unmarshal(IntervalsString, &o.Intervals)
	}
	
	if TimeZone, ok := DialersequencescheduleconfigchangesequencescheduleMap["timeZone"].(string); ok {
		o.TimeZone = &TimeZone
	}
	
	if Sequence, ok := DialersequencescheduleconfigchangesequencescheduleMap["sequence"].(map[string]interface{}); ok {
		SequenceString, _ := json.Marshal(Sequence)
		json.Unmarshal(SequenceString, &o.Sequence)
	}
	
	if AdditionalProperties, ok := DialersequencescheduleconfigchangesequencescheduleMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialersequencescheduleconfigchangesequenceschedule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
