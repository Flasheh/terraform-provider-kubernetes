// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetes

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func dataSourceKubernetesSecretV1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceKubernetesSecretV1Read,

		Schema: map[string]*schema.Schema{
			"metadata": namespacedMetadataSchema("secret", true),
			"data": {
				Type:        schema.TypeMap,
				Description: "A map of the secret data.",
				Computed:    true,
				Sensitive:   true,
			},
			"binary_data": {
				Type:        schema.TypeMap,
				Description: "A map of the secret data with values encoded in base64 format",
				Optional:    true,
				Sensitive:   true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Type of secret",
				Computed:    true,
			},
			"immutable": {
				Type:        schema.TypeBool,
				Description: "Ensures that data stored in the Secret cannot be updated (only object metadata can be modified).",
				Computed:    true,
			},
		},
	}
}

func dataSourceKubernetesSecretV1Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	om := meta_v1.ObjectMeta{
		Namespace: d.Get("metadata.0.namespace").(string),
		Name:      d.Get("metadata.0.name").(string),
	}
	d.SetId(buildId(om))

	return resourceKubernetesSecretV1Read(ctx, d, meta)
}
