package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Assignedlearningmodule - Learning module response
type Assignedlearningmodule struct { 
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


	// CurrentAssignments - The current assignments for the requested users
	CurrentAssignments *[]Learningassignment `json:"currentAssignments,omitempty"`


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

}

func (o *Assignedlearningmodule) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Assignedlearningmodule
	
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
		
		CurrentAssignments *[]Learningassignment `json:"currentAssignments,omitempty"`
		
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
		*Alias
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
		
		CurrentAssignments: o.CurrentAssignments,
		
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
		Alias:    (*Alias)(o),
	})
}

func (o *Assignedlearningmodule) UnmarshalJSON(b []byte) error {
	var AssignedlearningmoduleMap map[string]interface{}
	err := json.Unmarshal(b, &AssignedlearningmoduleMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AssignedlearningmoduleMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := AssignedlearningmoduleMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if CreatedBy, ok := AssignedlearningmoduleMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if dateCreatedString, ok := AssignedlearningmoduleMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if ModifiedBy, ok := AssignedlearningmoduleMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if dateModifiedString, ok := AssignedlearningmoduleMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := AssignedlearningmoduleMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if ExternalId, ok := AssignedlearningmoduleMap["externalId"].(string); ok {
		o.ExternalId = &ExternalId
	}
    
	if Source, ok := AssignedlearningmoduleMap["source"].(string); ok {
		o.Source = &Source
	}
    
	if Rule, ok := AssignedlearningmoduleMap["rule"].(map[string]interface{}); ok {
		RuleString, _ := json.Marshal(Rule)
		json.Unmarshal(RuleString, &o.Rule)
	}
	
	if CurrentAssignments, ok := AssignedlearningmoduleMap["currentAssignments"].([]interface{}); ok {
		CurrentAssignmentsString, _ := json.Marshal(CurrentAssignments)
		json.Unmarshal(CurrentAssignmentsString, &o.CurrentAssignments)
	}
	
	if SelfUri, ok := AssignedlearningmoduleMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if IsArchived, ok := AssignedlearningmoduleMap["isArchived"].(bool); ok {
		o.IsArchived = &IsArchived
	}
    
	if IsPublished, ok := AssignedlearningmoduleMap["isPublished"].(bool); ok {
		o.IsPublished = &IsPublished
	}
    
	if Description, ok := AssignedlearningmoduleMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if CompletionTimeInDays, ok := AssignedlearningmoduleMap["completionTimeInDays"].(float64); ok {
		CompletionTimeInDaysInt := int(CompletionTimeInDays)
		o.CompletionTimeInDays = &CompletionTimeInDaysInt
	}
	
	if VarType, ok := AssignedlearningmoduleMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if InformSteps, ok := AssignedlearningmoduleMap["informSteps"].([]interface{}); ok {
		InformStepsString, _ := json.Marshal(InformSteps)
		json.Unmarshal(InformStepsString, &o.InformSteps)
	}
	
	if AssessmentForm, ok := AssignedlearningmoduleMap["assessmentForm"].(map[string]interface{}); ok {
		AssessmentFormString, _ := json.Marshal(AssessmentForm)
		json.Unmarshal(AssessmentFormString, &o.AssessmentForm)
	}
	
	if SummaryData, ok := AssignedlearningmoduleMap["summaryData"].(map[string]interface{}); ok {
		SummaryDataString, _ := json.Marshal(SummaryData)
		json.Unmarshal(SummaryDataString, &o.SummaryData)
	}
	
	if CoverArt, ok := AssignedlearningmoduleMap["coverArt"].(map[string]interface{}); ok {
		CoverArtString, _ := json.Marshal(CoverArt)
		json.Unmarshal(CoverArtString, &o.CoverArt)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Assignedlearningmodule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
