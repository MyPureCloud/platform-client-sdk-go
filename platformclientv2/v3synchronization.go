package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// V3synchronization
type V3synchronization struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// VarType - The type of the synchronization.
	VarType *string `json:"type,omitempty"`

	// CreatedBy - The user who started the synchronization if the source is manually synchronized or the user who created the source for scheduled synchronization.
	CreatedBy *Userreference `json:"createdBy,omitempty"`

	// Source - The source of the synchronization.
	Source *V3sourceref `json:"source,omitempty"`

	// DateStart - The start time of the synchronization. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStart *time.Time `json:"dateStart,omitempty"`

	// DateEnd - The end time of the synchronization. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateEnd *time.Time `json:"dateEnd,omitempty"`

	// DateSourceIntervalStart - The start time of the interval to be synchronized from the source. Source item changes during that interval are included in this synchronization. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateSourceIntervalStart *time.Time `json:"dateSourceIntervalStart,omitempty"`

	// DateSourceIntervalEnd - The end time of the interval to be synchronized from the source. Source item changes during that interval are included in this synchronization. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateSourceIntervalEnd *time.Time `json:"dateSourceIntervalEnd,omitempty"`

	// TriggerType - The trigger type of the synchronization.
	TriggerType *string `json:"triggerType,omitempty"`

	// Status - The status of the synchronization.
	Status *string `json:"status,omitempty"`

	// Statistics - Statistics of the synchronization.
	Statistics *V3synchronizationstatistics `json:"statistics,omitempty"`

	// VarError - The error that occurred during the synchronization.
	VarError *Errorbody `json:"error,omitempty"`

	// IngestionStatus - The status of the ingestion.
	IngestionStatus *string `json:"ingestionStatus,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *V3synchronization) SetField(field string, fieldValue interface{}) {
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

func (o V3synchronization) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateStart","DateEnd","DateSourceIntervalStart","DateSourceIntervalEnd", }
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
	type Alias V3synchronization
	
	DateStart := new(string)
	if o.DateStart != nil {
		
		*DateStart = timeutil.Strftime(o.DateStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStart = nil
	}
	
	DateEnd := new(string)
	if o.DateEnd != nil {
		
		*DateEnd = timeutil.Strftime(o.DateEnd, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateEnd = nil
	}
	
	DateSourceIntervalStart := new(string)
	if o.DateSourceIntervalStart != nil {
		
		*DateSourceIntervalStart = timeutil.Strftime(o.DateSourceIntervalStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateSourceIntervalStart = nil
	}
	
	DateSourceIntervalEnd := new(string)
	if o.DateSourceIntervalEnd != nil {
		
		*DateSourceIntervalEnd = timeutil.Strftime(o.DateSourceIntervalEnd, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateSourceIntervalEnd = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		CreatedBy *Userreference `json:"createdBy,omitempty"`
		
		Source *V3sourceref `json:"source,omitempty"`
		
		DateStart *string `json:"dateStart,omitempty"`
		
		DateEnd *string `json:"dateEnd,omitempty"`
		
		DateSourceIntervalStart *string `json:"dateSourceIntervalStart,omitempty"`
		
		DateSourceIntervalEnd *string `json:"dateSourceIntervalEnd,omitempty"`
		
		TriggerType *string `json:"triggerType,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Statistics *V3synchronizationstatistics `json:"statistics,omitempty"`
		
		VarError *Errorbody `json:"error,omitempty"`
		
		IngestionStatus *string `json:"ingestionStatus,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		VarType: o.VarType,
		
		CreatedBy: o.CreatedBy,
		
		Source: o.Source,
		
		DateStart: DateStart,
		
		DateEnd: DateEnd,
		
		DateSourceIntervalStart: DateSourceIntervalStart,
		
		DateSourceIntervalEnd: DateSourceIntervalEnd,
		
		TriggerType: o.TriggerType,
		
		Status: o.Status,
		
		Statistics: o.Statistics,
		
		VarError: o.VarError,
		
		IngestionStatus: o.IngestionStatus,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *V3synchronization) UnmarshalJSON(b []byte) error {
	var V3synchronizationMap map[string]interface{}
	err := json.Unmarshal(b, &V3synchronizationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := V3synchronizationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if VarType, ok := V3synchronizationMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if CreatedBy, ok := V3synchronizationMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if Source, ok := V3synchronizationMap["source"].(map[string]interface{}); ok {
		SourceString, _ := json.Marshal(Source)
		json.Unmarshal(SourceString, &o.Source)
	}
	
	if dateStartString, ok := V3synchronizationMap["dateStart"].(string); ok {
		DateStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStartString)
		o.DateStart = &DateStart
	}
	
	if dateEndString, ok := V3synchronizationMap["dateEnd"].(string); ok {
		DateEnd, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateEndString)
		o.DateEnd = &DateEnd
	}
	
	if dateSourceIntervalStartString, ok := V3synchronizationMap["dateSourceIntervalStart"].(string); ok {
		DateSourceIntervalStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateSourceIntervalStartString)
		o.DateSourceIntervalStart = &DateSourceIntervalStart
	}
	
	if dateSourceIntervalEndString, ok := V3synchronizationMap["dateSourceIntervalEnd"].(string); ok {
		DateSourceIntervalEnd, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateSourceIntervalEndString)
		o.DateSourceIntervalEnd = &DateSourceIntervalEnd
	}
	
	if TriggerType, ok := V3synchronizationMap["triggerType"].(string); ok {
		o.TriggerType = &TriggerType
	}
    
	if Status, ok := V3synchronizationMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Statistics, ok := V3synchronizationMap["statistics"].(map[string]interface{}); ok {
		StatisticsString, _ := json.Marshal(Statistics)
		json.Unmarshal(StatisticsString, &o.Statistics)
	}
	
	if VarError, ok := V3synchronizationMap["error"].(map[string]interface{}); ok {
		VarErrorString, _ := json.Marshal(VarError)
		json.Unmarshal(VarErrorString, &o.VarError)
	}
	
	if IngestionStatus, ok := V3synchronizationMap["ingestionStatus"].(string); ok {
		o.IngestionStatus = &IngestionStatus
	}
    
	if SelfUri, ok := V3synchronizationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *V3synchronization) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
