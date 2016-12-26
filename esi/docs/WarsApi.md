# \WarsApi

All URIs are relative to *https://esi.tech.ccp.is/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWars**](WarsApi.md#GetWars) | **Get** /wars/ | List wars
[**GetWarsWarId**](WarsApi.md#GetWarsWarId) | **Get** /wars/{war_id}/ | Get war information
[**GetWarsWarIdKillmails**](WarsApi.md#GetWarsWarIdKillmails) | **Get** /wars/{war_id}/killmails/ | List kills for a war


# **GetWars**
> []int32 GetWars(optional)
List wars

Return a list of wars  ---  Alternate route: `/v1/wars/`  Alternate route: `/legacy/wars/`  Alternate route: `/dev/wars/`   ---  This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32**| Which page to query, starting at 1, 2000 wars per page. | [default to 1]
 **datasource** | **string**| The server name you would like data from | [default to tranquility]

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWarsWarId**
> GetWarsWarIdOk GetWarsWarId(warId, optional)
Get war information

Return details about a war  ---  Alternate route: `/v1/wars/{war_id}/`  Alternate route: `/legacy/wars/{war_id}/`  Alternate route: `/dev/wars/{war_id}/`   ---  This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **warId** | **int32**| ID for a war | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **warId** | **int32**| ID for a war | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]

### Return type

[**GetWarsWarIdOk**](get_wars_war_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWarsWarIdKillmails**
> []GetWarsWarIdKillmails200Ok GetWarsWarIdKillmails(warId, optional)
List kills for a war

Return a list of kills related to a war  ---  Alternate route: `/v1/wars/{war_id}/killmails/`  Alternate route: `/legacy/wars/{war_id}/killmails/`  Alternate route: `/dev/wars/{war_id}/killmails/`   ---  This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **warId** | **int32**| A valid war ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **warId** | **int32**| A valid war ID | 
 **page** | **int32**| Which page to query, starting at 1, 2000 killmails per page. | [default to 1]
 **datasource** | **string**| The server name you would like data from | [default to tranquility]

### Return type

[**[]GetWarsWarIdKillmails200Ok**](get_wars_war_id_killmails_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

