package platformclientv2
import (
	"encoding/json"
)

// Schemaquantitylimits
type Schemaquantitylimits struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// MinFieldNameCharacters - The minimum number of schema field name characters allowed.
	MinFieldNameCharacters *int32 `json:"minFieldNameCharacters,omitempty"`


	// MaxFieldNameCharacters - The maximum number of schema field name characters allowed.
	MaxFieldNameCharacters *int32 `json:"maxFieldNameCharacters,omitempty"`


	// MinFieldDescriptionCharacters - The minimum number of schema field description characters allowed.
	MinFieldDescriptionCharacters *int32 `json:"minFieldDescriptionCharacters,omitempty"`


	// MaxFieldDescriptionCharacters - The maximum number of schema field description characters allowed.
	MaxFieldDescriptionCharacters *int32 `json:"maxFieldDescriptionCharacters,omitempty"`


	// MinSchemaNameCharacters - The minimum number of schema name characters allowed.
	MinSchemaNameCharacters *int32 `json:"minSchemaNameCharacters,omitempty"`


	// MaxSchemaNameCharacters - The maximum number of schema name characters allowed.
	MaxSchemaNameCharacters *int32 `json:"maxSchemaNameCharacters,omitempty"`


	// MinSchemaDescriptionCharacters - The minimum number of schema description characters allowed.
	MinSchemaDescriptionCharacters *int32 `json:"minSchemaDescriptionCharacters,omitempty"`


	// MaxSchemaDescriptionCharacters - The maximum number of schema description characters allowed.
	MaxSchemaDescriptionCharacters *int32 `json:"maxSchemaDescriptionCharacters,omitempty"`


	// MaxNumberOfSchemasPerOrg - The maximum number of schema allowed per org.
	MaxNumberOfSchemasPerOrg *int32 `json:"maxNumberOfSchemasPerOrg,omitempty"`


	// MaxNumberOfFieldsPerSchema - The maximum number of schema fields allowed per schema.
	MaxNumberOfFieldsPerSchema *int32 `json:"maxNumberOfFieldsPerSchema,omitempty"`


	// MaxNumberOfFieldsPerOrg - The maximum number of schema fields allowed per organization across all of their schemas.
	MaxNumberOfFieldsPerOrg *int32 `json:"maxNumberOfFieldsPerOrg,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Schemaquantitylimits) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
