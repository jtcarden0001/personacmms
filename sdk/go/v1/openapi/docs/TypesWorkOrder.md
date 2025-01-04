# TypesWorkOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedDate** | Pointer to **string** |  | [optional] 
**CumulativeHours** | Pointer to **int32** |  | [optional] 
**CumulativeMiles** | Pointer to **int32** |  | [optional] 
**Instructions** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**Title** | **string** |  | 

## Methods

### NewTypesWorkOrder

`func NewTypesWorkOrder(status string, title string, ) *TypesWorkOrder`

NewTypesWorkOrder instantiates a new TypesWorkOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesWorkOrderWithDefaults

`func NewTypesWorkOrderWithDefaults() *TypesWorkOrder`

NewTypesWorkOrderWithDefaults instantiates a new TypesWorkOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedDate

`func (o *TypesWorkOrder) GetCompletedDate() string`

GetCompletedDate returns the CompletedDate field if non-nil, zero value otherwise.

### GetCompletedDateOk

`func (o *TypesWorkOrder) GetCompletedDateOk() (*string, bool)`

GetCompletedDateOk returns a tuple with the CompletedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedDate

`func (o *TypesWorkOrder) SetCompletedDate(v string)`

SetCompletedDate sets CompletedDate field to given value.

### HasCompletedDate

`func (o *TypesWorkOrder) HasCompletedDate() bool`

HasCompletedDate returns a boolean if a field has been set.

### GetCumulativeHours

`func (o *TypesWorkOrder) GetCumulativeHours() int32`

GetCumulativeHours returns the CumulativeHours field if non-nil, zero value otherwise.

### GetCumulativeHoursOk

`func (o *TypesWorkOrder) GetCumulativeHoursOk() (*int32, bool)`

GetCumulativeHoursOk returns a tuple with the CumulativeHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeHours

`func (o *TypesWorkOrder) SetCumulativeHours(v int32)`

SetCumulativeHours sets CumulativeHours field to given value.

### HasCumulativeHours

`func (o *TypesWorkOrder) HasCumulativeHours() bool`

HasCumulativeHours returns a boolean if a field has been set.

### GetCumulativeMiles

`func (o *TypesWorkOrder) GetCumulativeMiles() int32`

GetCumulativeMiles returns the CumulativeMiles field if non-nil, zero value otherwise.

### GetCumulativeMilesOk

`func (o *TypesWorkOrder) GetCumulativeMilesOk() (*int32, bool)`

GetCumulativeMilesOk returns a tuple with the CumulativeMiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeMiles

`func (o *TypesWorkOrder) SetCumulativeMiles(v int32)`

SetCumulativeMiles sets CumulativeMiles field to given value.

### HasCumulativeMiles

`func (o *TypesWorkOrder) HasCumulativeMiles() bool`

HasCumulativeMiles returns a boolean if a field has been set.

### GetInstructions

`func (o *TypesWorkOrder) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *TypesWorkOrder) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *TypesWorkOrder) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *TypesWorkOrder) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetNotes

`func (o *TypesWorkOrder) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *TypesWorkOrder) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *TypesWorkOrder) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *TypesWorkOrder) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetStatus

`func (o *TypesWorkOrder) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TypesWorkOrder) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TypesWorkOrder) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTitle

`func (o *TypesWorkOrder) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TypesWorkOrder) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TypesWorkOrder) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


