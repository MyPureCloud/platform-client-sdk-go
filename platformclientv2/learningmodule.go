package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningmodule - Learning module response
type Learningmodule struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The name of learning module
	Name *string `json:"name,omitempty"`

	// CreatedBy - The user who created learning module
	CreatedBy *Userreference `json:"createdBy,omitempty"`

	// DateCreated - The date/time learning module was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// ModifiedBy - The user who modified learning module
	ModifiedBy *Userreference `json:"modifiedBy,omitempty"`

	// DateModified - The date/time learning module was modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// Version - The version of published learning module
	Version *int `json:"version,omitempty"`

	// ExternalId - The external ID of the learning module
	ExternalId *string `json:"externalId,omitempty"`

	// Source - The source of the learning module
	Source *string `json:"source,omitempty"`

	// Rule - The rule for learning module; read-only, and only populated when requested via expand param.
	Rule *Learningmodulerule `json:"rule,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

	// IsArchived - If true, learning module is archived
	IsArchived *bool `json:"isArchived,omitempty"`

	// IsPublished - If true, learning module is published
	IsPublished *bool `json:"isPublished,omitempty"`

	// Description - The description of learning module
	Description *string `json:"description,omitempty"`

	// CompletionTimeInDays - The completion time of learning module in days
	CompletionTimeInDays *int `json:"completionTimeInDays,omitempty"`

	// VarType - The type for the learning module
	VarType *string `json:"type,omitempty"`

	// InformSteps - The list of inform steps in a learning module
	InformSteps *[]Learningmoduleinformstep `json:"informSteps,omitempty"`

	// AssessmentForm - The assessment form for learning module
	AssessmentForm *Assessmentform `json:"assessmentForm,omitempty"`

	// SummaryData - The learning module summary data
	SummaryData *Learningmodulesummary `json:"summaryData,omitempty"`

	// CoverArt - The cover art for the learning module
	CoverArt *Learningmodulecoverartresponse `json:"coverArt,omitempty"`

	// ArchivalMode - The mode of archival for learning module
	ArchivalMode *string `json:"archivalMode,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Learningmodule) SetField(field string, fieldValue interface{}) {
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

func (o Learningmodule) MarshalJSON() ([]byte, error) {
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
			if contains(dateTimeFields, fieldName) {
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
	type Alias Learningmodule
	
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
		
		CreatedBy *Userreference `json:"createdBy,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		ModifiedBy *Userreference `json:"modifiedBy,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		ExternalId *string `json:"externalId,omitempty"`
		
		Source *string `json:"source,omitempty"`
		
		Rule *Learningmodulerule `json:"rule,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		IsArchived *bool `json:"isArchived,omitempty"`
		
		IsPublished *bool `json:"isPublished,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		CompletionTimeInDays *int `json:"completionTimeInDays,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		InformSteps *[]Learningmoduleinformstep `json:"informSteps,omitempty"`
		
		AssessmentForm *Assessmentform `json:"assessmentForm,omitempty"`
		
		SummaryData *Learningmodulesummary `json:"summaryData,omitempty"`
		
		CoverArt *Learningmodulecoverartresponse `json:"coverArt,omitempty"`
		
		ArchivalMode *string `json:"archivalMode,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		CreatedBy: o.CreatedBy,
		
		DateCreated: DateCreated,
		
		ModifiedBy: o.ModifiedBy,
		
		DateModified: DateModified,
		
		Version: o.Version,
		
		ExternalId: o.ExternalId,
		
		Source: o.Source,
		
		Rule: o.Rule,
		
		SelfUri: o.SelfUri,
		
		IsArchived: o.IsArchived,
		
		IsPublished: o.IsPublished,
		
		Description: o.Description,
		
		CompletionTimeInDays: o.CompletionTimeInDays,
		
		VarType: o.VarType,
		
		InformSteps: o.InformSteps,
		
		AssessmentForm: o.AssessmentForm,
		
		SummaryData: o.SummaryData,
		
		CoverArt: o.CoverArt,
		
		ArchivalMode: o.ArchivalMode,
		Alias:    (Alias)(o),
	})
}

func (o *Learningmodule) UnmarshalJSON(b []byte) error {
	var LearningmoduleMap map[string]interface{}
	err := json.Unmarshal(b, &LearningmoduleMap)
	if err != nil {
		return err
	}
	
	if Id, ok := LearningmoduleMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := LearningmoduleMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if CreatedBy, ok := LearningmoduleMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if dateCreatedString, ok := LearningmoduleMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if ModifiedBy, ok := LearningmoduleMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if dateModifiedString, ok := LearningmoduleMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := LearningmoduleMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if ExternalId, ok := LearningmoduleMap["externalId"].(string); ok {
		o.ExternalId = &ExternalId
	}
    
	if Source, ok := LearningmoduleMap["source"].(string); ok {
		o.Source = &Source
	}
    
	if Rule, ok := LearningmoduleMap["rule"].(map[string]interface{}); ok {
		RuleString, _ := json.Marshal(Rule)
		json.Unmarshal(RuleString, &o.Rule)
	}
	
	if SelfUri, ok := LearningmoduleMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if IsArchived, ok := LearningmoduleMap["isArchived"].(bool); ok {
		o.IsArchived = &IsArchived
	}
    
	if IsPublished, ok := LearningmoduleMap["isPublished"].(bool); ok {
		o.IsPublished = &IsPublished
	}
    
	if Description, ok := LearningmoduleMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if CompletionTimeInDays, ok := LearningmoduleMap["completionTimeInDays"].(float64); ok {
		CompletionTimeInDaysInt := int(CompletionTimeInDays)
		o.CompletionTimeInDays = &CompletionTimeInDaysInt
	}
	
	if VarType, ok := LearningmoduleMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if InformSteps, ok := LearningmoduleMap["informSteps"].([]interface{}); ok {
		InformStepsString, _ := json.Marshal(InformSteps)
		json.Unmarshal(InformStepsString, &o.InformSteps)
	}
	
	if AssessmentForm, ok := LearningmoduleMap["assessmentForm"].(map[string]interface{}); ok {
		AssessmentFormString, _ := json.Marshal(AssessmentForm)
		json.Unmarshal(AssessmentFormString, &o.AssessmentForm)
	}
	
	if SummaryData, ok := LearningmoduleMap["summaryData"].(map[string]interface{}); ok {
		SummaryDataString, _ := json.Marshal(SummaryData)
		json.Unmarshal(SummaryDataString, &o.SummaryData)
	}
	
	if CoverArt, ok := LearningmoduleMap["coverArt"].(map[string]interface{}); ok {
		CoverArtString, _ := json.Marshal(CoverArt)
		json.Unmarshal(CoverArtString, &o.CoverArt)
	}
	
	if ArchivalMode, ok := LearningmoduleMap["archivalMode"].(string); ok {
		o.ArchivalMode = &ArchivalMode
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Learningmodule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
