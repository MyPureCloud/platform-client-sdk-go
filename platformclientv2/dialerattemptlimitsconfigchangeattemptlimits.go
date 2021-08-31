package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialerattemptlimitsconfigchangeattemptlimits
type Dialerattemptlimitsconfigchangeattemptlimits struct { 
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


	// MaxAttemptsPerContact
	MaxAttemptsPerContact *int `json:"maxAttemptsPerContact,omitempty"`


	// MaxAttemptsPerNumber
	MaxAttemptsPerNumber *int `json:"maxAttemptsPerNumber,omitempty"`


	// TimeZoneId
	TimeZoneId *string `json:"timeZoneId,omitempty"`


	// ResetPeriod
	ResetPeriod *string `json:"resetPeriod,omitempty"`


	// RecallEntries
	RecallEntries *map[string]Dialerattemptlimitsconfigchangerecallentry `json:"recallEntries,omitempty"`


	// BreadthFirstRecalls
	BreadthFirstRecalls *bool `json:"breadthFirstRecalls,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (o *Dialerattemptlimitsconfigchangeattemptlimits) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialerattemptlimitsconfigchangeattemptlimits
	
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
		
		MaxAttemptsPerContact *int `json:"maxAttemptsPerContact,omitempty"`
		
		MaxAttemptsPerNumber *int `json:"maxAttemptsPerNumber,omitempty"`
		
		TimeZoneId *string `json:"timeZoneId,omitempty"`
		
		ResetPeriod *string `json:"resetPeriod,omitempty"`
		
		RecallEntries *map[string]Dialerattemptlimitsconfigchangerecallentry `json:"recallEntries,omitempty"`
		
		BreadthFirstRecalls *bool `json:"breadthFirstRecalls,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		
		MaxAttemptsPerContact: o.MaxAttemptsPerContact,
		
		MaxAttemptsPerNumber: o.MaxAttemptsPerNumber,
		
		TimeZoneId: o.TimeZoneId,
		
		ResetPeriod: o.ResetPeriod,
		
		RecallEntries: o.RecallEntries,
		
		BreadthFirstRecalls: o.BreadthFirstRecalls,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialerattemptlimitsconfigchangeattemptlimits) UnmarshalJSON(b []byte) error {
	var DialerattemptlimitsconfigchangeattemptlimitsMap map[string]interface{}
	err := json.Unmarshal(b, &DialerattemptlimitsconfigchangeattemptlimitsMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DialerattemptlimitsconfigchangeattemptlimitsMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := DialerattemptlimitsconfigchangeattemptlimitsMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if dateCreatedString, ok := DialerattemptlimitsconfigchangeattemptlimitsMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := DialerattemptlimitsconfigchangeattemptlimitsMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := DialerattemptlimitsconfigchangeattemptlimitsMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if MaxAttemptsPerContact, ok := DialerattemptlimitsconfigchangeattemptlimitsMap["maxAttemptsPerContact"].(float64); ok {
		MaxAttemptsPerContactInt := int(MaxAttemptsPerContact)
		o.MaxAttemptsPerContact = &MaxAttemptsPerContactInt
	}
	
	if MaxAttemptsPerNumber, ok := DialerattemptlimitsconfigchangeattemptlimitsMap["maxAttemptsPerNumber"].(float64); ok {
		MaxAttemptsPerNumberInt := int(MaxAttemptsPerNumber)
		o.MaxAttemptsPerNumber = &MaxAttemptsPerNumberInt
	}
	
	if TimeZoneId, ok := DialerattemptlimitsconfigchangeattemptlimitsMap["timeZoneId"].(string); ok {
		o.TimeZoneId = &TimeZoneId
	}
	
	if ResetPeriod, ok := DialerattemptlimitsconfigchangeattemptlimitsMap["resetPeriod"].(string); ok {
		o.ResetPeriod = &ResetPeriod
	}
	
	if RecallEntries, ok := DialerattemptlimitsconfigchangeattemptlimitsMap["recallEntries"].(map[string]interface{}); ok {
		RecallEntriesString, _ := json.Marshal(RecallEntries)
		json.Unmarshal(RecallEntriesString, &o.RecallEntries)
	}
	
	if BreadthFirstRecalls, ok := DialerattemptlimitsconfigchangeattemptlimitsMap["breadthFirstRecalls"].(bool); ok {
		o.BreadthFirstRecalls = &BreadthFirstRecalls
	}
	
	if AdditionalProperties, ok := DialerattemptlimitsconfigchangeattemptlimitsMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialerattemptlimitsconfigchangeattemptlimits) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
