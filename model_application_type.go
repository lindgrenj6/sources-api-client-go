/*
 * Sources
 *
 * Sources
 *
 * API version: 3.0.0
 * Contact: support@redhat.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sourcesapi

import (
	"encoding/json"
	"time"
)

// ApplicationType struct for ApplicationType
type ApplicationType struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	DependentApplications *map[string]interface{} `json:"dependent_applications,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	// ID of the resource
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	SupportedAuthenticationTypes *map[string]interface{} `json:"supported_authentication_types,omitempty"`
	SupportedSourceTypes *map[string]interface{} `json:"supported_source_types,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewApplicationType instantiates a new ApplicationType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationType() *ApplicationType {
	this := ApplicationType{}
	return &this
}

// NewApplicationTypeWithDefaults instantiates a new ApplicationType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationTypeWithDefaults() *ApplicationType {
	this := ApplicationType{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ApplicationType) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationType) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ApplicationType) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ApplicationType) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDependentApplications returns the DependentApplications field value if set, zero value otherwise.
func (o *ApplicationType) GetDependentApplications() map[string]interface{} {
	if o == nil || o.DependentApplications == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.DependentApplications
}

// GetDependentApplicationsOk returns a tuple with the DependentApplications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationType) GetDependentApplicationsOk() (*map[string]interface{}, bool) {
	if o == nil || o.DependentApplications == nil {
		return nil, false
	}
	return o.DependentApplications, true
}

// HasDependentApplications returns a boolean if a field has been set.
func (o *ApplicationType) HasDependentApplications() bool {
	if o != nil && o.DependentApplications != nil {
		return true
	}

	return false
}

// SetDependentApplications gets a reference to the given map[string]interface{} and assigns it to the DependentApplications field.
func (o *ApplicationType) SetDependentApplications(v map[string]interface{}) {
	o.DependentApplications = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ApplicationType) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationType) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ApplicationType) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ApplicationType) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplicationType) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationType) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApplicationType) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApplicationType) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApplicationType) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationType) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApplicationType) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApplicationType) SetName(v string) {
	o.Name = &v
}

// GetSupportedAuthenticationTypes returns the SupportedAuthenticationTypes field value if set, zero value otherwise.
func (o *ApplicationType) GetSupportedAuthenticationTypes() map[string]interface{} {
	if o == nil || o.SupportedAuthenticationTypes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.SupportedAuthenticationTypes
}

// GetSupportedAuthenticationTypesOk returns a tuple with the SupportedAuthenticationTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationType) GetSupportedAuthenticationTypesOk() (*map[string]interface{}, bool) {
	if o == nil || o.SupportedAuthenticationTypes == nil {
		return nil, false
	}
	return o.SupportedAuthenticationTypes, true
}

// HasSupportedAuthenticationTypes returns a boolean if a field has been set.
func (o *ApplicationType) HasSupportedAuthenticationTypes() bool {
	if o != nil && o.SupportedAuthenticationTypes != nil {
		return true
	}

	return false
}

// SetSupportedAuthenticationTypes gets a reference to the given map[string]interface{} and assigns it to the SupportedAuthenticationTypes field.
func (o *ApplicationType) SetSupportedAuthenticationTypes(v map[string]interface{}) {
	o.SupportedAuthenticationTypes = &v
}

// GetSupportedSourceTypes returns the SupportedSourceTypes field value if set, zero value otherwise.
func (o *ApplicationType) GetSupportedSourceTypes() map[string]interface{} {
	if o == nil || o.SupportedSourceTypes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.SupportedSourceTypes
}

// GetSupportedSourceTypesOk returns a tuple with the SupportedSourceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationType) GetSupportedSourceTypesOk() (*map[string]interface{}, bool) {
	if o == nil || o.SupportedSourceTypes == nil {
		return nil, false
	}
	return o.SupportedSourceTypes, true
}

// HasSupportedSourceTypes returns a boolean if a field has been set.
func (o *ApplicationType) HasSupportedSourceTypes() bool {
	if o != nil && o.SupportedSourceTypes != nil {
		return true
	}

	return false
}

// SetSupportedSourceTypes gets a reference to the given map[string]interface{} and assigns it to the SupportedSourceTypes field.
func (o *ApplicationType) SetSupportedSourceTypes(v map[string]interface{}) {
	o.SupportedSourceTypes = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ApplicationType) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationType) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ApplicationType) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ApplicationType) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o ApplicationType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.DependentApplications != nil {
		toSerialize["dependent_applications"] = o.DependentApplications
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SupportedAuthenticationTypes != nil {
		toSerialize["supported_authentication_types"] = o.SupportedAuthenticationTypes
	}
	if o.SupportedSourceTypes != nil {
		toSerialize["supported_source_types"] = o.SupportedSourceTypes
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationType struct {
	value *ApplicationType
	isSet bool
}

func (v NullableApplicationType) Get() *ApplicationType {
	return v.value
}

func (v *NullableApplicationType) Set(val *ApplicationType) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationType) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationType(val *ApplicationType) *NullableApplicationType {
	return &NullableApplicationType{value: val, isSet: true}
}

func (v NullableApplicationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


