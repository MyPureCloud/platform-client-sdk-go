package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Workitemversion
type Workitemversion struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The name of the Workitem.
	Name *string `json:"name,omitempty"`

	// Division - The division to which this entity belongs.
	Division *Division `json:"division,omitempty"`

	// VarType - The Worktype of the Workitem.
	VarType *Worktypereference `json:"type,omitempty"`

	// Description - The description of the Workitem.
	Description *string `json:"description,omitempty"`

	// Language - The language of the Workitem.
	Language *Languagereference `json:"language,omitempty"`

	// UtilizationLabel - The utilization label of the Workitem.
	UtilizationLabel *Workitemutilizationlabelreference `json:"utilizationLabel,omitempty"`

	// Priority - The priority of the Workitem. The valid range is between -25,000,000 and 25,000,000.
	Priority *int `json:"priority,omitempty"`

	// DateCreated - The creation date of the Workitem. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - The modified date of the Workitem. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// DateDue - The due date of the Workitem. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateDue *time.Time `json:"dateDue,omitempty"`

	// DateExpires - The expiry date of the Workitem. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateExpires *time.Time `json:"dateExpires,omitempty"`

	// DurationSeconds - The estimated duration in seconds to complete the workitem.
	DurationSeconds *int `json:"durationSeconds,omitempty"`

	// Ttl - The time to live of the Workitem in seconds.
	Ttl *int `json:"ttl,omitempty"`

	// Status - The current Status of the Workitem.
	Status *Workitemstatusreference `json:"status,omitempty"`

	// StatusCategory - The Category of the current Status of the Workitem.
	StatusCategory *string `json:"statusCategory,omitempty"`

	// DateStatusChanged - The State change date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStatusChanged *time.Time `json:"dateStatusChanged,omitempty"`

	// DateClosed - The date the Workitem was closed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateClosed *time.Time `json:"dateClosed,omitempty"`

	// Workbin - The Workbin that contains the Workitem.
	Workbin *Workbinreference `json:"workbin,omitempty"`

	// Reporter - The reporter of the Workitem.
	Reporter *Userreferencewithname `json:"reporter,omitempty"`

	// Assignee - The assignee of the Workitem.
	Assignee *Userreferencewithname `json:"assignee,omitempty"`

	// ExternalContact - The external contact of the Workitem.
	ExternalContact *Externalcontactreference `json:"externalContact,omitempty"`

	// ExternalTag - The external tag of the Workitem.
	ExternalTag *string `json:"externalTag,omitempty"`

	// ModifiedBy - The User who modified the Workitem.
	ModifiedBy *Userreference `json:"modifiedBy,omitempty"`

	// Queue - The Workitems queue.
	Queue *Workitemqueuereference `json:"queue,omitempty"`

	// AssignmentState - The assignment state of the workitem.
	AssignmentState *string `json:"assignmentState,omitempty"`

	// DateAssignmentStateChanged - The assignment state change date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateAssignmentStateChanged *time.Time `json:"dateAssignmentStateChanged,omitempty"`

	// AlertTimeoutSeconds - The duration in seconds before an alert will timeout.
	AlertTimeoutSeconds *int `json:"alertTimeoutSeconds,omitempty"`

	// Skills - The skills of the Workitem.
	Skills *[]Routingskillreference `json:"skills,omitempty"`

	// PreferredAgents - The preferred agents of the Workitem.
	PreferredAgents *[]Userreference `json:"preferredAgents,omitempty"`

	// AutoStatusTransition - Set it to false to disable auto status transition. By default, it is enabled.
	AutoStatusTransition *bool `json:"autoStatusTransition,omitempty"`

	// Schema - The schema defining the custom fields of the Workitem. The schema is inherited from the Workitems Worktype at creation time.
	Schema *Workitemschema `json:"schema,omitempty"`

	// CustomFields - Custom fields defined in the schema referenced by the Workitem.
	CustomFields *map[string]interface{} `json:"customFields,omitempty"`

	// AutoStatusTransitionDetail - Auto status transition details of Workitem.
	AutoStatusTransitionDetail *Autostatustransitiondetail `json:"autoStatusTransitionDetail,omitempty"`

	// ScoredAgents - A list of scored agents for the Workitem.
	ScoredAgents *[]Workitemscoredagent `json:"scoredAgents,omitempty"`

	// Version - Version
	Version *int `json:"version,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Workitemversion) SetField(field string, fieldValue interface{}) {
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

func (o Workitemversion) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateModified","DateDue","DateExpires","DateStatusChanged","DateClosed","DateAssignmentStateChanged", }
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
	type Alias Workitemversion
	
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
	
	DateDue := new(string)
	if o.DateDue != nil {
		
		*DateDue = timeutil.Strftime(o.DateDue, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateDue = nil
	}
	
	DateExpires := new(string)
	if o.DateExpires != nil {
		
		*DateExpires = timeutil.Strftime(o.DateExpires, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateExpires = nil
	}
	
	DateStatusChanged := new(string)
	if o.DateStatusChanged != nil {
		
		*DateStatusChanged = timeutil.Strftime(o.DateStatusChanged, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStatusChanged = nil
	}
	
	DateClosed := new(string)
	if o.DateClosed != nil {
		
		*DateClosed = timeutil.Strftime(o.DateClosed, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateClosed = nil
	}
	
	DateAssignmentStateChanged := new(string)
	if o.DateAssignmentStateChanged != nil {
		
		*DateAssignmentStateChanged = timeutil.Strftime(o.DateAssignmentStateChanged, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateAssignmentStateChanged = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Division *Division `json:"division,omitempty"`
		
		VarType *Worktypereference `json:"type,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Language *Languagereference `json:"language,omitempty"`
		
		UtilizationLabel *Workitemutilizationlabelreference `json:"utilizationLabel,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		DateDue *string `json:"dateDue,omitempty"`
		
		DateExpires *string `json:"dateExpires,omitempty"`
		
		DurationSeconds *int `json:"durationSeconds,omitempty"`
		
		Ttl *int `json:"ttl,omitempty"`
		
		Status *Workitemstatusreference `json:"status,omitempty"`
		
		StatusCategory *string `json:"statusCategory,omitempty"`
		
		DateStatusChanged *string `json:"dateStatusChanged,omitempty"`
		
		DateClosed *string `json:"dateClosed,omitempty"`
		
		Workbin *Workbinreference `json:"workbin,omitempty"`
		
		Reporter *Userreferencewithname `json:"reporter,omitempty"`
		
		Assignee *Userreferencewithname `json:"assignee,omitempty"`
		
		ExternalContact *Externalcontactreference `json:"externalContact,omitempty"`
		
		ExternalTag *string `json:"externalTag,omitempty"`
		
		ModifiedBy *Userreference `json:"modifiedBy,omitempty"`
		
		Queue *Workitemqueuereference `json:"queue,omitempty"`
		
		AssignmentState *string `json:"assignmentState,omitempty"`
		
		DateAssignmentStateChanged *string `json:"dateAssignmentStateChanged,omitempty"`
		
		AlertTimeoutSeconds *int `json:"alertTimeoutSeconds,omitempty"`
		
		Skills *[]Routingskillreference `json:"skills,omitempty"`
		
		PreferredAgents *[]Userreference `json:"preferredAgents,omitempty"`
		
		AutoStatusTransition *bool `json:"autoStatusTransition,omitempty"`
		
		Schema *Workitemschema `json:"schema,omitempty"`
		
		CustomFields *map[string]interface{} `json:"customFields,omitempty"`
		
		AutoStatusTransitionDetail *Autostatustransitiondetail `json:"autoStatusTransitionDetail,omitempty"`
		
		ScoredAgents *[]Workitemscoredagent `json:"scoredAgents,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Division: o.Division,
		
		VarType: o.VarType,
		
		Description: o.Description,
		
		Language: o.Language,
		
		UtilizationLabel: o.UtilizationLabel,
		
		Priority: o.Priority,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		DateDue: DateDue,
		
		DateExpires: DateExpires,
		
		DurationSeconds: o.DurationSeconds,
		
		Ttl: o.Ttl,
		
		Status: o.Status,
		
		StatusCategory: o.StatusCategory,
		
		DateStatusChanged: DateStatusChanged,
		
		DateClosed: DateClosed,
		
		Workbin: o.Workbin,
		
		Reporter: o.Reporter,
		
		Assignee: o.Assignee,
		
		ExternalContact: o.ExternalContact,
		
		ExternalTag: o.ExternalTag,
		
		ModifiedBy: o.ModifiedBy,
		
		Queue: o.Queue,
		
		AssignmentState: o.AssignmentState,
		
		DateAssignmentStateChanged: DateAssignmentStateChanged,
		
		AlertTimeoutSeconds: o.AlertTimeoutSeconds,
		
		Skills: o.Skills,
		
		PreferredAgents: o.PreferredAgents,
		
		AutoStatusTransition: o.AutoStatusTransition,
		
		Schema: o.Schema,
		
		CustomFields: o.CustomFields,
		
		AutoStatusTransitionDetail: o.AutoStatusTransitionDetail,
		
		ScoredAgents: o.ScoredAgents,
		
		Version: o.Version,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Workitemversion) UnmarshalJSON(b []byte) error {
	var WorkitemversionMap map[string]interface{}
	err := json.Unmarshal(b, &WorkitemversionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WorkitemversionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := WorkitemversionMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Division, ok := WorkitemversionMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if VarType, ok := WorkitemversionMap["type"].(map[string]interface{}); ok {
		VarTypeString, _ := json.Marshal(VarType)
		json.Unmarshal(VarTypeString, &o.VarType)
	}
	
	if Description, ok := WorkitemversionMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Language, ok := WorkitemversionMap["language"].(map[string]interface{}); ok {
		LanguageString, _ := json.Marshal(Language)
		json.Unmarshal(LanguageString, &o.Language)
	}
	
	if UtilizationLabel, ok := WorkitemversionMap["utilizationLabel"].(map[string]interface{}); ok {
		UtilizationLabelString, _ := json.Marshal(UtilizationLabel)
		json.Unmarshal(UtilizationLabelString, &o.UtilizationLabel)
	}
	
	if Priority, ok := WorkitemversionMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	
	if dateCreatedString, ok := WorkitemversionMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := WorkitemversionMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if dateDueString, ok := WorkitemversionMap["dateDue"].(string); ok {
		DateDue, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateDueString)
		o.DateDue = &DateDue
	}
	
	if dateExpiresString, ok := WorkitemversionMap["dateExpires"].(string); ok {
		DateExpires, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateExpiresString)
		o.DateExpires = &DateExpires
	}
	
	if DurationSeconds, ok := WorkitemversionMap["durationSeconds"].(float64); ok {
		DurationSecondsInt := int(DurationSeconds)
		o.DurationSeconds = &DurationSecondsInt
	}
	
	if Ttl, ok := WorkitemversionMap["ttl"].(float64); ok {
		TtlInt := int(Ttl)
		o.Ttl = &TtlInt
	}
	
	if Status, ok := WorkitemversionMap["status"].(map[string]interface{}); ok {
		StatusString, _ := json.Marshal(Status)
		json.Unmarshal(StatusString, &o.Status)
	}
	
	if StatusCategory, ok := WorkitemversionMap["statusCategory"].(string); ok {
		o.StatusCategory = &StatusCategory
	}
    
	if dateStatusChangedString, ok := WorkitemversionMap["dateStatusChanged"].(string); ok {
		DateStatusChanged, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStatusChangedString)
		o.DateStatusChanged = &DateStatusChanged
	}
	
	if dateClosedString, ok := WorkitemversionMap["dateClosed"].(string); ok {
		DateClosed, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateClosedString)
		o.DateClosed = &DateClosed
	}
	
	if Workbin, ok := WorkitemversionMap["workbin"].(map[string]interface{}); ok {
		WorkbinString, _ := json.Marshal(Workbin)
		json.Unmarshal(WorkbinString, &o.Workbin)
	}
	
	if Reporter, ok := WorkitemversionMap["reporter"].(map[string]interface{}); ok {
		ReporterString, _ := json.Marshal(Reporter)
		json.Unmarshal(ReporterString, &o.Reporter)
	}
	
	if Assignee, ok := WorkitemversionMap["assignee"].(map[string]interface{}); ok {
		AssigneeString, _ := json.Marshal(Assignee)
		json.Unmarshal(AssigneeString, &o.Assignee)
	}
	
	if ExternalContact, ok := WorkitemversionMap["externalContact"].(map[string]interface{}); ok {
		ExternalContactString, _ := json.Marshal(ExternalContact)
		json.Unmarshal(ExternalContactString, &o.ExternalContact)
	}
	
	if ExternalTag, ok := WorkitemversionMap["externalTag"].(string); ok {
		o.ExternalTag = &ExternalTag
	}
    
	if ModifiedBy, ok := WorkitemversionMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if Queue, ok := WorkitemversionMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if AssignmentState, ok := WorkitemversionMap["assignmentState"].(string); ok {
		o.AssignmentState = &AssignmentState
	}
    
	if dateAssignmentStateChangedString, ok := WorkitemversionMap["dateAssignmentStateChanged"].(string); ok {
		DateAssignmentStateChanged, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateAssignmentStateChangedString)
		o.DateAssignmentStateChanged = &DateAssignmentStateChanged
	}
	
	if AlertTimeoutSeconds, ok := WorkitemversionMap["alertTimeoutSeconds"].(float64); ok {
		AlertTimeoutSecondsInt := int(AlertTimeoutSeconds)
		o.AlertTimeoutSeconds = &AlertTimeoutSecondsInt
	}
	
	if Skills, ok := WorkitemversionMap["skills"].([]interface{}); ok {
		SkillsString, _ := json.Marshal(Skills)
		json.Unmarshal(SkillsString, &o.Skills)
	}
	
	if PreferredAgents, ok := WorkitemversionMap["preferredAgents"].([]interface{}); ok {
		PreferredAgentsString, _ := json.Marshal(PreferredAgents)
		json.Unmarshal(PreferredAgentsString, &o.PreferredAgents)
	}
	
	if AutoStatusTransition, ok := WorkitemversionMap["autoStatusTransition"].(bool); ok {
		o.AutoStatusTransition = &AutoStatusTransition
	}
    
	if Schema, ok := WorkitemversionMap["schema"].(map[string]interface{}); ok {
		SchemaString, _ := json.Marshal(Schema)
		json.Unmarshal(SchemaString, &o.Schema)
	}
	
	if CustomFields, ok := WorkitemversionMap["customFields"].(map[string]interface{}); ok {
		CustomFieldsString, _ := json.Marshal(CustomFields)
		json.Unmarshal(CustomFieldsString, &o.CustomFields)
	}
	
	if AutoStatusTransitionDetail, ok := WorkitemversionMap["autoStatusTransitionDetail"].(map[string]interface{}); ok {
		AutoStatusTransitionDetailString, _ := json.Marshal(AutoStatusTransitionDetail)
		json.Unmarshal(AutoStatusTransitionDetailString, &o.AutoStatusTransitionDetail)
	}
	
	if ScoredAgents, ok := WorkitemversionMap["scoredAgents"].([]interface{}); ok {
		ScoredAgentsString, _ := json.Marshal(ScoredAgents)
		json.Unmarshal(ScoredAgentsString, &o.ScoredAgents)
	}
	
	if Version, ok := WorkitemversionMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if SelfUri, ok := WorkitemversionMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Workitemversion) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
