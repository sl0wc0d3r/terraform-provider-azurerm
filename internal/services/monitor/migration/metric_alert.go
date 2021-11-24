package migration

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-provider-azurerm/helpers/azure"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/monitor/parse"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tags"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
)

var _ pluginsdk.StateUpgrade = MetricAlertUpgradeV0ToV1{}

type MetricAlertUpgradeV0ToV1 struct{}

func (MetricAlertUpgradeV0ToV1) Schema() map[string]*pluginsdk.Schema {
	return metricAlertSchemaForV0AndV1()
}

func (MetricAlertUpgradeV0ToV1) UpgradeFunc() pluginsdk.StateUpgraderFunc {
	return func(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
		// old
		// 	/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/metricAlerts/alert1
		// new:
		// 	/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/metricAlerts/alert1
		oldId, err := azure.ParseAzureResourceID(rawState["id"].(string))
		if err != nil {
			return rawState, err
		}

		alertName := ""
		for key, value := range oldId.Path {
			if strings.EqualFold(key, "metricAlerts") {
				alertName = value
				break
			}
		}

		if alertName == "" {
			return rawState, fmt.Errorf("couldn't find the `metricAlerts` segment in the old resource id %q", oldId)
		}

		newId := parse.NewMetricAlertID(oldId.SubscriptionID, oldId.ResourceGroup, alertName)

		log.Printf("[DEBUG] Updating ID from %q to %q", oldId, newId.ID())

		rawState["id"] = newId.ID()

		return rawState, nil
	}
}

func metricAlertSchemaForV0AndV1() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"name": {
			Type:     pluginsdk.TypeString,
			Required: true,
		},

		"resource_group_name": azure.SchemaResourceGroupName(),

		"scopes": {
			Type:     pluginsdk.TypeSet,
			Required: true,
			MinItems: 1,
			Elem: &pluginsdk.Schema{
				Type: pluginsdk.TypeString,
			},
		},

		"target_resource_type": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			Computed: true,
		},

		"target_resource_location": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			Computed: true,
		},

		// static criteria
		"criteria": {
			Type:     pluginsdk.TypeSet,
			Optional: true,
			MinItems: 1,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"metric_namespace": {
						Type:     pluginsdk.TypeString,
						Required: true,
					},
					"metric_name": {
						Type:     pluginsdk.TypeString,
						Required: true,
					},
					"aggregation": {
						Type:     pluginsdk.TypeString,
						Required: true,
					},
					"dimension": {
						Type:     pluginsdk.TypeList,
						Optional: true,
						Elem: &pluginsdk.Resource{
							Schema: map[string]*pluginsdk.Schema{
								"name": {
									Type:     pluginsdk.TypeString,
									Required: true,
								},
								"operator": {
									Type:     pluginsdk.TypeString,
									Required: true,
								},
								"values": {
									Type:     pluginsdk.TypeList,
									Required: true,
									MinItems: 1,
									Elem: &pluginsdk.Schema{
										Type: pluginsdk.TypeString,
									},
								},
							},
						},
					},
					"operator": {
						Type:     pluginsdk.TypeString,
						Required: true,
					},
					"threshold": {
						Type:     pluginsdk.TypeFloat,
						Required: true,
					},
					"skip_metric_validation": {
						Type:     pluginsdk.TypeBool,
						Optional: true,
					},
				},
			},
		},

		// lintignore: S018
		"dynamic_criteria": {
			Type:     pluginsdk.TypeSet,
			Optional: true,
			MinItems: 1,
			// Curently, it allows to define only one dynamic criteria in one metric alert.
			MaxItems: 1,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"metric_namespace": {
						Type:     pluginsdk.TypeString,
						Required: true,
					},
					"metric_name": {
						Type:     pluginsdk.TypeString,
						Required: true,
					},
					"aggregation": {
						Type:     pluginsdk.TypeString,
						Required: true,
					},
					"dimension": {
						Type:     pluginsdk.TypeList,
						Optional: true,
						Elem: &pluginsdk.Resource{
							Schema: map[string]*pluginsdk.Schema{
								"name": {
									Type:     pluginsdk.TypeString,
									Required: true,
								},
								"operator": {
									Type:     pluginsdk.TypeString,
									Required: true,
								},
								"values": {
									Type:     pluginsdk.TypeList,
									Required: true,
									MinItems: 1,
									Elem: &pluginsdk.Schema{
										Type: pluginsdk.TypeString,
									},
								},
							},
						},
					},
					"operator": {
						Type:     pluginsdk.TypeString,
						Required: true,
					},
					"alert_sensitivity": {
						Type:     pluginsdk.TypeString,
						Required: true,
					},

					"evaluation_total_count": {
						Type:     pluginsdk.TypeInt,
						Optional: true,
					},

					"evaluation_failure_count": {
						Type:     pluginsdk.TypeInt,
						Optional: true,
					},

					"ignore_data_before": {
						Type:     pluginsdk.TypeString,
						Optional: true,
					},
					"skip_metric_validation": {
						Type:     pluginsdk.TypeBool,
						Optional: true,
					},
				},
			},
		},

		"application_insights_web_test_location_availability_criteria": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			MinItems: 1,
			MaxItems: 1,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"web_test_id": {
						Type:     pluginsdk.TypeString,
						Required: true,
					},
					"component_id": {
						Type:     pluginsdk.TypeString,
						Required: true,
					},
					"failed_location_count": {
						Type:     pluginsdk.TypeInt,
						Required: true,
					},
				},
			},
		},

		"action": {
			Type:     pluginsdk.TypeSet,
			Optional: true,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"action_group_id": {
						Type:     pluginsdk.TypeString,
						Required: true,
					},
					"webhook_properties": {
						Type:     pluginsdk.TypeMap,
						Optional: true,
						Elem: &pluginsdk.Schema{
							Type: pluginsdk.TypeString,
						},
					},
				},
			},
		},

		"auto_mitigate": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
		},

		"description": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},

		"enabled": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
		},

		"frequency": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},

		"severity": {
			Type:     pluginsdk.TypeInt,
			Optional: true,
		},

		"window_size": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},

		"tags": tags.Schema(),
	}
}
