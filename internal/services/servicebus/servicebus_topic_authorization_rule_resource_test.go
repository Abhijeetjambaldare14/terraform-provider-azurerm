package servicebus_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/servicebus/parse"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/utils"
)

type ServiceBusTopicAuthorizationRuleResource struct {
}

func TestAccServiceBusTopicAuthorizationRule_listen(t *testing.T) {
	testAccServiceBusTopicAuthorizationRule(t, true, false, false)
}

func TestAccServiceBusTopicAuthorizationRule_send(t *testing.T) {
	testAccServiceBusTopicAuthorizationRule(t, false, true, false)
}

func TestAccServiceBusTopicAuthorizationRule_listensend(t *testing.T) {
	testAccServiceBusTopicAuthorizationRule(t, true, true, false)
}

func TestAccServiceBusTopicAuthorizationRule_manage(t *testing.T) {
	testAccServiceBusTopicAuthorizationRule(t, true, true, true)
}

func testAccServiceBusTopicAuthorizationRule(t *testing.T, listen, send, manage bool) {
	data := acceptance.BuildTestData(t, "azurerm_servicebus_topic_authorization_rule", "test")
	r := ServiceBusTopicAuthorizationRuleResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.base(data, listen, send, manage),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("name").Exists(),
				check.That(data.ResourceName).Key("namespace_name").Exists(),
				check.That(data.ResourceName).Key("primary_key").Exists(),
				check.That(data.ResourceName).Key("secondary_key").Exists(),
				check.That(data.ResourceName).Key("primary_connection_string").Exists(),
				check.That(data.ResourceName).Key("secondary_connection_string").Exists(),
				check.That(data.ResourceName).Key("primary_connection_string_alias").HasValue(""),
				check.That(data.ResourceName).Key("secondary_connection_string_alias").HasValue(""),
				check.That(data.ResourceName).Key("listen").HasValue(strconv.FormatBool(listen)),
				check.That(data.ResourceName).Key("send").HasValue(strconv.FormatBool(send)),
				check.That(data.ResourceName).Key("manage").HasValue(strconv.FormatBool(manage)),
			),
		},
		data.ImportStep(),
	})
}

func TestAccServiceBusTopicAuthorizationRule_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_servicebus_topic_authorization_rule", "test")
	r := ServiceBusTopicAuthorizationRuleResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.base(data, true, false, false),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("listen").HasValue("true"),
				check.That(data.ResourceName).Key("send").HasValue("false"),
				check.That(data.ResourceName).Key("manage").HasValue("false"),
			),
		},
		{
			Config:      r.requiresImport(data, true, false, false),
			ExpectError: acceptance.RequiresImportError("azurerm_servicebus_topic_authorization_rule"),
		},
	})
}

func TestAccServiceBusTopicAuthorizationRule_rightsUpdate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_servicebus_topic_authorization_rule", "test")
	r := ServiceBusTopicAuthorizationRuleResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.base(data, true, false, false),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("listen").HasValue("true"),
				check.That(data.ResourceName).Key("send").HasValue("false"),
				check.That(data.ResourceName).Key("manage").HasValue("false"),
			),
		},
		{
			Config: r.base(data, true, true, true),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("name").Exists(),
				check.That(data.ResourceName).Key("namespace_name").Exists(),
				check.That(data.ResourceName).Key("primary_key").Exists(),
				check.That(data.ResourceName).Key("secondary_key").Exists(),
				check.That(data.ResourceName).Key("primary_connection_string").Exists(),
				check.That(data.ResourceName).Key("secondary_connection_string").Exists(),
				check.That(data.ResourceName).Key("primary_connection_string_alias").HasValue(""),
				check.That(data.ResourceName).Key("secondary_connection_string_alias").HasValue(""),
				check.That(data.ResourceName).Key("listen").HasValue("true"),
				check.That(data.ResourceName).Key("send").HasValue("true"),
				check.That(data.ResourceName).Key("manage").HasValue("true"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccServiceBusTopicAuthorizationRule_withAliasConnectionString(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_servicebus_topic_authorization_rule", "test")
	r := ServiceBusTopicAuthorizationRuleResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.withAliasConnectionString(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		{
			Config: r.withAliasConnectionString(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("primary_connection_string_alias").Exists(),
				check.That(data.ResourceName).Key("secondary_connection_string_alias").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func (t ServiceBusTopicAuthorizationRuleResource) Exists(ctx context.Context, clients *clients.Client, state *pluginsdk.InstanceState) (*bool, error) {
	id, err := parse.TopicAuthorizationRuleID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := clients.ServiceBus.TopicsClient.GetAuthorizationRule(ctx, id.ResourceGroup, id.NamespaceName, id.TopicName, id.AuthorizationRuleName)
	if err != nil {
		return nil, fmt.Errorf("reading Service Bus Topic Authorization Rule (%s): %+v", id.String(), err)
	}

	return utils.Bool(resp.ID != nil), nil
}

func (ServiceBusTopicAuthorizationRuleResource) base(data acceptance.TestData, listen, send, manage bool) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%[1]d"
  location = "%[2]s"
}

resource "azurerm_servicebus_namespace" "test" {
  name                = "acctest-%[1]d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  sku                 = "Standard"
}

resource "azurerm_servicebus_topic" "test" {
  name                = "acctestservicebustopic-%[1]d"
  namespace_name      = azurerm_servicebus_namespace.test.name
  resource_group_name = azurerm_resource_group.test.name
}

resource "azurerm_servicebus_topic_authorization_rule" "test" {
  name                = "acctest-%[1]d"
  namespace_name      = azurerm_servicebus_namespace.test.name
  resource_group_name = azurerm_resource_group.test.name
  topic_name          = azurerm_servicebus_topic.test.name

  listen = %[3]t
  send   = %[4]t
  manage = %[5]t
}
`, data.RandomInteger, data.Locations.Primary, listen, send, manage)
}

func (r ServiceBusTopicAuthorizationRuleResource) requiresImport(data acceptance.TestData, listen, send, manage bool) string {
	return fmt.Sprintf(`
%s

resource "azurerm_servicebus_topic_authorization_rule" "import" {
  name                = azurerm_servicebus_topic_authorization_rule.test.name
  namespace_name      = azurerm_servicebus_topic_authorization_rule.test.namespace_name
  resource_group_name = azurerm_servicebus_topic_authorization_rule.test.resource_group_name
  topic_name          = azurerm_servicebus_topic_authorization_rule.test.topic_name

  listen = azurerm_servicebus_topic_authorization_rule.test.listen
  send   = azurerm_servicebus_topic_authorization_rule.test.send
  manage = azurerm_servicebus_topic_authorization_rule.test.manage
}
`, r.base(data, listen, send, manage))
}

func (ServiceBusTopicAuthorizationRuleResource) withAliasConnectionString(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "primary" {
  name     = "acctest1RG-%[1]d"
  location = "%[2]s"
}

resource "azurerm_resource_group" "secondary" {
  name     = "acctest2RG-%[1]d"
  location = "%[3]s"
}

resource "azurerm_servicebus_namespace" "primary_namespace_test" {
  name                = "acctest1-%[1]d"
  location            = azurerm_resource_group.primary.location
  resource_group_name = azurerm_resource_group.primary.name
  sku                 = "Premium"
  capacity            = "1"
}

resource "azurerm_servicebus_topic" "example" {
  name                = "topic-test"
  resource_group_name = azurerm_resource_group.primary.name
  namespace_name      = azurerm_servicebus_namespace.primary_namespace_test.name
}

resource "azurerm_servicebus_namespace" "secondary_namespace_test" {
  name                = "acctest2-%[1]d"
  location            = azurerm_resource_group.secondary.location
  resource_group_name = azurerm_resource_group.secondary.name
  sku                 = "Premium"
  capacity            = "1"
}

resource "azurerm_servicebus_namespace_disaster_recovery_config" "pairing_test" {
  name                 = "acctest-alias-%[1]d"
  primary_namespace_id = azurerm_servicebus_namespace.primary_namespace_test.id
  partner_namespace_id = azurerm_servicebus_namespace.secondary_namespace_test.id
}

resource "azurerm_servicebus_topic_authorization_rule" "test" {
  name                = "example_topic_rule"
  namespace_name      = azurerm_servicebus_namespace.primary_namespace_test.name
  topic_name          = azurerm_servicebus_topic.example.name
  resource_group_name = azurerm_resource_group.primary.name
  manage              = true
  listen              = true
  send                = true

  depends_on = [
    azurerm_servicebus_namespace_disaster_recovery_config.pairing_test
  ]
}
`, data.RandomInteger, data.Locations.Primary, data.Locations.Secondary)
}
