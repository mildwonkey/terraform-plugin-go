package fromproto

import (
	tfproto "github.com/hashicorp/terraform-protocol-go"
	"github.com/hashicorp/terraform-protocol-go/internal/tfplugin5"
	"github.com/hashicorp/terraform-protocol-go/tfprotov5"
)

func CtyBlock(in tfplugin5.DynamicValue) (tfprotov5.CtyBlock, error) {
	// TODO: figure out how to do this
	return tfprotov5.CtyBlock{}, tfproto.ErrUnimplemented
}