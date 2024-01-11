package main

import (
	"os"
	"path/filepath"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi-terraform/sdk/v5/go/state"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		tfStatePath := filepath.Join(cwd, "../tf/terraform.tfstate")

		state, err := state.NewRemoteStateReference(ctx, "localstate", &state.LocalStateArgs{
            Path: pulumi.String(tfStatePath),
        })
        if err != nil {
            return err
        }

		naclId := state.Outputs.ApplyT(func(args interface{}) (string, error) {
			id := args.(map[string]interface{})["nacl_id"].(string)
            return id, nil
        }).(pulumi.StringOutput)

		sgId := state.Outputs.ApplyT(func(args interface{}) (string, error) {
			id := args.(map[string]interface{})["sg_id"].(string)
            return id, nil
        }).(pulumi.StringOutput)

		naclRuleIngressHttp, err := ec2.NewNetworkAclRule(ctx, "nacl_ingress_http", &ec2.NetworkAclRuleArgs{
			NetworkAclId: 	naclId,
			RuleNumber: 	pulumi.Int(100),
			Egress: 		pulumi.Bool(false),
			Protocol: 		pulumi.String("tcp"),
			RuleAction: 	pulumi.String("allow"),
			CidrBlock: 		pulumi.String("0.0.0.0/0"),
			FromPort: 		pulumi.Int(80),
			ToPort: 		pulumi.Int(80),
		})
        if err != nil {
            return err
        }

		sgRuleIngressHttp, err := ec2.NewSecurityGroupRule(ctx, "sg_ingress_http", &ec2.SecurityGroupRuleArgs{
			Type:     pulumi.String("ingress"),
			FromPort: pulumi.Int(80),
			ToPort:   pulumi.Int(80),
			Protocol: pulumi.String("tcp"),
			CidrBlocks: pulumi.StringArray{
				pulumi.String("0.0.0.0/0"),
			},
			SecurityGroupId: sgId,
		})
		if err != nil {
			return err
		}

		ctx.Export("naclRuleIngressHttp ID", naclRuleIngressHttp.ID())
		ctx.Export("sgRuleIngressHttp ID", sgRuleIngressHttp.ID())

		return nil
	})
}
