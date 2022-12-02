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
	// Intervals - a list of start and end times
	Intervals *[]Dialersequencescheduleconfigchangescheduleinterval `json:"intervals,omitempty"`


	// TimeZone - time zone identifier to be applied to the intervals; for example Africa/Abidjan
	TimeZone *string `json:"timeZone,omitempty"`


	// Sequence
	Sequence *Dialersequencescheduleconfigchangeurireference `json:"sequence,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`


	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The UI-visible name of the object
	Name *string `json:"name,omitempty"`


	// DateCreated - Creation time of the entity
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Last modified time of the entity
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version - Required for updates, must match the version number of the most recent update
	Version *int `json:"version,omitempty"`

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
		Intervals *[]Dialersequencescheduleconfigchangescheduleinterval `json:"intervals,omitempty"`
		
		TimeZone *string `json:"timeZone,omitempty"`
		
		Sequence *Dialersequencescheduleconfigchangeurireference `json:"sequence,omitempty"`
		
		AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Version *int `json:"version,omitempty"`
		*Alias
	}{ 
		Intervals: o.Intervals,
		
		TimeZone: o.TimeZone,
		
		Sequence: o.Sequence,
		
		AdditionalProperties: o.AdditionalProperties,
		
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialersequencescheduleconfigchangesequenceschedule) UnmarshalJSON(b []byte) error {
	var DialersequencescheduleconfigchangesequencescheduleMap map[string]interface{}
	err := json.Unmarshal(b, &DialersequencescheduleconfigchangesequencescheduleMap)
	if err != nil {
		return err
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
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialersequencescheduleconfigchangesequenceschedule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
