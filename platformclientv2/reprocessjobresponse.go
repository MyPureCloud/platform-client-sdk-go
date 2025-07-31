package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Reprocessjobresponse
type Reprocessjobresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Description - The description of the job.
	Description *string `json:"description,omitempty"`

	// DateStart - The date from which the reprocessing should begin. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStart *time.Time `json:"dateStart,omitempty"`

	// DateEnd - The date at which the reprocessing should end. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateEnd *time.Time `json:"dateEnd,omitempty"`

	// MediaTypes - Media types used to filter interactions.
	MediaTypes *[]string `json:"mediaTypes,omitempty"`

	// Programs - The mapped programs to be included in the job.
	Programs *[]string `json:"programs,omitempty"`

	// Dialects - Language dialects used to filter interactions.
	Dialects *[]string `json:"dialects,omitempty"`

	// CreatedBy - The user who created the job.
	CreatedBy *Addressableentityref `json:"createdBy,omitempty"`

	// DateCreated - The date the job was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// JobStatus - The status of the job.
	JobStatus *string `json:"jobStatus,omitempty"`

	// QueueOrder - The position of the job in the queued jobs.
	QueueOrder *int `json:"queueOrder,omitempty"`

	// ProcessedInteractionsCount - The amount of interactions successfully reprocessed.
	ProcessedInteractionsCount *int `json:"processedInteractionsCount,omitempty"`

	// FailedInteractionsCount - The amount of failed interactions.
	FailedInteractionsCount *int `json:"failedInteractionsCount,omitempty"`

	// TotalInteractionsCount - The amount of interactions in the job.
	TotalInteractionsCount *int `json:"totalInteractionsCount,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Reprocessjobresponse) SetField(field string, fieldValue interface{}) {
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

func (o Reprocessjobresponse) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateStart","DateEnd","DateCreated", }
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
	type Alias Reprocessjobresponse
	
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
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		DateStart *string `json:"dateStart,omitempty"`
		
		DateEnd *string `json:"dateEnd,omitempty"`
		
		MediaTypes *[]string `json:"mediaTypes,omitempty"`
		
		Programs *[]string `json:"programs,omitempty"`
		
		Dialects *[]string `json:"dialects,omitempty"`
		
		CreatedBy *Addressableentityref `json:"createdBy,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		JobStatus *string `json:"jobStatus,omitempty"`
		
		QueueOrder *int `json:"queueOrder,omitempty"`
		
		ProcessedInteractionsCount *int `json:"processedInteractionsCount,omitempty"`
		
		FailedInteractionsCount *int `json:"failedInteractionsCount,omitempty"`
		
		TotalInteractionsCount *int `json:"totalInteractionsCount,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		DateStart: DateStart,
		
		DateEnd: DateEnd,
		
		MediaTypes: o.MediaTypes,
		
		Programs: o.Programs,
		
		Dialects: o.Dialects,
		
		CreatedBy: o.CreatedBy,
		
		DateCreated: DateCreated,
		
		JobStatus: o.JobStatus,
		
		QueueOrder: o.QueueOrder,
		
		ProcessedInteractionsCount: o.ProcessedInteractionsCount,
		
		FailedInteractionsCount: o.FailedInteractionsCount,
		
		TotalInteractionsCount: o.TotalInteractionsCount,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Reprocessjobresponse) UnmarshalJSON(b []byte) error {
	var ReprocessjobresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ReprocessjobresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ReprocessjobresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ReprocessjobresponseMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := ReprocessjobresponseMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if dateStartString, ok := ReprocessjobresponseMap["dateStart"].(string); ok {
		DateStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStartString)
		o.DateStart = &DateStart
	}
	
	if dateEndString, ok := ReprocessjobresponseMap["dateEnd"].(string); ok {
		DateEnd, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateEndString)
		o.DateEnd = &DateEnd
	}
	
	if MediaTypes, ok := ReprocessjobresponseMap["mediaTypes"].([]interface{}); ok {
		MediaTypesString, _ := json.Marshal(MediaTypes)
		json.Unmarshal(MediaTypesString, &o.MediaTypes)
	}
	
	if Programs, ok := ReprocessjobresponseMap["programs"].([]interface{}); ok {
		ProgramsString, _ := json.Marshal(Programs)
		json.Unmarshal(ProgramsString, &o.Programs)
	}
	
	if Dialects, ok := ReprocessjobresponseMap["dialects"].([]interface{}); ok {
		DialectsString, _ := json.Marshal(Dialects)
		json.Unmarshal(DialectsString, &o.Dialects)
	}
	
	if CreatedBy, ok := ReprocessjobresponseMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if dateCreatedString, ok := ReprocessjobresponseMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if JobStatus, ok := ReprocessjobresponseMap["jobStatus"].(string); ok {
		o.JobStatus = &JobStatus
	}
    
	if QueueOrder, ok := ReprocessjobresponseMap["queueOrder"].(float64); ok {
		QueueOrderInt := int(QueueOrder)
		o.QueueOrder = &QueueOrderInt
	}
	
	if ProcessedInteractionsCount, ok := ReprocessjobresponseMap["processedInteractionsCount"].(float64); ok {
		ProcessedInteractionsCountInt := int(ProcessedInteractionsCount)
		o.ProcessedInteractionsCount = &ProcessedInteractionsCountInt
	}
	
	if FailedInteractionsCount, ok := ReprocessjobresponseMap["failedInteractionsCount"].(float64); ok {
		FailedInteractionsCountInt := int(FailedInteractionsCount)
		o.FailedInteractionsCount = &FailedInteractionsCountInt
	}
	
	if TotalInteractionsCount, ok := ReprocessjobresponseMap["totalInteractionsCount"].(float64); ok {
		TotalInteractionsCountInt := int(TotalInteractionsCount)
		o.TotalInteractionsCount = &TotalInteractionsCountInt
	}
	
	if SelfUri, ok := ReprocessjobresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Reprocessjobresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
