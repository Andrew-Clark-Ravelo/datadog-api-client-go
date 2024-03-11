// Update a CSM Threats Agent rule returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "agent_rule_rc" in the system
	AgentRuleDataID := os.Getenv("AGENT_RULE_DATA_ID")

	body := datadogV2.CloudWorkloadSecurityAgentRuleUpdateRequest{
		Data: datadogV2.CloudWorkloadSecurityAgentRuleUpdateData{
			Attributes: datadogV2.CloudWorkloadSecurityAgentRuleUpdateAttributes{
				Description: datadog.PtrString("Test Agent rule"),
				Enabled:     datadog.PtrBool(true),
				Expression:  datadog.PtrString(`exec.file.name == "sh"`),
			},
			Type: datadogV2.CLOUDWORKLOADSECURITYAGENTRULETYPE_AGENT_RULE,
			Id:   datadog.PtrString(AgentRuleDataID),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudWorkloadSecurityApi(apiClient)
	resp, r, err := api.UpdateCSMThreatsAgentRule(ctx, AgentRuleDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudWorkloadSecurityApi.UpdateCSMThreatsAgentRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudWorkloadSecurityApi.UpdateCSMThreatsAgentRule`:\n%s\n", responseContent)
}
