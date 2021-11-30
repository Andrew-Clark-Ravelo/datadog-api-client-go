@endpoint(logs) @endpoint(logs-v2)
Feature: Logs
  Search your logs and send them to your Datadog platform over HTTP.

  Background:
    Given a valid "apiKeyAuth" key in the system
    And an instance of "Logs" API

  Scenario: Aggregate compute events returns "OK" response
    Given a valid "appKeyAuth" key in the system
    And new "AggregateLogs" request
    And body with value {"compute": [{"aggregation": "count", "interval": "5m", "type": "timeseries"}], "filter": {"from": "now-15m", "indexes": ["main"], "query": "*", "to": "now"}}
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Aggregate events returns "Bad Request" response
    Given a valid "appKeyAuth" key in the system
    And new "AggregateLogs" request
    And body with value {"compute": [{"aggregation": "pc90", "interval": "5m", "metric": "@duration", "type": "total"}], "filter": {"from": "now-15m", "indexes": ["main", "web"], "query": "service:web* AND @http.status_code:[200 TO 299]", "to": "now"}, "group_by": [{"facet": "host", "histogram": {"interval": 10, "max": 100, "min": 50}, "limit": 10, "missing": null, "sort": {"aggregation": "count", "order": "asc"}, "total": false}], "options": {"timeOffset": null, "timezone": "GMT"}, "page": {"cursor": "eyJzdGFydEF0IjoiQVFBQUFYS2tMS3pPbm40NGV3QUFBQUJCV0V0clRFdDZVbG8zY3pCRmNsbHJiVmxDWlEifQ=="}}
    When the request is sent
    Then the response status is 400 Bad Request

  Scenario: Aggregate events returns "OK" response
    Given a valid "appKeyAuth" key in the system
    And new "AggregateLogs" request
    And body with value {"filter": {"from": "now-15m", "indexes": ["main"], "query": "*", "to": "now"}}
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Get a list of logs returns "Bad Request" response
    Given a valid "appKeyAuth" key in the system
    And new "ListLogsGet" request
    When the request is sent
    Then the response status is 400 Bad Request

  Scenario: Search logs returns "OK" response
    Given a valid "appKeyAuth" key in the system
    And operation "ListLogs" enabled
    And new "ListLogs" request
    And body with value {"filter": {"query": "datadog-agent", "indexes": ["main"], "from": "2020-09-17T11:48:36+01:00", "to": "2020-09-17T12:48:36+01:00"}, "sort": "timestamp", "page": {"limit": 5}}
    When the request is sent
    Then the response status is 200 OK

  Scenario: Get a quick list of logs returns "OK" response
    Given a valid "appKeyAuth" key in the system
    And operation "ListLogsGet" enabled
    And new "ListLogsGet" request
    And request contains "filter[query]" parameter with value "datadog-agent"
    And request contains "filter[index]" parameter with value "main"
    And request contains "filter[from]" parameter with value "2020-09-17T11:48:36+01:00"
    And request contains "filter[to]" parameter with value "2020-09-17T12:48:36+01:00"
    And request contains "page[limit]" parameter with value 5
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Search logs returns "Bad Request" response
    Given a valid "appKeyAuth" key in the system
    And new "ListLogs" request
    And body with value {"filter": {"from": "now-15m", "indexes": ["main", "web"], "query": "service:web* AND @http.status_code:[200 TO 299]", "to": "now"}, "options": {"timeOffset": null, "timezone": "GMT"}, "page": {"cursor": "eyJzdGFydEF0IjoiQVFBQUFYS2tMS3pPbm40NGV3QUFBQUJCV0V0clRFdDZVbG8zY3pCRmNsbHJiVmxDWlEifQ==", "limit": 25}, "sort": "timestamp"}
    When the request is sent
    Then the response status is 400 Bad Request

  @integration-only
  Scenario: Send deflat logs returns "Request accepted for processing (always 202 empty JSON)." response
    Given new "SubmitLog" request
    And body with value [{"ddsource": "nginx", "ddtags": "env:staging,version:5.1", "hostname": "i-012345678", "message": "2019-11-19T14:37:58,995 INFO [process.name][20081] Hello World", "service": "payment"}]
    And request contains "content_encoding" parameter with value "deflate"
    When the request is sent
    Then the response status is 202 Response from server (always 202 empty JSON).

  @integration-only
  Scenario: Send gzip logs returns "Request accepted for processing (always 202 empty JSON)." response
    Given new "SubmitLog" request
    And body with value [{"ddsource": "nginx", "ddtags": "env:staging,version:5.1", "hostname": "i-012345678", "message": "2019-11-19T14:37:58,995 INFO [process.name][20081] Hello World", "service": "payment"}]
    And request contains "content_encoding" parameter with value "gzip"
    When the request is sent
    Then the response status is 202 Request accepted for processing (always 202 empty JSON).

  @generated @skip
  Scenario: Send logs returns "Bad Request" response
    Given new "SubmitLog" request
    And body with value [{"ddsource": "nginx", "ddtags": "env:staging,version:5.1", "hostname": "i-012345678", "message": "2019-11-19T14:37:58,995 INFO [process.name][20081] Hello World", "service": "payment"}]
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip
  Scenario: Send logs returns "Payload Too Large" response
    Given new "SubmitLog" request
    And body with value [{"ddsource": "nginx", "ddtags": "env:staging,version:5.1", "hostname": "i-012345678", "message": "2019-11-19T14:37:58,995 INFO [process.name][20081] Hello World", "service": "payment"}]
    When the request is sent
    Then the response status is 413 Payload Too Large

  @generated @skip
  Scenario: Send logs returns "Request Timeout" response
    Given new "SubmitLog" request
    And body with value [{"ddsource": "nginx", "ddtags": "env:staging,version:5.1", "hostname": "i-012345678", "message": "2019-11-19T14:37:58,995 INFO [process.name][20081] Hello World", "service": "payment"}]
    When the request is sent
    Then the response status is 408 Request Timeout

  @generated @skip
  Scenario: Send logs returns "Request accepted for processing (always 202 empty JSON)." response
    Given new "SubmitLog" request
    And body with value [{"ddsource": "nginx", "ddtags": "env:staging,version:5.1", "hostname": "i-012345678", "message": "2019-11-19T14:37:58,995 INFO [process.name][20081] Hello World", "service": "payment"}]
    When the request is sent
    Then the response status is 202 Request accepted for processing (always 202 empty JSON).