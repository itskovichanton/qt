#ifdef __cplusplus
extern "C" {
#endif

void QGeoCodeReply_TimerEvent(void* ptr, void* event);
void QGeoCodeReply_TimerEventDefault(void* ptr, void* event);
void QGeoCodeReply_ChildEvent(void* ptr, void* event);
void QGeoCodeReply_ChildEventDefault(void* ptr, void* event);
void QGeoCodeReply_CustomEvent(void* ptr, void* event);
void QGeoCodeReply_CustomEventDefault(void* ptr, void* event);
void QGeoCodingManager_TimerEvent(void* ptr, void* event);
void QGeoCodingManager_TimerEventDefault(void* ptr, void* event);
void QGeoCodingManager_ChildEvent(void* ptr, void* event);
void QGeoCodingManager_ChildEventDefault(void* ptr, void* event);
void QGeoCodingManager_CustomEvent(void* ptr, void* event);
void QGeoCodingManager_CustomEventDefault(void* ptr, void* event);
void QGeoCodingManagerEngine_TimerEvent(void* ptr, void* event);
void QGeoCodingManagerEngine_TimerEventDefault(void* ptr, void* event);
void QGeoCodingManagerEngine_ChildEvent(void* ptr, void* event);
void QGeoCodingManagerEngine_ChildEventDefault(void* ptr, void* event);
void QGeoCodingManagerEngine_CustomEvent(void* ptr, void* event);
void QGeoCodingManagerEngine_CustomEventDefault(void* ptr, void* event);
void* QGeoManeuver_NewQGeoManeuver();
void* QGeoManeuver_NewQGeoManeuver2(void* other);
int QGeoManeuver_Direction(void* ptr);
double QGeoManeuver_DistanceToNextInstruction(void* ptr);
char* QGeoManeuver_InstructionText(void* ptr);
int QGeoManeuver_IsValid(void* ptr);
void QGeoManeuver_SetDirection(void* ptr, int direction);
void QGeoManeuver_SetDistanceToNextInstruction(void* ptr, double distance);
void QGeoManeuver_SetInstructionText(void* ptr, char* instructionText);
void QGeoManeuver_SetPosition(void* ptr, void* position);
void QGeoManeuver_SetTimeToNextInstruction(void* ptr, int secs);
void QGeoManeuver_SetWaypoint(void* ptr, void* coordinate);
int QGeoManeuver_TimeToNextInstruction(void* ptr);
void QGeoManeuver_DestroyQGeoManeuver(void* ptr);
void* QGeoRoute_NewQGeoRoute();
void* QGeoRoute_NewQGeoRoute2(void* other);
double QGeoRoute_Distance(void* ptr);
char* QGeoRoute_RouteId(void* ptr);
void QGeoRoute_SetBounds(void* ptr, void* bounds);
void QGeoRoute_SetDistance(void* ptr, double distance);
void QGeoRoute_SetFirstRouteSegment(void* ptr, void* routeSegment);
void QGeoRoute_SetRequest(void* ptr, void* request);
void QGeoRoute_SetRouteId(void* ptr, char* id);
void QGeoRoute_SetTravelMode(void* ptr, int mode);
void QGeoRoute_SetTravelTime(void* ptr, int secs);
int QGeoRoute_TravelMode(void* ptr);
int QGeoRoute_TravelTime(void* ptr);
void QGeoRoute_DestroyQGeoRoute(void* ptr);
void* QGeoRouteReply_NewQGeoRouteReply(int error, char* errorString, void* parent);
void QGeoRouteReply_Abort(void* ptr);
void QGeoRouteReply_AbortDefault(void* ptr);
void QGeoRouteReply_ConnectError2(void* ptr);
void QGeoRouteReply_DisconnectError2(void* ptr);
void QGeoRouteReply_Error2(void* ptr, int error, char* errorString);
int QGeoRouteReply_Error(void* ptr);
char* QGeoRouteReply_ErrorString(void* ptr);
void QGeoRouteReply_ConnectFinished(void* ptr);
void QGeoRouteReply_DisconnectFinished(void* ptr);
void QGeoRouteReply_Finished(void* ptr);
int QGeoRouteReply_IsFinished(void* ptr);
void QGeoRouteReply_DestroyQGeoRouteReply(void* ptr);
void QGeoRouteReply_TimerEvent(void* ptr, void* event);
void QGeoRouteReply_TimerEventDefault(void* ptr, void* event);
void QGeoRouteReply_ChildEvent(void* ptr, void* event);
void QGeoRouteReply_ChildEventDefault(void* ptr, void* event);
void QGeoRouteReply_CustomEvent(void* ptr, void* event);
void QGeoRouteReply_CustomEventDefault(void* ptr, void* event);
void* QGeoRouteRequest_NewQGeoRouteRequest2(void* origin, void* destination);
void* QGeoRouteRequest_NewQGeoRouteRequest3(void* other);
int QGeoRouteRequest_FeatureWeight(void* ptr, int featureType);
int QGeoRouteRequest_ManeuverDetail(void* ptr);
int QGeoRouteRequest_NumberAlternativeRoutes(void* ptr);
int QGeoRouteRequest_RouteOptimization(void* ptr);
int QGeoRouteRequest_SegmentDetail(void* ptr);
void QGeoRouteRequest_SetFeatureWeight(void* ptr, int featureType, int featureWeight);
void QGeoRouteRequest_SetManeuverDetail(void* ptr, int maneuverDetail);
void QGeoRouteRequest_SetNumberAlternativeRoutes(void* ptr, int alternatives);
void QGeoRouteRequest_SetRouteOptimization(void* ptr, int optimization);
void QGeoRouteRequest_SetSegmentDetail(void* ptr, int segmentDetail);
void QGeoRouteRequest_SetTravelModes(void* ptr, int travelModes);
int QGeoRouteRequest_TravelModes(void* ptr);
void QGeoRouteRequest_DestroyQGeoRouteRequest(void* ptr);
void* QGeoRouteSegment_NewQGeoRouteSegment();
void* QGeoRouteSegment_NewQGeoRouteSegment2(void* other);
double QGeoRouteSegment_Distance(void* ptr);
int QGeoRouteSegment_IsValid(void* ptr);
void QGeoRouteSegment_SetDistance(void* ptr, double distance);
void QGeoRouteSegment_SetManeuver(void* ptr, void* maneuver);
void QGeoRouteSegment_SetNextRouteSegment(void* ptr, void* routeSegment);
void QGeoRouteSegment_SetTravelTime(void* ptr, int secs);
int QGeoRouteSegment_TravelTime(void* ptr);
void QGeoRouteSegment_DestroyQGeoRouteSegment(void* ptr);
void* QGeoRoutingManager_CalculateRoute(void* ptr, void* request);
void QGeoRoutingManager_ConnectError(void* ptr);
void QGeoRoutingManager_DisconnectError(void* ptr);
void QGeoRoutingManager_Error(void* ptr, void* reply, int error, char* errorString);
void QGeoRoutingManager_ConnectFinished(void* ptr);
void QGeoRoutingManager_DisconnectFinished(void* ptr);
void QGeoRoutingManager_Finished(void* ptr, void* reply);
char* QGeoRoutingManager_ManagerName(void* ptr);
int QGeoRoutingManager_ManagerVersion(void* ptr);
int QGeoRoutingManager_MeasurementSystem(void* ptr);
void QGeoRoutingManager_SetLocale(void* ptr, void* locale);
void QGeoRoutingManager_SetMeasurementSystem(void* ptr, int system);
int QGeoRoutingManager_SupportedFeatureTypes(void* ptr);
int QGeoRoutingManager_SupportedFeatureWeights(void* ptr);
int QGeoRoutingManager_SupportedManeuverDetails(void* ptr);
int QGeoRoutingManager_SupportedRouteOptimizations(void* ptr);
int QGeoRoutingManager_SupportedSegmentDetails(void* ptr);
int QGeoRoutingManager_SupportedTravelModes(void* ptr);
void* QGeoRoutingManager_UpdateRoute(void* ptr, void* route, void* position);
void QGeoRoutingManager_DestroyQGeoRoutingManager(void* ptr);
void QGeoRoutingManager_TimerEvent(void* ptr, void* event);
void QGeoRoutingManager_TimerEventDefault(void* ptr, void* event);
void QGeoRoutingManager_ChildEvent(void* ptr, void* event);
void QGeoRoutingManager_ChildEventDefault(void* ptr, void* event);
void QGeoRoutingManager_CustomEvent(void* ptr, void* event);
void QGeoRoutingManager_CustomEventDefault(void* ptr, void* event);
void* QGeoRoutingManagerEngine_CalculateRoute(void* ptr, void* request);
void QGeoRoutingManagerEngine_ConnectError(void* ptr);
void QGeoRoutingManagerEngine_DisconnectError(void* ptr);
void QGeoRoutingManagerEngine_Error(void* ptr, void* reply, int error, char* errorString);
void QGeoRoutingManagerEngine_ConnectFinished(void* ptr);
void QGeoRoutingManagerEngine_DisconnectFinished(void* ptr);
void QGeoRoutingManagerEngine_Finished(void* ptr, void* reply);
char* QGeoRoutingManagerEngine_ManagerName(void* ptr);
int QGeoRoutingManagerEngine_ManagerVersion(void* ptr);
int QGeoRoutingManagerEngine_MeasurementSystem(void* ptr);
void QGeoRoutingManagerEngine_SetLocale(void* ptr, void* locale);
void QGeoRoutingManagerEngine_SetMeasurementSystem(void* ptr, int system);
int QGeoRoutingManagerEngine_SupportedFeatureTypes(void* ptr);
int QGeoRoutingManagerEngine_SupportedFeatureWeights(void* ptr);
int QGeoRoutingManagerEngine_SupportedManeuverDetails(void* ptr);
int QGeoRoutingManagerEngine_SupportedRouteOptimizations(void* ptr);
int QGeoRoutingManagerEngine_SupportedSegmentDetails(void* ptr);
int QGeoRoutingManagerEngine_SupportedTravelModes(void* ptr);
void* QGeoRoutingManagerEngine_UpdateRoute(void* ptr, void* route, void* position);
void* QGeoRoutingManagerEngine_UpdateRouteDefault(void* ptr, void* route, void* position);
void QGeoRoutingManagerEngine_DestroyQGeoRoutingManagerEngine(void* ptr);
void QGeoRoutingManagerEngine_TimerEvent(void* ptr, void* event);
void QGeoRoutingManagerEngine_TimerEventDefault(void* ptr, void* event);
void QGeoRoutingManagerEngine_ChildEvent(void* ptr, void* event);
void QGeoRoutingManagerEngine_ChildEventDefault(void* ptr, void* event);
void QGeoRoutingManagerEngine_CustomEvent(void* ptr, void* event);
void QGeoRoutingManagerEngine_CustomEventDefault(void* ptr, void* event);
int QGeoServiceProvider_OnlineGeocodingFeature_Type();
int QGeoServiceProvider_OfflineGeocodingFeature_Type();
int QGeoServiceProvider_ReverseGeocodingFeature_Type();
int QGeoServiceProvider_LocalizedGeocodingFeature_Type();
int QGeoServiceProvider_AnyGeocodingFeatures_Type();
int QGeoServiceProvider_OnlineMappingFeature_Type();
int QGeoServiceProvider_OfflineMappingFeature_Type();
int QGeoServiceProvider_LocalizedMappingFeature_Type();
int QGeoServiceProvider_AnyMappingFeatures_Type();
int QGeoServiceProvider_OnlinePlacesFeature_Type();
int QGeoServiceProvider_OfflinePlacesFeature_Type();
int QGeoServiceProvider_SavePlaceFeature_Type();
int QGeoServiceProvider_RemovePlaceFeature_Type();
int QGeoServiceProvider_SaveCategoryFeature_Type();
int QGeoServiceProvider_RemoveCategoryFeature_Type();
int QGeoServiceProvider_PlaceRecommendationsFeature_Type();
int QGeoServiceProvider_SearchSuggestionsFeature_Type();
int QGeoServiceProvider_LocalizedPlacesFeature_Type();
int QGeoServiceProvider_NotificationsFeature_Type();
int QGeoServiceProvider_PlaceMatchingFeature_Type();
int QGeoServiceProvider_AnyPlacesFeatures_Type();
int QGeoServiceProvider_OnlineRoutingFeature_Type();
int QGeoServiceProvider_OfflineRoutingFeature_Type();
int QGeoServiceProvider_LocalizedRoutingFeature_Type();
int QGeoServiceProvider_RouteUpdatesFeature_Type();
int QGeoServiceProvider_AlternativeRoutesFeature_Type();
int QGeoServiceProvider_ExcludeAreasRoutingFeature_Type();
int QGeoServiceProvider_AnyRoutingFeatures_Type();
char* QGeoServiceProvider_QGeoServiceProvider_AvailableServiceProviders();
int QGeoServiceProvider_Error(void* ptr);
char* QGeoServiceProvider_ErrorString(void* ptr);
int QGeoServiceProvider_GeocodingFeatures(void* ptr);
void* QGeoServiceProvider_GeocodingManager(void* ptr);
int QGeoServiceProvider_MappingFeatures(void* ptr);
void* QGeoServiceProvider_PlaceManager(void* ptr);
int QGeoServiceProvider_PlacesFeatures(void* ptr);
int QGeoServiceProvider_RoutingFeatures(void* ptr);
void* QGeoServiceProvider_RoutingManager(void* ptr);
void QGeoServiceProvider_SetAllowExperimental(void* ptr, int allow);
void QGeoServiceProvider_SetLocale(void* ptr, void* locale);
void QGeoServiceProvider_DestroyQGeoServiceProvider(void* ptr);
void QGeoServiceProvider_TimerEvent(void* ptr, void* event);
void QGeoServiceProvider_TimerEventDefault(void* ptr, void* event);
void QGeoServiceProvider_ChildEvent(void* ptr, void* event);
void QGeoServiceProvider_ChildEventDefault(void* ptr, void* event);
void QGeoServiceProvider_CustomEvent(void* ptr, void* event);
void QGeoServiceProvider_CustomEventDefault(void* ptr, void* event);
void QGeoServiceProviderFactory_DestroyQGeoServiceProviderFactory(void* ptr);
char* QGeoServiceProviderFactory_ObjectNameAbs(void* ptr);
void QGeoServiceProviderFactory_SetObjectNameAbs(void* ptr, char* name);
void QPlaceContentReply_TimerEvent(void* ptr, void* event);
void QPlaceContentReply_TimerEventDefault(void* ptr, void* event);
void QPlaceContentReply_ChildEvent(void* ptr, void* event);
void QPlaceContentReply_ChildEventDefault(void* ptr, void* event);
void QPlaceContentReply_CustomEvent(void* ptr, void* event);
void QPlaceContentReply_CustomEventDefault(void* ptr, void* event);
void QPlaceDetailsReply_TimerEvent(void* ptr, void* event);
void QPlaceDetailsReply_TimerEventDefault(void* ptr, void* event);
void QPlaceDetailsReply_ChildEvent(void* ptr, void* event);
void QPlaceDetailsReply_ChildEventDefault(void* ptr, void* event);
void QPlaceDetailsReply_CustomEvent(void* ptr, void* event);
void QPlaceDetailsReply_CustomEventDefault(void* ptr, void* event);
void QPlaceIdReply_TimerEvent(void* ptr, void* event);
void QPlaceIdReply_TimerEventDefault(void* ptr, void* event);
void QPlaceIdReply_ChildEvent(void* ptr, void* event);
void QPlaceIdReply_ChildEventDefault(void* ptr, void* event);
void QPlaceIdReply_CustomEvent(void* ptr, void* event);
void QPlaceIdReply_CustomEventDefault(void* ptr, void* event);
void QPlaceManager_TimerEvent(void* ptr, void* event);
void QPlaceManager_TimerEventDefault(void* ptr, void* event);
void QPlaceManager_ChildEvent(void* ptr, void* event);
void QPlaceManager_ChildEventDefault(void* ptr, void* event);
void QPlaceManager_CustomEvent(void* ptr, void* event);
void QPlaceManager_CustomEventDefault(void* ptr, void* event);
void QPlaceManagerEngine_TimerEvent(void* ptr, void* event);
void QPlaceManagerEngine_TimerEventDefault(void* ptr, void* event);
void QPlaceManagerEngine_ChildEvent(void* ptr, void* event);
void QPlaceManagerEngine_ChildEventDefault(void* ptr, void* event);
void QPlaceManagerEngine_CustomEvent(void* ptr, void* event);
void QPlaceManagerEngine_CustomEventDefault(void* ptr, void* event);
void QPlaceMatchReply_TimerEvent(void* ptr, void* event);
void QPlaceMatchReply_TimerEventDefault(void* ptr, void* event);
void QPlaceMatchReply_ChildEvent(void* ptr, void* event);
void QPlaceMatchReply_ChildEventDefault(void* ptr, void* event);
void QPlaceMatchReply_CustomEvent(void* ptr, void* event);
void QPlaceMatchReply_CustomEventDefault(void* ptr, void* event);
void QPlaceReply_TimerEvent(void* ptr, void* event);
void QPlaceReply_TimerEventDefault(void* ptr, void* event);
void QPlaceReply_ChildEvent(void* ptr, void* event);
void QPlaceReply_ChildEventDefault(void* ptr, void* event);
void QPlaceReply_CustomEvent(void* ptr, void* event);
void QPlaceReply_CustomEventDefault(void* ptr, void* event);
void QPlaceSearchReply_TimerEvent(void* ptr, void* event);
void QPlaceSearchReply_TimerEventDefault(void* ptr, void* event);
void QPlaceSearchReply_ChildEvent(void* ptr, void* event);
void QPlaceSearchReply_ChildEventDefault(void* ptr, void* event);
void QPlaceSearchReply_CustomEvent(void* ptr, void* event);
void QPlaceSearchReply_CustomEventDefault(void* ptr, void* event);
void QPlaceSearchSuggestionReply_TimerEvent(void* ptr, void* event);
void QPlaceSearchSuggestionReply_TimerEventDefault(void* ptr, void* event);
void QPlaceSearchSuggestionReply_ChildEvent(void* ptr, void* event);
void QPlaceSearchSuggestionReply_ChildEventDefault(void* ptr, void* event);
void QPlaceSearchSuggestionReply_CustomEvent(void* ptr, void* event);
void QPlaceSearchSuggestionReply_CustomEventDefault(void* ptr, void* event);

#ifdef __cplusplus
}
#endif