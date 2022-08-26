package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgeimportjobstatistics
type Knowledgeimportjobstatistics struct { 
	// CountDocumentImportActivityCreate - Number of documents will be created by the import.
	CountDocumentImportActivityCreate *int `json:"countDocumentImportActivityCreate,omitempty"`


	// CountDocumentImportActivityUpdate - Number of documents will be updated by the import.
	CountDocumentImportActivityUpdate *int `json:"countDocumentImportActivityUpdate,omitempty"`


	// CountDocumentStateDraft - Number of documents will be imported as draft.
	CountDocumentStateDraft *int `json:"countDocumentStateDraft,omitempty"`


	// CountDocumentStatePublished - Number of documents will be imported as published.
	CountDocumentStatePublished *int `json:"countDocumentStatePublished,omitempty"`


	// CountDocumentValidationSuccess - Number of documents that validated successfully for import.
	CountDocumentValidationSuccess *int `json:"countDocumentValidationSuccess,omitempty"`


	// CountDocumentValidationFailure - Number of documents that failed validation for import.
	CountDocumentValidationFailure *int `json:"countDocumentValidationFailure,omitempty"`


	// CountDocumentImportSuccess - Number of imported documents.
	CountDocumentImportSuccess *int `json:"countDocumentImportSuccess,omitempty"`


	// CountDocumentImportFailure - Number of documents failed to import.
	CountDocumentImportFailure *int `json:"countDocumentImportFailure,omitempty"`

}

func (o *Knowledgeimportjobstatistics) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgeimportjobstatistics
	
	return json.Marshal(&struct { 
		CountDocumentImportActivityCreate *int `json:"countDocumentImportActivityCreate,omitempty"`
		
		CountDocumentImportActivityUpdate *int `json:"countDocumentImportActivityUpdate,omitempty"`
		
		CountDocumentStateDraft *int `json:"countDocumentStateDraft,omitempty"`
		
		CountDocumentStatePublished *int `json:"countDocumentStatePublished,omitempty"`
		
		CountDocumentValidationSuccess *int `json:"countDocumentValidationSuccess,omitempty"`
		
		CountDocumentValidationFailure *int `json:"countDocumentValidationFailure,omitempty"`
		
		CountDocumentImportSuccess *int `json:"countDocumentImportSuccess,omitempty"`
		
		CountDocumentImportFailure *int `json:"countDocumentImportFailure,omitempty"`
		*Alias
	}{ 
		CountDocumentImportActivityCreate: o.CountDocumentImportActivityCreate,
		
		CountDocumentImportActivityUpdate: o.CountDocumentImportActivityUpdate,
		
		CountDocumentStateDraft: o.CountDocumentStateDraft,
		
		CountDocumentStatePublished: o.CountDocumentStatePublished,
		
		CountDocumentValidationSuccess: o.CountDocumentValidationSuccess,
		
		CountDocumentValidationFailure: o.CountDocumentValidationFailure,
		
		CountDocumentImportSuccess: o.CountDocumentImportSuccess,
		
		CountDocumentImportFailure: o.CountDocumentImportFailure,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgeimportjobstatistics) UnmarshalJSON(b []byte) error {
	var KnowledgeimportjobstatisticsMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgeimportjobstatisticsMap)
	if err != nil {
		return err
	}
	
	if CountDocumentImportActivityCreate, ok := KnowledgeimportjobstatisticsMap["countDocumentImportActivityCreate"].(float64); ok {
		CountDocumentImportActivityCreateInt := int(CountDocumentImportActivityCreate)
		o.CountDocumentImportActivityCreate = &CountDocumentImportActivityCreateInt
	}
	
	if CountDocumentImportActivityUpdate, ok := KnowledgeimportjobstatisticsMap["countDocumentImportActivityUpdate"].(float64); ok {
		CountDocumentImportActivityUpdateInt := int(CountDocumentImportActivityUpdate)
		o.CountDocumentImportActivityUpdate = &CountDocumentImportActivityUpdateInt
	}
	
	if CountDocumentStateDraft, ok := KnowledgeimportjobstatisticsMap["countDocumentStateDraft"].(float64); ok {
		CountDocumentStateDraftInt := int(CountDocumentStateDraft)
		o.CountDocumentStateDraft = &CountDocumentStateDraftInt
	}
	
	if CountDocumentStatePublished, ok := KnowledgeimportjobstatisticsMap["countDocumentStatePublished"].(float64); ok {
		CountDocumentStatePublishedInt := int(CountDocumentStatePublished)
		o.CountDocumentStatePublished = &CountDocumentStatePublishedInt
	}
	
	if CountDocumentValidationSuccess, ok := KnowledgeimportjobstatisticsMap["countDocumentValidationSuccess"].(float64); ok {
		CountDocumentValidationSuccessInt := int(CountDocumentValidationSuccess)
		o.CountDocumentValidationSuccess = &CountDocumentValidationSuccessInt
	}
	
	if CountDocumentValidationFailure, ok := KnowledgeimportjobstatisticsMap["countDocumentValidationFailure"].(float64); ok {
		CountDocumentValidationFailureInt := int(CountDocumentValidationFailure)
		o.CountDocumentValidationFailure = &CountDocumentValidationFailureInt
	}
	
	if CountDocumentImportSuccess, ok := KnowledgeimportjobstatisticsMap["countDocumentImportSuccess"].(float64); ok {
		CountDocumentImportSuccessInt := int(CountDocumentImportSuccess)
		o.CountDocumentImportSuccess = &CountDocumentImportSuccessInt
	}
	
	if CountDocumentImportFailure, ok := KnowledgeimportjobstatisticsMap["countDocumentImportFailure"].(float64); ok {
		CountDocumentImportFailureInt := int(CountDocumentImportFailure)
		o.CountDocumentImportFailure = &CountDocumentImportFailureInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgeimportjobstatistics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
