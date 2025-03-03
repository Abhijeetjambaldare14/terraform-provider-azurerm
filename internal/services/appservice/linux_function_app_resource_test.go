package appservice_test

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/appservice/parse"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/utils"
)

type LinuxFunctionAppResource struct{}

const (
	SkuConsumptionPlan    = "Y1"
	SkuElasticPremiumPlan = "EP1"
	SkuStandardPlan       = "S1"
	SkuBasicPlan          = "B1"
	SkuPremiumPlan        = "P1v2"
)

// Plan types
func TestAccLinuxFunctionApp_basicBasicPlan(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data, SkuBasicPlan),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccLinuxFunctionApp_basicConsumptionPlan(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data, SkuConsumptionPlan),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccLinuxFunctionApp_basicElasticPremiumPlan(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data, SkuElasticPremiumPlan),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccLinuxFunctionApp_basicPremiumAppServicePlan(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data, SkuPremiumPlan),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccLinuxFunctionApp_basicStandardPlan(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data, SkuStandardPlan),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
	})
}

// App Settings by Plan Type

func TestAccLinuxFunctionApp_withAppSettingsBasic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.appSettings(data, SkuBasicPlan),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
				check.That(data.ResourceName).Key("app_settings.%").HasValue("2"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccLinuxFunctionApp_withAppSettingsConsumption(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.appSettings(data, SkuConsumptionPlan),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
				check.That(data.ResourceName).Key("app_settings.%").HasValue("2"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccLinuxFunctionApp_withAppSettingsElasticPremiumPlan(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.appSettings(data, SkuElasticPremiumPlan),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
				check.That(data.ResourceName).Key("app_settings.%").HasValue("2"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccLinuxFunctionApp_withAppSettingsPremiumPlan(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.appSettings(data, SkuPremiumPlan),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
				check.That(data.ResourceName).Key("app_settings.%").HasValue("2"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccLinuxFunctionApp_withAppSettingsStandardPlan(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.appSettings(data, SkuStandardPlan),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
				check.That(data.ResourceName).Key("app_settings.%").HasValue("2"),
			),
		},
		data.ImportStep(),
	})
}

// backup by plan type

func TestAccLinuxFunctionApp_withBackupElasticPremiumPlan(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.backup(data, SkuElasticPremiumPlan),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccLinuxFunctionApp_withBackupPremiumPlan(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.backup(data, SkuPremiumPlan),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccLinuxFunctionApp_withBackupStandardPlan(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.backup(data, SkuStandardPlan),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
	})
}

// Completes by plan type

func TestAccLinuxFunctionApp_consumptionComplete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.consumptionComplete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccLinuxFunctionApp_consumptionCompleteUpdate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data, SkuConsumptionPlan),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
		{
			Config: r.consumptionComplete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data, SkuConsumptionPlan),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
	})
}

func TestAccLinuxFunctionApp_elasticPremiumComplete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.elasticComplete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccLinuxFunctionApp_standardComplete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.standardComplete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

// Individual Settings / Blocks

func TestAccLinuxFunctionApp_withAuthSettingsStandard(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.withAuthSettings(data, SkuStandardPlan),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccLinuxFunctionApp_withAuthSettingsConsumption(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.withAuthSettings(data, SkuConsumptionPlan),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccLinuxFunctionApp_builtInLogging(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.builtInLogging(data, SkuStandardPlan, true),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccLinuxFunctionApp_withConnectionStrings(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.builtInLogging(data, SkuStandardPlan, true),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccLinuxFunctionApp_withUserIdentity(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.userIdentity(data, SkuStandardPlan),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccLinuxFunctionApp_withConnectionStringsUpdate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data, SkuStandardPlan),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
		{
			Config: r.connectionStrings(data, SkuStandardPlan),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
		{
			Config: r.connectionStringsUpdate(data, SkuStandardPlan),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data, SkuStandardPlan),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccLinuxFunctionApp_dailyTimeQuotaConsumptionPlan(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.dailyTimeLimitQuota(data, SkuConsumptionPlan, 1000),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccLinuxFunctionApp_dailyTimeQuotaElasticPremiumPlan(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.dailyTimeLimitQuota(data, SkuElasticPremiumPlan, 2000),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccLinuxFunctionApp_healthCheckPath(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.healthCheckPath(data, "S1"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccLinuxFunctionApp_healthCheckPathWithEviction(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.healthCheckPathWithEviction(data, "S1"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccLinuxFunctionApp_healthCheckPathWithEvictionUpdate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data, "S1"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.healthCheckPathWithEviction(data, "S1"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data, "S1"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccLinuxFunctionApp_appServiceLogging(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.appServiceLogs(data, SkuStandardPlan),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccLinuxFunctionApp_appServiceLoggingUpdate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data, SkuStandardPlan),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
		{
			Config: r.appServiceLogs(data, SkuStandardPlan),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data, SkuStandardPlan),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
	})
}

// App Stacks

func TestAccLinuxFunctionApp_appStackDotNet31(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.appStackDotNet(data, SkuBasicPlan, "3.1"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccLinuxFunctionApp_appStackDotNet6(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.appStackDotNet(data, SkuBasicPlan, "6"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccLinuxFunctionApp_appStackPython(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.appStackPython(data, SkuBasicPlan, "3.7"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccLinuxFunctionApp_appStackPythonUpdate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.appStackPython(data, SkuBasicPlan, "3.7"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
		{
			Config: r.appStackPython(data, SkuBasicPlan, "3.9"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccLinuxFunctionApp_appStackNode(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.appStackNode(data, SkuBasicPlan, "14"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccLinuxFunctionApp_appStackNodeUpdate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.appStackNode(data, SkuBasicPlan, "12"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
		{
			Config: r.appStackNode(data, SkuBasicPlan, "14"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccLinuxFunctionApp_appStackJava(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.appStackJava(data, SkuBasicPlan, "11"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccLinuxFunctionApp_appStackJavaUpdate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.appStackJava(data, SkuBasicPlan, "8"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(), {
			Config: r.appStackJava(data, SkuBasicPlan, "11"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccLinuxFunctionApp_appStackDocker(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.appStackDocker(data, SkuBasicPlan),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux,container"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccLinuxFunctionApp_appStackDockerManagedServiceIdentity(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.appStackDockerUseMSI(data, SkuBasicPlan),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux,container"),
			),
		},
		data.ImportStep(),
	})
}

// Others

func TestAccLinuxFunctionApp_updateServicePlan(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data, SkuStandardPlan),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
		{
			Config: r.servicePlanUpdate(data, SkuStandardPlan),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccLinuxFunctionApp_updateStorageAccount(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data, SkuStandardPlan),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
		{
			Config: r.updateStorageAccount(data, SkuStandardPlan),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("functionapp,linux"),
			),
		},
		data.ImportStep(),
	})
}

// CustomDiff tests
func TestAccLinuxFunctionApp_consumptionPlanBackupShouldError(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config:      r.backup(data, SkuConsumptionPlan),
			ExpectError: regexp.MustCompile("cannot specify backup configuration for Dynamic tier Service Plans"),
		},
	})
}

func TestAccLinuxFunctionApp_basicPlanBackupShouldError(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_linux_function_app", "test")
	r := LinuxFunctionAppResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config:      r.backup(data, SkuBasicPlan),
			ExpectError: regexp.MustCompile("cannot specify backup configuration for Basic tier Service Plans"),
		},
	})
}

// Configs

func (r LinuxFunctionAppResource) Exists(ctx context.Context, client *clients.Client, state *pluginsdk.InstanceState) (*bool, error) {
	id, err := parse.FunctionAppID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := client.AppService.WebAppsClient.Get(ctx, id.ResourceGroup, id.SiteName)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			return utils.Bool(false), nil
		}
		return nil, fmt.Errorf("retrieving Linux Web App %s: %+v", id, err)
	}
	if utils.ResponseWasNotFound(resp.Response) {
		return utils.Bool(false), nil
	}
	return utils.Bool(true), nil
}

func (r LinuxFunctionAppResource) basic(data acceptance.TestData, planSku string) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_linux_function_app" "test" {
  name                = "acctest-FA-%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  service_plan_id     = azurerm_service_plan.test.id

  storage_account_name       = azurerm_storage_account.test.name
  storage_account_access_key = azurerm_storage_account.test.primary_access_key

  site_config {}
}
`, r.template(data, planSku), data.RandomInteger)
}

func (r LinuxFunctionAppResource) healthCheckPath(data acceptance.TestData, planSku string) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_linux_function_app" "test" {
  name                = "acctest-FA-%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  service_plan_id     = azurerm_service_plan.test.id

  storage_account_name       = azurerm_storage_account.test.name
  storage_account_access_key = azurerm_storage_account.test.primary_access_key

  site_config {
    health_check_path = "/health"
  }
}
`, r.template(data, planSku), data.RandomInteger)
}

func (r LinuxFunctionAppResource) healthCheckPathWithEviction(data acceptance.TestData, planSku string) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_linux_function_app" "test" {
  name                = "acctest-FA-%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  service_plan_id     = azurerm_service_plan.test.id

  storage_account_name       = azurerm_storage_account.test.name
  storage_account_access_key = azurerm_storage_account.test.primary_access_key

  site_config {
    health_check_path                 = "/health"
    health_check_eviction_time_in_min = 3
  }
}
`, r.template(data, planSku), data.RandomInteger)
}

func (r LinuxFunctionAppResource) userIdentity(data acceptance.TestData, planSku string) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_linux_function_app" "test" {
  name                = "acctest-FA-%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  service_plan_id     = azurerm_service_plan.test.id

  storage_account_name       = azurerm_storage_account.test.name
  storage_account_access_key = azurerm_storage_account.test.primary_access_key

  site_config {}

  identity {
    type         = "UserAssigned"
    identity_ids = [azurerm_user_assigned_identity.test.id]
  }
}
`, r.identityTemplate(data, planSku), data.RandomInteger)
}

func (r LinuxFunctionAppResource) builtInLogging(data acceptance.TestData, planSku string, builtInLogging bool) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_linux_function_app" "test" {
  name                = "acctest-FA-%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  service_plan_id     = azurerm_service_plan.test.id

  storage_account_name       = azurerm_storage_account.test.name
  storage_account_access_key = azurerm_storage_account.test.primary_access_key

  builtin_logging_enabled = %t

  site_config {}
}
`, r.template(data, planSku), data.RandomInteger, builtInLogging)
}

func (r LinuxFunctionAppResource) appSettings(data acceptance.TestData, planSku string) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_linux_function_app" "test" {
  name                = "acctest-FA-%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  service_plan_id     = azurerm_service_plan.test.id

  storage_account_name       = azurerm_storage_account.test.name
  storage_account_access_key = azurerm_storage_account.test.primary_access_key

  app_settings = {
    foo    = "bar"
    secret = "sauce"
  }

  site_config {}
}
`, r.template(data, planSku), data.RandomInteger)
}

func (r LinuxFunctionAppResource) connectionStrings(data acceptance.TestData, planSku string) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_linux_function_app" "test" {
  name                = "acctest-FA-%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  service_plan_id     = azurerm_service_plan.test.id

  storage_account_name       = azurerm_storage_account.test.name
  storage_account_access_key = azurerm_storage_account.test.primary_access_key

  connection_string {
    name  = "Example"
    value = "some-postgresql-connection-string"
    type  = "PostgreSQL"
  }

  site_config {}
}
`, r.template(data, planSku), data.RandomInteger)
}

func (r LinuxFunctionAppResource) dailyTimeLimitQuota(data acceptance.TestData, planSku string, quota int) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_linux_function_app" "test" {
  name                = "acctest-FA-%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  service_plan_id     = azurerm_service_plan.test.id

  storage_account_name       = azurerm_storage_account.test.name
  storage_account_access_key = azurerm_storage_account.test.primary_access_key

  daily_memory_time_quota = %d

  site_config {}
}
`, r.template(data, planSku), data.RandomInteger, quota)
}

func (r LinuxFunctionAppResource) withAuthSettings(data acceptance.TestData, planSku string) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_linux_function_app" "test" {
  name                = "acctest-FA-%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  service_plan_id     = azurerm_service_plan.test.id

  storage_account_name       = azurerm_storage_account.test.name
  storage_account_access_key = azurerm_storage_account.test.primary_access_key

  auth_settings {
    enabled                       = true
    issuer                        = "https://sts.windows.net/%s"
    runtime_version               = "1.0"
    unauthenticated_client_action = "RedirectToLoginPage"
    token_refresh_extension_hours = 75
    token_store_enabled           = true

    additional_login_parameters = {
      test_key = "test_value"
    }

    allowed_external_redirect_urls = [
      "https://terra.form",
    ]

    active_directory {
      client_id     = "aadclientid"
      client_secret = "aadsecret"

      allowed_audiences = [
        "activedirectorytokenaudiences",
      ]
    }
  }

  site_config {}
}
`, r.template(data, planSku), data.RandomInteger, data.RandomString)
}

func (r LinuxFunctionAppResource) connectionStringsUpdate(data acceptance.TestData, planSku string) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_linux_function_app" "test" {
  name                = "acctest-FA-%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  service_plan_id     = azurerm_service_plan.test.id

  storage_account_name       = azurerm_storage_account.test.name
  storage_account_access_key = azurerm_storage_account.test.primary_access_key

  connection_string {
    name  = "Example"
    value = "some-postgresql-connection-string"
    type  = "PostgreSQL"
  }

  connection_string {
    name  = "AnotherExample"
    value = "some-other-connection-string"
    type  = "Custom"
  }

  site_config {}
}
`, r.template(data, planSku), data.RandomInteger)
}

func (r LinuxFunctionAppResource) appServiceLogs(data acceptance.TestData, planSku string) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_linux_function_app" "test" {
  name                = "acctest-FA-%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  service_plan_id     = azurerm_service_plan.test.id

  storage_account_name       = azurerm_storage_account.test.name
  storage_account_access_key = azurerm_storage_account.test.primary_access_key

  site_config {
    app_service_logs {
      disk_quota_mb         = 25
      retention_period_days = 7
    }
  }
}
`, r.template(data, planSku), data.RandomInteger)
}

func (r LinuxFunctionAppResource) appStackDotNet(data acceptance.TestData, planSku string, version string) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_linux_function_app" "test" {
  name                = "acctest-FA-%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  service_plan_id     = azurerm_service_plan.test.id

  storage_account_name       = azurerm_storage_account.test.name
  storage_account_access_key = azurerm_storage_account.test.primary_access_key

  site_config {
    application_stack {
      dotnet_version = "%s"
    }
  }
}
`, r.template(data, planSku), data.RandomInteger, version)
}

func (r LinuxFunctionAppResource) appStackPython(data acceptance.TestData, planSku string, pythonVersion string) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_linux_function_app" "test" {
  name                = "acctest-FA-%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  service_plan_id     = azurerm_service_plan.test.id

  storage_account_name       = azurerm_storage_account.test.name
  storage_account_access_key = azurerm_storage_account.test.primary_access_key

  site_config {
    application_stack {
      python_version = "%s"
    }
  }
}
`, r.template(data, planSku), data.RandomInteger, pythonVersion)
}

func (r LinuxFunctionAppResource) appStackJava(data acceptance.TestData, planSku string, javaVersion string) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_linux_function_app" "test" {
  name                = "acctest-FA-%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  service_plan_id     = azurerm_service_plan.test.id

  storage_account_name       = azurerm_storage_account.test.name
  storage_account_access_key = azurerm_storage_account.test.primary_access_key

  site_config {
    application_stack {
      java_version = "%s"
    }
  }
}
`, r.template(data, planSku), data.RandomInteger, javaVersion)
}

func (r LinuxFunctionAppResource) appStackNode(data acceptance.TestData, planSku string, nodeVersion string) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_linux_function_app" "test" {
  name                = "acctest-FA-%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  service_plan_id     = azurerm_service_plan.test.id

  storage_account_name       = azurerm_storage_account.test.name
  storage_account_access_key = azurerm_storage_account.test.primary_access_key

  site_config {
    application_stack {
      node_version = "%s"
    }
  }
}
`, r.template(data, planSku), data.RandomInteger, nodeVersion)
}

func (r LinuxFunctionAppResource) appStackDocker(data acceptance.TestData, planSku string) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_linux_function_app" "test" {
  name                = "acctest-FA-%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  service_plan_id     = azurerm_service_plan.test.id

  storage_account_name       = azurerm_storage_account.test.name
  storage_account_access_key = azurerm_storage_account.test.primary_access_key

  site_config {
    application_stack {
      docker {
        registry_url = "https://mcr.microsoft.com"
        image_name   = "azure-app-service/samples/aspnethelloworld"
        image_tag    = "latest"
      }
    }
  }
}
`, r.template(data, planSku), data.RandomInteger)
}

func (r LinuxFunctionAppResource) appStackDockerUseMSI(data acceptance.TestData, planSku string) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_linux_function_app" "test" {
  name                = "acctest-FA-%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  service_plan_id     = azurerm_service_plan.test.id

  storage_account_name       = azurerm_storage_account.test.name
  storage_account_access_key = azurerm_storage_account.test.primary_access_key

  identity {
    type         = "UserAssigned"
    identity_ids = [azurerm_user_assigned_identity.test.id]
  }
  site_config {
    container_registry_use_managed_identity       = true
    container_registry_managed_identity_client_id = azurerm_user_assigned_identity.test.client_id

    application_stack {
      docker {
        registry_url = "https://mcr.microsoft.com"
        image_name   = "azure-app-service/samples/aspnethelloworld"
        image_tag    = "latest"
      }
    }
  }
}
`, r.identityTemplate(data, planSku), data.RandomInteger)
}

func (r LinuxFunctionAppResource) backup(data acceptance.TestData, planSku string) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_linux_function_app" "test" {
  name                = "acctest-FA-%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  service_plan_id     = azurerm_service_plan.test.id

  storage_account_name       = azurerm_storage_account.test.name
  storage_account_access_key = azurerm_storage_account.test.primary_access_key

  backup {
    name                = "acctest"
    storage_account_url = "https://${azurerm_storage_account.test.name}.blob.core.windows.net/${azurerm_storage_container.test.name}${data.azurerm_storage_account_sas.test.sas}&sr=b"
    schedule {
      frequency_interval = 7
      frequency_unit     = "Day"
    }
  }

  site_config {}
}
`, r.storageContainerTemplate(data, planSku), data.RandomInteger)
}

func (r LinuxFunctionAppResource) consumptionComplete(data acceptance.TestData) string {
	planSku := "Y1"
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_user_assigned_identity" "test" {
  name                = "acct-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
}

resource "azurerm_application_insights" "test" {
  name                = "acctestappinsights-%[2]d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  application_type    = "web"
}

resource "azurerm_linux_function_app" "test" {
  name                = "acctest-FA-%[2]d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  service_plan_id     = azurerm_service_plan.test.id

  storage_account_name       = azurerm_storage_account.test.name
  storage_account_access_key = azurerm_storage_account.test.primary_access_key

  app_settings = {
    foo    = "bar"
    secret = "sauce"
  }

  auth_settings {
    enabled = true
    issuer  = "https://sts.windows.net/%[3]s"

    additional_login_parameters = {
      test_key = "test_value"
    }

    active_directory {
      client_id     = "aadclientid"
      client_secret = "aadsecret"

      allowed_audiences = [
        "activedirectorytokenaudiences",
      ]
    }

    facebook {
      app_id     = "facebookappid"
      app_secret = "facebookappsecret"

      oauth_scopes = [
        "facebookscope",
      ]
    }
  }

  builtin_logging_enabled = false
  client_cert_enabled     = true
  client_cert_mode        = "Required"

  connection_string {
    name  = "Second"
    value = "some-postgresql-connection-string"
    type  = "PostgreSQL"
  }

  enabled                     = false
  functions_extension_version = "~3"
  https_only                  = true

  identity {
    type         = "UserAssigned"
    identity_ids = [azurerm_user_assigned_identity.test.id]
  }

  site_config {
    app_command_line   = "whoami"
    api_definition_url = "https://example.com/azure_function_app_def.json"
    app_scale_limit    = 3
    // api_management_api_id = ""  // TODO
    application_insights_key               = azurerm_application_insights.test.instrumentation_key
    application_insights_connection_string = azurerm_application_insights.test.connection_string

    container_registry_use_managed_identity       = true
    container_registry_managed_identity_client_id = azurerm_user_assigned_identity.test.client_id

    default_documents = [
      "first.html",
      "second.jsp",
      "third.aspx",
      "hostingstart.html",
    ]

    http2_enabled = true
    ip_restriction {
      ip_address = "10.10.10.10/32"
      name       = "test-restriction"
      priority   = 123
      action     = "Allow"
      headers {
        x_azure_fdid      = ["55ce4ed1-4b06-4bf1-b40e-4638452104da"]
        x_fd_health_probe = ["1"]
        x_forwarded_for   = ["9.9.9.9/32", "2002::1234:abcd:ffff:c0a8:101/64"]
        x_forwarded_host  = ["example.com"]
      }
    }
    load_balancing_mode      = "LeastResponseTime"
    remote_debugging         = true
    remote_debugging_version = "VS2019"

    scm_ip_restriction {
      ip_address = "10.20.20.20/32"
      name       = "test-scm-restriction"
      priority   = 123
      action     = "Allow"
      headers {
        x_azure_fdid      = ["55ce4ed1-4b06-4bf1-b40e-4638452104da"]
        x_fd_health_probe = ["1"]
        x_forwarded_for   = ["9.9.9.9/32", "2002::1234:abcd:ffff:c0a8:101/64"]
        x_forwarded_host  = ["example.com"]
      }
    }

    use_32_bit_worker  = true
    websockets_enabled = true
    ftps_state         = "FtpsOnly"
    health_check_path  = "/health-check"

    application_stack {
      python_version = "3.9"
    }

    minimum_tls_version     = "1.1"
    scm_minimum_tls_version = "1.1"

    cors {
      allowed_origins = [
        "https://www.contoso.com",
        "www.contoso.com",
      ]

      support_credentials = true
    }
  }

  tags = {
    terraform = "true"
    Env       = "AccTest"
  }
}
`, r.template(data, planSku), data.RandomInteger, data.Client().TenantID)
}

func (r LinuxFunctionAppResource) standardComplete(data acceptance.TestData) string {
	planSku := "S1"
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_user_assigned_identity" "test" {
  name                = "acct-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
}

resource "azurerm_application_insights" "test" {
  name                = "acctestappinsights-%[2]d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  application_type    = "web"
}

resource "azurerm_linux_function_app" "test" {
  name                = "acctest-FA-%[2]d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  service_plan_id     = azurerm_service_plan.test.id

  storage_account_name       = azurerm_storage_account.test.name
  storage_account_access_key = azurerm_storage_account.test.primary_access_key

  app_settings = {
    foo    = "bar"
    secret = "sauce"
  }

  auth_settings {
    enabled = true
    issuer  = "https://sts.windows.net/%[3]s"

    additional_login_parameters = {
      test_key = "test_value"
    }

    active_directory {
      client_id     = "aadclientid"
      client_secret = "aadsecret"

      allowed_audiences = [
        "activedirectorytokenaudiences",
      ]
    }

    facebook {
      app_id     = "facebookappid"
      app_secret = "facebookappsecret"

      oauth_scopes = [
        "facebookscope",
      ]
    }
  }

  backup {
    name                = "acctest"
    storage_account_url = "https://${azurerm_storage_account.test.name}.blob.core.windows.net/${azurerm_storage_container.test.name}${data.azurerm_storage_account_sas.test.sas}&sr=b"
    schedule {
      frequency_interval = 7
      frequency_unit     = "Day"
    }
  }

  builtin_logging_enabled = false
  client_cert_enabled     = true
  client_cert_mode        = "OptionalInteractiveUser"

  connection_string {
    name  = "First"
    value = "some-postgresql-connection-string"
    type  = "PostgreSQL"
  }

  enabled                     = false
  functions_extension_version = "~3"
  https_only                  = true

  identity {
    type         = "UserAssigned"
    identity_ids = [azurerm_user_assigned_identity.test.id]
  }

  site_config {
    always_on          = true
    app_command_line   = "whoami"
    api_definition_url = "https://example.com/azure_function_app_def.json"
    // api_management_api_id = ""  // TODO
    application_insights_key               = azurerm_application_insights.test.instrumentation_key
    application_insights_connection_string = azurerm_application_insights.test.connection_string

    application_stack {
      python_version = "3.8"
    }

    container_registry_use_managed_identity       = true
    container_registry_managed_identity_client_id = azurerm_user_assigned_identity.test.client_id

    default_documents = [
      "first.html",
      "second.jsp",
      "third.aspx",
      "hostingstart.html",
    ]

    http2_enabled = true

    ip_restriction {
      ip_address = "10.10.10.10/32"
      name       = "test-restriction"
      priority   = 123
      action     = "Allow"
      headers {
        x_azure_fdid      = ["55ce4ed1-4b06-4bf1-b40e-4638452104da"]
        x_fd_health_probe = ["1"]
        x_forwarded_for   = ["9.9.9.9/32", "2002::1234:abcd:ffff:c0a8:101/64"]
        x_forwarded_host  = ["example.com"]
      }
    }

    load_balancing_mode       = "LeastResponseTime"
    pre_warmed_instance_count = 2
    remote_debugging          = true
    remote_debugging_version  = "VS2017"

    scm_ip_restriction {
      ip_address = "10.20.20.20/32"
      name       = "test-scm-restriction"
      priority   = 123
      action     = "Allow"
      headers {
        x_azure_fdid      = ["55ce4ed1-4b06-4bf1-b40e-4638452104da"]
        x_fd_health_probe = ["1"]
        x_forwarded_for   = ["9.9.9.9/32", "2002::1234:abcd:ffff:c0a8:101/64"]
        x_forwarded_host  = ["example.com"]
      }
    }

    use_32_bit_worker  = true
    websockets_enabled = true
    ftps_state         = "FtpsOnly"
    health_check_path  = "/health-check"
    number_of_workers  = 3

    minimum_tls_version     = "1.1"
    scm_minimum_tls_version = "1.1"

    cors {
      allowed_origins = [
        "https://www.contoso.com",
        "www.contoso.com",
      ]

      support_credentials = true
    }

    vnet_route_all_enabled = true
  }

  tags = {
    terraform = "true"
    Env       = "AccTest"
  }
}
`, r.storageContainerTemplate(data, planSku), data.RandomInteger, data.Client().TenantID)
}

func (r LinuxFunctionAppResource) elasticComplete(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_linux_function_app" "test" {
  name                = "acctest-FA-%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  service_plan_id     = azurerm_service_plan.test.id

  storage_account_name       = azurerm_storage_account.test.name
  storage_account_access_key = azurerm_storage_account.test.primary_access_key

  app_settings = {
    foo    = "bar"
    secret = "sauce"
  }

  backup {
    name                = "acctest"
    storage_account_url = "https://${azurerm_storage_account.test.name}.blob.core.windows.net/${azurerm_storage_container.test.name}${data.azurerm_storage_account_sas.test.sas}&sr=b"
    schedule {
      frequency_interval = 7
      frequency_unit     = "Day"
    }
  }

  connection_string {
    name  = "Example"
    value = "some-postgresql-connection-string"
    type  = "PostgreSQL"
  }

  site_config {}
}
`, r.storageContainerTemplate(data, SkuElasticPremiumPlan), data.RandomInteger)
}

func (r LinuxFunctionAppResource) servicePlanUpdate(data acceptance.TestData, planSku string) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_linux_function_app" "test" {
  name                = "acctest-FA-%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  service_plan_id     = azurerm_service_plan.update.id

  storage_account_name       = azurerm_storage_account.test.name
  storage_account_access_key = azurerm_storage_account.test.primary_access_key

  site_config {}

  depends_on = [azurerm_service_plan.update]
}
`, r.templateServicePlanUpdate(data, planSku), data.RandomInteger)
}

func (r LinuxFunctionAppResource) updateStorageAccount(data acceptance.TestData, planSku string) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_linux_function_app" "test" {
  name                = "acctest-FA-%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  service_plan_id     = azurerm_service_plan.test.id

  storage_account_name       = azurerm_storage_account.update.name
  storage_account_access_key = azurerm_storage_account.update.primary_access_key

  site_config {}
}
`, r.templateExtraStorageAccount(data, planSku), data.RandomInteger)
}

func (LinuxFunctionAppResource) template(data acceptance.TestData, planSku string) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_storage_account" "test" {
  name                     = "acctestsa%s"
  resource_group_name      = azurerm_resource_group.test.name
  location                 = azurerm_resource_group.test.location
  account_tier             = "Standard"
  account_replication_type = "LRS"
}

resource "azurerm_service_plan" "test" {
  name                = "acctestASP-%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  os_type             = "Linux"
  sku_name            = "%s"
}
`, data.RandomInteger, data.Locations.Primary, data.RandomString, data.RandomInteger, planSku)
}

func (LinuxFunctionAppResource) templateExtraStorageAccount(data acceptance.TestData, planSku string) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%[1]d"
  location = "%[2]s"
}

resource "azurerm_storage_account" "test" {
  name                     = "acctestsa%[3]s"
  resource_group_name      = azurerm_resource_group.test.name
  location                 = azurerm_resource_group.test.location
  account_tier             = "Standard"
  account_replication_type = "LRS"
}

resource "azurerm_storage_account" "update" {
  name                     = "acctestsa2%[3]s"
  resource_group_name      = azurerm_resource_group.test.name
  location                 = azurerm_resource_group.test.location
  account_tier             = "Standard"
  account_replication_type = "LRS"
}

resource "azurerm_service_plan" "test" {
  name                = "acctestASP-%[1]d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  os_type             = "Linux"
  sku_name            = "%[4]s"
}
`, data.RandomInteger, data.Locations.Primary, data.RandomString, planSku)
}

func (r LinuxFunctionAppResource) templateServicePlanUpdate(data acceptance.TestData, planSku string) string {
	return fmt.Sprintf(`
%s

resource "azurerm_service_plan" "update" {
  name                = "acctestASP2-%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  os_type             = "Linux"
  sku_name            = "%s"
}
`, r.template(data, planSku), data.RandomInteger, planSku)
}

func (r LinuxFunctionAppResource) storageContainerTemplate(data acceptance.TestData, planSku string) string {
	return fmt.Sprintf(`
%s

resource "azurerm_storage_container" "test" {
  name                  = "test"
  storage_account_name  = azurerm_storage_account.test.name
  container_access_type = "private"
}

data "azurerm_storage_account_sas" "test" {
  connection_string = azurerm_storage_account.test.primary_connection_string
  https_only        = true

  resource_types {
    service   = false
    container = false
    object    = true
  }

  services {
    blob  = true
    queue = false
    table = false
    file  = false
  }

  start  = "2021-04-01"
  expiry = "2024-03-30"

  permissions {
    read    = false
    write   = true
    delete  = false
    list    = false
    add     = false
    create  = false
    update  = false
    process = false
  }
}
`, r.template(data, planSku))
}

func (r LinuxFunctionAppResource) identityTemplate(data acceptance.TestData, planSku string) string {
	return fmt.Sprintf(`
%s

resource "azurerm_user_assigned_identity" "test" {
  name                = "acct-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
}
`, r.template(data, planSku), data.RandomInteger)
}
