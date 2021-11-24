package consumption

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/consumption/parse"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/consumption/validate"
	resourceParse "github.com/hashicorp/terraform-provider-azurerm/internal/services/resource/parse"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/timeouts"
	"github.com/hashicorp/terraform-provider-azurerm/utils"
)

func resourceArmConsumptionBudgetResourceGroupDataSource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Read: resourceArmConsumptionBudgetResourceGroupDataSourceRead,
		Timeouts: &pluginsdk.ResourceTimeout{
			Read: pluginsdk.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*pluginsdk.Schema{
			"name": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ValidateFunc: validate.ConsumptionBudgetName(),
			},

			"resource_group_id": {
				Type:     pluginsdk.TypeString,
				Required: true,
			},

			"amount": {
				Type:     pluginsdk.TypeFloat,
				Computed: true,
			},

			"filter": {
				Type:     pluginsdk.TypeList,
				Computed: true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"dimension": {
							Type:     pluginsdk.TypeList,
							Computed: true,
							Elem: &pluginsdk.Resource{
								Schema: map[string]*pluginsdk.Schema{
									"name": {
										Type:     pluginsdk.TypeString,
										Computed: true,
									},
									"operator": {
										Type:     pluginsdk.TypeString,
										Computed: true,
									},
									"values": {
										Type:     pluginsdk.TypeList,
										Computed: true,
										Elem: &pluginsdk.Schema{
											Type: pluginsdk.TypeString,
										},
									},
								},
							},
						},
						"tag": {
							Type:     pluginsdk.TypeList,
							Computed: true,
							Elem: &pluginsdk.Resource{
								Schema: map[string]*pluginsdk.Schema{
									"name": {
										Type:     pluginsdk.TypeString,
										Computed: true,
									},
									"operator": {
										Type:     pluginsdk.TypeString,
										Computed: true,
									},
									"values": {
										Type:     pluginsdk.TypeList,
										Computed: true,
										Elem: &pluginsdk.Schema{
											Type: pluginsdk.TypeString,
										},
									},
								},
							},
						},
						"not": {
							Type:     pluginsdk.TypeList,
							Computed: true,
							Elem: &pluginsdk.Resource{
								Schema: map[string]*pluginsdk.Schema{
									"dimension": {
										Type:     pluginsdk.TypeList,
										Computed: true,
										Elem: &pluginsdk.Resource{
											Schema: map[string]*pluginsdk.Schema{
												"name": {
													Type:     pluginsdk.TypeString,
													Computed: true,
												},
												"operator": {
													Type:     pluginsdk.TypeString,
													Computed: true,
												},
												"values": {
													Type:     pluginsdk.TypeList,
													Computed: true,
													Elem: &pluginsdk.Schema{
														Type: pluginsdk.TypeString,
													},
												},
											},
										},
									},
									"tag": {
										Type:     pluginsdk.TypeList,
										Computed: true,
										Elem: &pluginsdk.Resource{
											Schema: map[string]*pluginsdk.Schema{
												"name": {
													Type:     pluginsdk.TypeString,
													Computed: true,
												},
												"operator": {
													Type:     pluginsdk.TypeString,
													Computed: true,
												},
												"values": {
													Type:     pluginsdk.TypeList,
													Computed: true,
													Elem: &pluginsdk.Schema{
														Type: pluginsdk.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},

			"notification": {
				Type:     pluginsdk.TypeList,
				Computed: true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"enabled": {
							Type:     pluginsdk.TypeBool,
							Computed: true,
						},
						"threshold": {
							Type:     pluginsdk.TypeInt,
							Computed: true,
						},
						"threshold_type": {
							Type:     pluginsdk.TypeString,
							Computed: true,
						},
						"operator": {
							Type:     pluginsdk.TypeString,
							Computed: true,
						},

						"contact_emails": {
							Type:     pluginsdk.TypeList,
							Computed: true,
							Elem: &pluginsdk.Schema{
								Type: pluginsdk.TypeString,
							},
						},

						"contact_groups": {
							Type:     pluginsdk.TypeList,
							Computed: true,
							Elem: &pluginsdk.Schema{
								Type: pluginsdk.TypeString,
							},
						},

						"contact_roles": {
							Type:     pluginsdk.TypeList,
							Computed: true,
							Elem: &pluginsdk.Schema{
								Type: pluginsdk.TypeString,
							},
						},
					},
				},
			},

			"time_grain": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},

			"time_period": {
				Type:     pluginsdk.TypeList,
				Computed: true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"start_date": {
							Type:     pluginsdk.TypeString,
							Computed: true,
						},
						"end_date": {
							Type:     pluginsdk.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceArmConsumptionBudgetResourceGroupDataSourceRead(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Consumption.BudgetsClient
	subscriptionID := meta.(*clients.Client).Account.SubscriptionId
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	resourceGroupId, err := resourceParse.ResourceGroupID(d.Get("resource_group_id").(string))
	if err != nil {
		return err
	}

	id := parse.NewConsumptionBudgetResourceGroupID(subscriptionID, resourceGroupId.ResourceGroup, d.Get("name").(string))
	d.SetId(id.ID())

	resp, err := client.Get(ctx, resourceGroupId.ID(), id.BudgetName)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("making read request on %s: %+v", id, err)
	}

	d.Set("name", resp.Name)
	if resp.Amount != nil {
		amount, _ := resp.Amount.Float64()
		d.Set("amount", amount)
	}

	d.Set("time_grain", string(resp.TimeGrain))
	d.Set("time_period", FlattenConsumptionBudgetTimePeriod(resp.TimePeriod))
	d.Set("notification", FlattenConsumptionBudgetNotifications(resp.Notifications))
	d.Set("filter", FlattenConsumptionBudgetFilter(resp.Filter))

	return nil
}