package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// V3sourcedetailedwitherrorresponse
type V3sourcedetailedwitherrorresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The name of the source.
	Name *string `json:"name,omitempty"`

	// ConnectionId - The connectionId of the source.
	ConnectionId *string `json:"connectionId,omitempty"`

	// VarType - The type of the source.
	VarType *string `json:"type,omitempty"`

	// TriggerType - The trigger type of the source.
	TriggerType *string `json:"triggerType,omitempty"`

	// Status - The current status of the source.
	Status *string `json:"status,omitempty"`

	// CreatedBy - The user who created the source.
	CreatedBy *Userreference `json:"createdBy,omitempty"`

	// ModifiedBy - The user who modified the document.
	ModifiedBy *Userreference `json:"modifiedBy,omitempty"`

	// DateCreated - Source creation date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - Source last modification date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// LastSync - The last synchronization of the source.
	LastSync *V3sourcelastsynchronization `json:"lastSync,omitempty"`

	// ScheduleSettings - Settings that determine when the source starts a sync.
	ScheduleSettings *V3sourceschedulesettings `json:"scheduleSettings,omitempty"`

	// Filters - Filters that determine what documents are synced.
	Filters *V3sourcefilter `json:"filters,omitempty"`

	// FilterDetails - Additional details to the source's filters.
	FilterDetails *V3sourcefilterdetails `json:"filterDetails,omitempty"`

	// VarError - Optional error details of an errored source.
	VarError *Errorbody `json:"error,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *V3sourcedetailedwitherrorresponse) SetField(field string, fieldValue interface{}) {
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

func (o V3sourcedetailedwitherrorresponse) MarshalJSON() ([]byte, error) {
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
	type Alias V3sourcedetailedwitherrorresponse
	
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
		
		ConnectionId *string `json:"connectionId,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		TriggerType *string `json:"triggerType,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		CreatedBy *Userreference `json:"createdBy,omitempty"`
		
		ModifiedBy *Userreference `json:"modifiedBy,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		LastSync *V3sourcelastsynchronization `json:"lastSync,omitempty"`
		
		ScheduleSettings *V3sourceschedulesettings `json:"scheduleSettings,omitempty"`
		
		Filters *V3sourcefilter `json:"filters,omitempty"`
		
		FilterDetails *V3sourcefilterdetails `json:"filterDetails,omitempty"`
		
		VarError *Errorbody `json:"error,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		ConnectionId: o.ConnectionId,
		
		VarType: o.VarType,
		
		TriggerType: o.TriggerType,
		
		Status: o.Status,
		
		CreatedBy: o.CreatedBy,
		
		ModifiedBy: o.ModifiedBy,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		LastSync: o.LastSync,
		
		ScheduleSettings: o.ScheduleSettings,
		
		Filters: o.Filters,
		
		FilterDetails: o.FilterDetails,
		
		VarError: o.VarError,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *V3sourcedetailedwitherrorresponse) UnmarshalJSON(b []byte) error {
	var V3sourcedetailedwitherrorresponseMap map[string]interface{}
	err := json.Unmarshal(b, &V3sourcedetailedwitherrorresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := V3sourcedetailedwitherrorresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := V3sourcedetailedwitherrorresponseMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ConnectionId, ok := V3sourcedetailedwitherrorresponseMap["connectionId"].(string); ok {
		o.ConnectionId = &ConnectionId
	}
    
	if VarType, ok := V3sourcedetailedwitherrorresponseMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if TriggerType, ok := V3sourcedetailedwitherrorresponseMap["triggerType"].(string); ok {
		o.TriggerType = &TriggerType
	}
    
	if Status, ok := V3sourcedetailedwitherrorresponseMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if CreatedBy, ok := V3sourcedetailedwitherrorresponseMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if ModifiedBy, ok := V3sourcedetailedwitherrorresponseMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if dateCreatedString, ok := V3sourcedetailedwitherrorresponseMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := V3sourcedetailedwitherrorresponseMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if LastSync, ok := V3sourcedetailedwitherrorresponseMap["lastSync"].(map[string]interface{}); ok {
		LastSyncString, _ := json.Marshal(LastSync)
		json.Unmarshal(LastSyncString, &o.LastSync)
	}
	
	if ScheduleSettings, ok := V3sourcedetailedwitherrorresponseMap["scheduleSettings"].(map[string]interface{}); ok {
		ScheduleSettingsString, _ := json.Marshal(ScheduleSettings)
		json.Unmarshal(ScheduleSettingsString, &o.ScheduleSettings)
	}
	
	if Filters, ok := V3sourcedetailedwitherrorresponseMap["filters"].(map[string]interface{}); ok {
		FiltersString, _ := json.Marshal(Filters)
		json.Unmarshal(FiltersString, &o.Filters)
	}
	
	if FilterDetails, ok := V3sourcedetailedwitherrorresponseMap["filterDetails"].(map[string]interface{}); ok {
		FilterDetailsString, _ := json.Marshal(FilterDetails)
		json.Unmarshal(FilterDetailsString, &o.FilterDetails)
	}
	
	if VarError, ok := V3sourcedetailedwitherrorresponseMap["error"].(map[string]interface{}); ok {
		VarErrorString, _ := json.Marshal(VarError)
		json.Unmarshal(VarErrorString, &o.VarError)
	}
	
	if SelfUri, ok := V3sourcedetailedwitherrorresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *V3sourcedetailedwitherrorresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
