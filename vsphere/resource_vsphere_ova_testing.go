package vsphere

import (
    "fmt"
    "log"
    "strings"
    "archive/tar"
    "errors"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/vmware/govmomi"
    "github.com/vmware/govmomi/find"
    "github.com/vmware/govmomi/object"
    "github.com/vmware/govmomi/vim25/soap"
    "github.com/vmware/govmomi/vim25/types"
    "golang.org/x/net/context"
)

type ova_file struct {
    templateName  string
    sourceFile    string
    datacenter    string
    resourcePool  string
    datastore     string
}

func resourceVSphereOvaTemplate() *schema.Resource {
    return &schema.resource{
        Create: resourceVsphereOvaTemplateCreate,
        Read: resourceVsphereOvaTemplateRead,
        Delete: resourceVsphereOvaTemplateDelete,

        Schema: map[string]*schema.Schema{
            "template_name": {
                Type:     schema.TypeString,
                Required: true,
            },

            "source_file": {
                Type:     schema.TypeString,
                Required: true,
            },

            "datacenter": {
                Type:     schema.TypeString,
                Optional: true,
            },

            "resource_pool": {
                Type:      schema.TypeString,
                Optional:  true,
            },

            "datastore": {
                Type:   schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceVsphereOvaTemplateCreate(d * schema.ResourceData, meta interface{}) error {
    log.Printf("[DEBUG] creating template %#v", d)
    client := meta.(*govmomi.Client)

    ova := ova_file{}

    if v, ok := d.GetOk("source_file"); ok {
        ova.sourceFile = v.(string)
    } else {
        return fmt.Errorf("source_file argument is required")
    }
    if v, ok := d.GetOk("template_name"); ok && v != "" {
        ova.templateName = v (string)
    }  else {
       return fmt.Errorf("template_name argument is required")
    }

    if v, ok := d.GetOk("datacenter"); ok {
        ova.datacenter = v.(string)
    }

    if v, ok := d.GetOk("datastore"); ok {
        ova.datastore = v.(string)
    }  else {
        return fmt.Errorf("datastore argument is required")
    }

    if v, ok := d.GetOk("resource_pool"); ok {
        ova.resourcePool = v.(string)
    }
    err:= processOva(client, &ova)
    if err != nil {
        return err
    }

    log.Printf("[INFO] Uploaded template: %s", ova.templateName)
    return
}

func processOva(client *govomi.Client, ova *file) error {
    buf := new(bytes.Buffer)
    


}
"""
Functions needed:
    GetDatacenter:  resourceVSphereGetDatacenter,
    GetDataStore: resourceVsphereGetDataStore,
    GetResourcePool: resourceVsphereGetResourcePool,
    GetLargestFreeResourcePool: resourceVsphereGetLargestFreeResourcePool,
    GetLargestFreeDataStore: resourceVsphereGetFreeDataStore,
    GetTarFileSize: resourceVsphereGetTarFileSize
"""
