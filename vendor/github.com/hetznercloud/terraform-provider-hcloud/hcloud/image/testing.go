package image

import (
	"fmt"

	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/testtemplate"
)

// DData defines the fields for the "testdata/d/hcloud_image"
// template.
type DData struct {
	testtemplate.DataCommon

	ImageID       string
	ImageName     string
	LabelSelector string
}

// TFID returns the data source identifier.
func (d *DData) TFID() string {
	return fmt.Sprintf("data.%s.%s", DataSourceType, d.RName())
}
