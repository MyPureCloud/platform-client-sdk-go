package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Alternativeshiftjobresponse
type Alternativeshiftjobresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Status - The status of the alternative shift job
	Status *string `json:"status,omitempty"`

	// VarType - The type of job
	VarType *string `json:"type,omitempty"`

	// DownloadUrl - The URL where completed results are available, only set if status == 'Complete'
	DownloadUrl *string `json:"downloadUrl,omitempty"`

	// VarError - Any error information, only set if the status == 'Error'
	VarError *Errorbody `json:"error,omitempty"`

	// ViewOffersResults - Schema template for deserializing data returned from the downloadUrl. Use if type == 'ListOffers' or 'SearchOffers'
	ViewOffersResults *Alternativeshiftoffersviewresponsetemplate `json:"viewOffersResults,omitempty"`

	// ViewTradesResults - Schema template for deserializing data returned from the downloadUrl. Use if type == 'ListUserTrades' or 'SearchTrades'
	ViewTradesResults *Alternativeshifttradesviewresponsetemplate `json:"viewTradesResults,omitempty"`

	// BulkUpdateTradesResults - Schema template for deserializing data returned from the downloadUrl. Use if type == 'BulkUpdateTrades'
	BulkUpdateTradesResults *Alternativeshiftbulkupdatetradesresponsetemplate `json:"bulkUpdateTradesResults,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Alternativeshiftjobresponse) SetField(field string, fieldValue interface{}) {
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

func (o Alternativeshiftjobresponse) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
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
	type Alias Alternativeshiftjobresponse
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		DownloadUrl *string `json:"downloadUrl,omitempty"`
		
		VarError *Errorbody `json:"error,omitempty"`
		
		ViewOffersResults *Alternativeshiftoffersviewresponsetemplate `json:"viewOffersResults,omitempty"`
		
		ViewTradesResults *Alternativeshifttradesviewresponsetemplate `json:"viewTradesResults,omitempty"`
		
		BulkUpdateTradesResults *Alternativeshiftbulkupdatetradesresponsetemplate `json:"bulkUpdateTradesResults,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Status: o.Status,
		
		VarType: o.VarType,
		
		DownloadUrl: o.DownloadUrl,
		
		VarError: o.VarError,
		
		ViewOffersResults: o.ViewOffersResults,
		
		ViewTradesResults: o.ViewTradesResults,
		
		BulkUpdateTradesResults: o.BulkUpdateTradesResults,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Alternativeshiftjobresponse) UnmarshalJSON(b []byte) error {
	var AlternativeshiftjobresponseMap map[string]interface{}
	err := json.Unmarshal(b, &AlternativeshiftjobresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AlternativeshiftjobresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Status, ok := AlternativeshiftjobresponseMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if VarType, ok := AlternativeshiftjobresponseMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if DownloadUrl, ok := AlternativeshiftjobresponseMap["downloadUrl"].(string); ok {
		o.DownloadUrl = &DownloadUrl
	}
    
	if VarError, ok := AlternativeshiftjobresponseMap["error"].(map[string]interface{}); ok {
		VarErrorString, _ := json.Marshal(VarError)
		json.Unmarshal(VarErrorString, &o.VarError)
	}
	
	if ViewOffersResults, ok := AlternativeshiftjobresponseMap["viewOffersResults"].(map[string]interface{}); ok {
		ViewOffersResultsString, _ := json.Marshal(ViewOffersResults)
		json.Unmarshal(ViewOffersResultsString, &o.ViewOffersResults)
	}
	
	if ViewTradesResults, ok := AlternativeshiftjobresponseMap["viewTradesResults"].(map[string]interface{}); ok {
		ViewTradesResultsString, _ := json.Marshal(ViewTradesResults)
		json.Unmarshal(ViewTradesResultsString, &o.ViewTradesResults)
	}
	
	if BulkUpdateTradesResults, ok := AlternativeshiftjobresponseMap["bulkUpdateTradesResults"].(map[string]interface{}); ok {
		BulkUpdateTradesResultsString, _ := json.Marshal(BulkUpdateTradesResults)
		json.Unmarshal(BulkUpdateTradesResultsString, &o.BulkUpdateTradesResults)
	}
	
	if SelfUri, ok := AlternativeshiftjobresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Alternativeshiftjobresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
