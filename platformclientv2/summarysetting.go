package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Summarysetting
type Summarysetting struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - Name of the summary setting.
	Name *string `json:"name,omitempty"`

	// Language - Language of the generated summary, e.g. en-US, it-IT.
	Language *string `json:"language,omitempty"`

	// SummaryType - Level of detail of the generated summary.
	SummaryType *string `json:"summaryType,omitempty"`

	// Format - Format of the generated summary.
	Format *string `json:"format,omitempty"`

	// MaskPII - Displaying PII in the generated summary.
	MaskPII *Summarysettingpii `json:"maskPII,omitempty"`

	// ParticipantLabels - How to refer to interaction participants in the generated summary.
	ParticipantLabels *Summarysettingparticipantlabels `json:"participantLabels,omitempty"`

	// PredefinedInsights - Set which insights to include in the generated summary by default.
	PredefinedInsights *[]string `json:"predefinedInsights,omitempty"`

	// CustomEntities - Custom entity definition.
	CustomEntities *[]Summarysettingcustomentity `json:"customEntities,omitempty"`

	// SettingType - Type of the summary setting.
	SettingType *string `json:"settingType,omitempty"`

	// Prompt - Custom prompt of summary setting.
	Prompt *string `json:"prompt,omitempty"`

	// ServiceType - Service type for summarization. Can be 'Native' for Genesys native summarization engine or 'External' for external service. If specified as 'External', integrationId must be provided.
	ServiceType *string `json:"serviceType,omitempty"`

	// IntegrationId - Integration ID for the external summarization service. Required when serviceType is External.
	IntegrationId *string `json:"integrationId,omitempty"`

	// TimeoutDuration - Timeout duration in seconds for the external summarization service request.
	TimeoutDuration *int `json:"timeoutDuration,omitempty"`

	// DateCreated - The date and time the setting was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - The date and time the setting was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Summarysetting) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Summarysetting) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateModified", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Summarysetting
	
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
		
		Language *string `json:"language,omitempty"`
		
		SummaryType *string `json:"summaryType,omitempty"`
		
		Format *string `json:"format,omitempty"`
		
		MaskPII *Summarysettingpii `json:"maskPII,omitempty"`
		
		ParticipantLabels *Summarysettingparticipantlabels `json:"participantLabels,omitempty"`
		
		PredefinedInsights *[]string `json:"predefinedInsights,omitempty"`
		
		CustomEntities *[]Summarysettingcustomentity `json:"customEntities,omitempty"`
		
		SettingType *string `json:"settingType,omitempty"`
		
		Prompt *string `json:"prompt,omitempty"`
		
		ServiceType *string `json:"serviceType,omitempty"`
		
		IntegrationId *string `json:"integrationId,omitempty"`
		
		TimeoutDuration *int `json:"timeoutDuration,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Language: o.Language,
		
		SummaryType: o.SummaryType,
		
		Format: o.Format,
		
		MaskPII: o.MaskPII,
		
		ParticipantLabels: o.ParticipantLabels,
		
		PredefinedInsights: o.PredefinedInsights,
		
		CustomEntities: o.CustomEntities,
		
		SettingType: o.SettingType,
		
		Prompt: o.Prompt,
		
		ServiceType: o.ServiceType,
		
		IntegrationId: o.IntegrationId,
		
		TimeoutDuration: o.TimeoutDuration,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Summarysetting) UnmarshalJSON(b []byte) error {
	var SummarysettingMap map[string]interface{}
	err := json.Unmarshal(b, &SummarysettingMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SummarysettingMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := SummarysettingMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Language, ok := SummarysettingMap["language"].(string); ok {
		o.Language = &Language
	}
    
	if SummaryType, ok := SummarysettingMap["summaryType"].(string); ok {
		o.SummaryType = &SummaryType
	}
    
	if Format, ok := SummarysettingMap["format"].(string); ok {
		o.Format = &Format
	}
    
	if MaskPII, ok := SummarysettingMap["maskPII"].(map[string]interface{}); ok {
		MaskPIIString, _ := json.Marshal(MaskPII)
		json.Unmarshal(MaskPIIString, &o.MaskPII)
	}
	
	if ParticipantLabels, ok := SummarysettingMap["participantLabels"].(map[string]interface{}); ok {
		ParticipantLabelsString, _ := json.Marshal(ParticipantLabels)
		json.Unmarshal(ParticipantLabelsString, &o.ParticipantLabels)
	}
	
	if PredefinedInsights, ok := SummarysettingMap["predefinedInsights"].([]interface{}); ok {
		PredefinedInsightsString, _ := json.Marshal(PredefinedInsights)
		json.Unmarshal(PredefinedInsightsString, &o.PredefinedInsights)
	}
	
	if CustomEntities, ok := SummarysettingMap["customEntities"].([]interface{}); ok {
		CustomEntitiesString, _ := json.Marshal(CustomEntities)
		json.Unmarshal(CustomEntitiesString, &o.CustomEntities)
	}
	
	if SettingType, ok := SummarysettingMap["settingType"].(string); ok {
		o.SettingType = &SettingType
	}
    
	if Prompt, ok := SummarysettingMap["prompt"].(string); ok {
		o.Prompt = &Prompt
	}
    
	if ServiceType, ok := SummarysettingMap["serviceType"].(string); ok {
		o.ServiceType = &ServiceType
	}
    
	if IntegrationId, ok := SummarysettingMap["integrationId"].(string); ok {
		o.IntegrationId = &IntegrationId
	}
    
	if TimeoutDuration, ok := SummarysettingMap["timeoutDuration"].(float64); ok {
		TimeoutDurationInt := int(TimeoutDuration)
		o.TimeoutDuration = &TimeoutDurationInt
	}
	
	if dateCreatedString, ok := SummarysettingMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := SummarysettingMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if SelfUri, ok := SummarysettingMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Summarysetting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
