package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTechsupportmanagementCollectionControlPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTechsupportmanagementCollectionControlPolicyRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"deployment_type": {
				Description: "Deployment type defines whether the policy is associated with a SaaS or Appliance account.\n* `None` - Service deployment type None.\n* `SaaS` - Service deployment type SaaS.\n* `Appliance` - Service deployment type Appliance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"tech_support_collection": {
				Description: "Enable or Disable techsupport collection for a specific account.\n* `Enable` - Enable techsupport collection.\n* `Disable` - Disable techsupport collection.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceTechsupportmanagementCollectionControlPolicy().Schema},
				Computed: true,
			}},
	}
}

func dataSourceTechsupportmanagementCollectionControlPolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.TechsupportmanagementCollectionControlPolicy{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("deployment_type"); ok {
		x := (v.(string))
		o.SetDeploymentType(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("tech_support_collection"); ok {
		x := (v.(string))
		o.SetTechSupportCollection(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of TechsupportmanagementCollectionControlPolicy object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.TechsupportmanagementApi.GetTechsupportmanagementCollectionControlPolicyList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of TechsupportmanagementCollectionControlPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.TechsupportmanagementCollectionControlPolicyList.GetCount()
	var i int32
	var techsupportmanagementCollectionControlPolicyResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.TechsupportmanagementApi.GetTechsupportmanagementCollectionControlPolicyList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching TechsupportmanagementCollectionControlPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.TechsupportmanagementCollectionControlPolicyList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for TechsupportmanagementCollectionControlPolicy data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})

				temp["account"] = flattenMapIamAccountRelationship(s.GetAccount(), d)
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["deployment_type"] = (s.GetDeploymentType())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["tech_support_collection"] = (s.GetTechSupportCollection())
				techsupportmanagementCollectionControlPolicyResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(techsupportmanagementCollectionControlPolicyResults))
	if err := d.Set("results", techsupportmanagementCollectionControlPolicyResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(techsupportmanagementCollectionControlPolicyResults[0]["moid"].(string))
	return de
}