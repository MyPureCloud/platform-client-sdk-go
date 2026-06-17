package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Agenteffectivebid
type Agenteffectivebid struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The ID of the schedule bid
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// EffectiveDate - The effective date of the bid relative to the business unit time zone in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	EffectiveDate *time.Time `json:"effectiveDate,omitempty"`

	// EndDate - The end date of the bid, relative to the business unit time zone in yyyy-MM-dd format. Null denotes an active schedule bid. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	EndDate *time.Time `json:"endDate,omitempty"`

	// DownloadUrl - The download URL to fetch the list of schedule sets and the agents assigned to them
	DownloadUrl *string `json:"downloadUrl,omitempty"`

	// DownloadTemplate - This field will always be null. Effective schedule sets are returned through the download URL. The schema is included here for documentation purposes
	DownloadTemplate *Agentassignedschedulesetlist `json:"downloadTemplate,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Agenteffectivebid) SetField(field string, fieldValue interface{}) {
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

func (o Agenteffectivebid) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{ "EffectiveDate","EndDate", }

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
	type Alias Agenteffectivebid
	
	EffectiveDate := new(string)
	if o.EffectiveDate != nil {
		*EffectiveDate = timeutil.Strftime(o.EffectiveDate, "%Y-%m-%d")
	} else {
		EffectiveDate = nil
	}
	
	EndDate := new(string)
	if o.EndDate != nil {
		*EndDate = timeutil.Strftime(o.EndDate, "%Y-%m-%d")
	} else {
		EndDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		EffectiveDate *string `json:"effectiveDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		
		DownloadUrl *string `json:"downloadUrl,omitempty"`
		
		DownloadTemplate *Agentassignedschedulesetlist `json:"downloadTemplate,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		EffectiveDate: EffectiveDate,
		
		EndDate: EndDate,
		
		DownloadUrl: o.DownloadUrl,
		
		DownloadTemplate: o.DownloadTemplate,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Agenteffectivebid) UnmarshalJSON(b []byte) error {
	var AgenteffectivebidMap map[string]interface{}
	err := json.Unmarshal(b, &AgenteffectivebidMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AgenteffectivebidMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := AgenteffectivebidMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if effectiveDateString, ok := AgenteffectivebidMap["effectiveDate"].(string); ok {
		EffectiveDate, _ := time.Parse("2006-01-02", effectiveDateString)
		o.EffectiveDate = &EffectiveDate
	}
	
	if endDateString, ok := AgenteffectivebidMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02", endDateString)
		o.EndDate = &EndDate
	}
	
	if DownloadUrl, ok := AgenteffectivebidMap["downloadUrl"].(string); ok {
		o.DownloadUrl = &DownloadUrl
	}
    
	if DownloadTemplate, ok := AgenteffectivebidMap["downloadTemplate"].(map[string]interface{}); ok {
		DownloadTemplateString, _ := json.Marshal(DownloadTemplate)
		json.Unmarshal(DownloadTemplateString, &o.DownloadTemplate)
	}
	
	if SelfUri, ok := AgenteffectivebidMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Agenteffectivebid) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
