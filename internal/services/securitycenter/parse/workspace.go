package parse

// NOTE: this file is generated via 'go:generate' - manual changes will be overwritten

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-provider-azurerm/helpers/azure"
)

type WorkspaceId struct {
	SubscriptionId       string
	WorkspaceSettingName string
}

func NewWorkspaceID(subscriptionId, workspaceSettingName string) WorkspaceId {
	return WorkspaceId{
		SubscriptionId:       subscriptionId,
		WorkspaceSettingName: workspaceSettingName,
	}
}

func (id WorkspaceId) String() string {
	segments := []string{
		fmt.Sprintf("Workspace Setting Name %q", id.WorkspaceSettingName),
	}
	segmentsStr := strings.Join(segments, " / ")
	return fmt.Sprintf("%s: (%s)", "Workspace", segmentsStr)
}

func (id WorkspaceId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Security/workspaceSettings/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.WorkspaceSettingName)
}

// WorkspaceID parses a Workspace ID into an WorkspaceId struct
func WorkspaceID(input string) (*WorkspaceId, error) {
	id, err := azure.ParseAzureResourceID(input)
	if err != nil {
		return nil, err
	}

	resourceId := WorkspaceId{
		SubscriptionId: id.SubscriptionID,
	}

	if resourceId.SubscriptionId == "" {
		return nil, fmt.Errorf("ID was missing the 'subscriptions' element")
	}

	if resourceId.WorkspaceSettingName, err = id.PopSegment("workspaceSettings"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &resourceId, nil
}
