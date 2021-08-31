package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercallabletimesetconfigchangecallabletimeset
type Dialercallabletimesetconfigchangecallabletimeset struct { 
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


	// CallableTimes
	CallableTimes *[]Dialercallabletimesetconfigchangecallabletime `json:"callableTimes,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (o *Dialercallabletimesetconfigchangecallabletimeset) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercallabletimesetconfigchangecallabletimeset
	
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
		
		CallableTimes *[]Dialercallabletimesetconfigchangecallabletime `json:"callableTimes,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		
		CallableTimes: o.CallableTimes,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialercallabletimesetconfigchangecallabletimeset) UnmarshalJSON(b []byte) error {
	var DialercallabletimesetconfigchangecallabletimesetMap map[string]interface{}
	err := json.Unmarshal(b, &DialercallabletimesetconfigchangecallabletimesetMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DialercallabletimesetconfigchangecallabletimesetMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := DialercallabletimesetconfigchangecallabletimesetMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if dateCreatedString, ok := DialercallabletimesetconfigchangecallabletimesetMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := DialercallabletimesetconfigchangecallabletimesetMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := DialercallabletimesetconfigchangecallabletimesetMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if CallableTimes, ok := DialercallabletimesetconfigchangecallabletimesetMap["callableTimes"].([]interface{}); ok {
		CallableTimesString, _ := json.Marshal(CallableTimes)
		json.Unmarshal(CallableTimesString, &o.CallableTimes)
	}
	
	if AdditionalProperties, ok := DialercallabletimesetconfigchangecallabletimesetMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercallabletimesetconfigchangecallabletimeset) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
