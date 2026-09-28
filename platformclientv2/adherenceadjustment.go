package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Adherenceadjustment
type Adherenceadjustment struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Agent - The agent to whom this adherence adjustment applies
	Agent *Userreference `json:"agent,omitempty"`

	// ManagementUnit - The management unit to which the agent belonged when the adherence adjustment was submitted
	ManagementUnit *Managementunitreference `json:"managementUnit,omitempty"`

	// BusinessUnit - The business unit to which the agent belonged when the adherence adjustment was submitted
	BusinessUnit *Businessunitreference `json:"businessUnit,omitempty"`

	// StartDate - The start timestamp of the adherence adjustment in ISO-8601 format
	StartDate *time.Time `json:"startDate,omitempty"`

	// LengthMinutes - The length of the adherence adjustment in minutes
	LengthMinutes *int `json:"lengthMinutes,omitempty"`

	// ReasonCode - The reason code for this adherence adjustment
	ReasonCode *Adherenceadjustmentsreasoncodereference `json:"reasonCode,omitempty"`

	// Status - The status of the adherence adjustment
	Status *string `json:"status,omitempty"`

	// Expired - Indicates if the adherence adjustment is expired
	Expired *bool `json:"expired,omitempty"`

	// SubmitterNotes - Notes provided by the submitter for this adherence adjustment
	SubmitterNotes *string `json:"submitterNotes,omitempty"`

	// ReviewerNotes - Notes provided by the reviewer for this adherence adjustment
	ReviewerNotes *string `json:"reviewerNotes,omitempty"`

	// ReviewedBy - The user who reviewed the adherence adjustment, if applicable. The id may be 'System' if it was an automated process
	ReviewedBy *Userreference `json:"reviewedBy,omitempty"`

	// ReviewedDate - The date the adherence adjustment was reviewed, if applicable. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ReviewedDate *time.Time `json:"reviewedDate,omitempty"`

	// Metadata - Version metadata for the adherence adjustment
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Adherenceadjustment) SetField(field string, fieldValue interface{}) {
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

func (o Adherenceadjustment) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "StartDate","ReviewedDate", }
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
	type Alias Adherenceadjustment
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	ReviewedDate := new(string)
	if o.ReviewedDate != nil {
		
		*ReviewedDate = timeutil.Strftime(o.ReviewedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ReviewedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Agent *Userreference `json:"agent,omitempty"`
		
		ManagementUnit *Managementunitreference `json:"managementUnit,omitempty"`
		
		BusinessUnit *Businessunitreference `json:"businessUnit,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		LengthMinutes *int `json:"lengthMinutes,omitempty"`
		
		ReasonCode *Adherenceadjustmentsreasoncodereference `json:"reasonCode,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Expired *bool `json:"expired,omitempty"`
		
		SubmitterNotes *string `json:"submitterNotes,omitempty"`
		
		ReviewerNotes *string `json:"reviewerNotes,omitempty"`
		
		ReviewedBy *Userreference `json:"reviewedBy,omitempty"`
		
		ReviewedDate *string `json:"reviewedDate,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Agent: o.Agent,
		
		ManagementUnit: o.ManagementUnit,
		
		BusinessUnit: o.BusinessUnit,
		
		StartDate: StartDate,
		
		LengthMinutes: o.LengthMinutes,
		
		ReasonCode: o.ReasonCode,
		
		Status: o.Status,
		
		Expired: o.Expired,
		
		SubmitterNotes: o.SubmitterNotes,
		
		ReviewerNotes: o.ReviewerNotes,
		
		ReviewedBy: o.ReviewedBy,
		
		ReviewedDate: ReviewedDate,
		
		Metadata: o.Metadata,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Adherenceadjustment) UnmarshalJSON(b []byte) error {
	var AdherenceadjustmentMap map[string]interface{}
	err := json.Unmarshal(b, &AdherenceadjustmentMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AdherenceadjustmentMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Agent, ok := AdherenceadjustmentMap["agent"].(map[string]interface{}); ok {
		AgentString, _ := json.Marshal(Agent)
		json.Unmarshal(AgentString, &o.Agent)
	}
	
	if ManagementUnit, ok := AdherenceadjustmentMap["managementUnit"].(map[string]interface{}); ok {
		ManagementUnitString, _ := json.Marshal(ManagementUnit)
		json.Unmarshal(ManagementUnitString, &o.ManagementUnit)
	}
	
	if BusinessUnit, ok := AdherenceadjustmentMap["businessUnit"].(map[string]interface{}); ok {
		BusinessUnitString, _ := json.Marshal(BusinessUnit)
		json.Unmarshal(BusinessUnitString, &o.BusinessUnit)
	}
	
	if startDateString, ok := AdherenceadjustmentMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if LengthMinutes, ok := AdherenceadjustmentMap["lengthMinutes"].(float64); ok {
		LengthMinutesInt := int(LengthMinutes)
		o.LengthMinutes = &LengthMinutesInt
	}
	
	if ReasonCode, ok := AdherenceadjustmentMap["reasonCode"].(map[string]interface{}); ok {
		ReasonCodeString, _ := json.Marshal(ReasonCode)
		json.Unmarshal(ReasonCodeString, &o.ReasonCode)
	}
	
	if Status, ok := AdherenceadjustmentMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Expired, ok := AdherenceadjustmentMap["expired"].(bool); ok {
		o.Expired = &Expired
	}
    
	if SubmitterNotes, ok := AdherenceadjustmentMap["submitterNotes"].(string); ok {
		o.SubmitterNotes = &SubmitterNotes
	}
    
	if ReviewerNotes, ok := AdherenceadjustmentMap["reviewerNotes"].(string); ok {
		o.ReviewerNotes = &ReviewerNotes
	}
    
	if ReviewedBy, ok := AdherenceadjustmentMap["reviewedBy"].(map[string]interface{}); ok {
		ReviewedByString, _ := json.Marshal(ReviewedBy)
		json.Unmarshal(ReviewedByString, &o.ReviewedBy)
	}
	
	if reviewedDateString, ok := AdherenceadjustmentMap["reviewedDate"].(string); ok {
		ReviewedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", reviewedDateString)
		o.ReviewedDate = &ReviewedDate
	}
	
	if Metadata, ok := AdherenceadjustmentMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if SelfUri, ok := AdherenceadjustmentMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Adherenceadjustment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
