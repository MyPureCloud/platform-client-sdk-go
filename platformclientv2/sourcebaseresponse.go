package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Sourcebaseresponse
type Sourcebaseresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - Name of the source.
	Name *string `json:"name,omitempty"`

	// DateCreated - Source creation date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - Source last modification date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// VarType - The source type.
	VarType *string `json:"type,omitempty"`

	// Integration - The reference to the integration associated with the source.
	Integration *Knowledgeintegrationreference `json:"integration,omitempty"`

	// SchedulePeriod - The schedule period of the source in hours.
	SchedulePeriod *int `json:"schedulePeriod,omitempty"`

	// LastSync - Additional information about the last synchronization of the source.
	LastSync *Sourcelastsync `json:"lastSync,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Sourcebaseresponse) SetField(field string, fieldValue interface{}) {
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

func (o Sourcebaseresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Sourcebaseresponse
	
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
		
		VarType *string `json:"type,omitempty"`
		
		Integration *Knowledgeintegrationreference `json:"integration,omitempty"`
		
		SchedulePeriod *int `json:"schedulePeriod,omitempty"`
		
		LastSync *Sourcelastsync `json:"lastSync,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		VarType: o.VarType,
		
		Integration: o.Integration,
		
		SchedulePeriod: o.SchedulePeriod,
		
		LastSync: o.LastSync,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Sourcebaseresponse) UnmarshalJSON(b []byte) error {
	var SourcebaseresponseMap map[string]interface{}
	err := json.Unmarshal(b, &SourcebaseresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SourcebaseresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := SourcebaseresponseMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if dateCreatedString, ok := SourcebaseresponseMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := SourcebaseresponseMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if VarType, ok := SourcebaseresponseMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Integration, ok := SourcebaseresponseMap["integration"].(map[string]interface{}); ok {
		IntegrationString, _ := json.Marshal(Integration)
		json.Unmarshal(IntegrationString, &o.Integration)
	}
	
	if SchedulePeriod, ok := SourcebaseresponseMap["schedulePeriod"].(float64); ok {
		SchedulePeriodInt := int(SchedulePeriod)
		o.SchedulePeriod = &SchedulePeriodInt
	}
	
	if LastSync, ok := SourcebaseresponseMap["lastSync"].(map[string]interface{}); ok {
		LastSyncString, _ := json.Marshal(LastSync)
		json.Unmarshal(LastSyncString, &o.LastSync)
	}
	
	if SelfUri, ok := SourcebaseresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Sourcebaseresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
