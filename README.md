# pulumi-tf-demo

## Possible Approaches
1. Coexist with resources provisioned by Terraform by referencing a `.tfstate` file.
2. Import existing resources into Pulumi [in the usual way](https://www.pulumi.com/docs/using-pulumi/adopting-pulumi/import/).
3. Convert any Terraform HCL to Pulumi code using `pulumi convert --from terraform`.

## Approach
Using local Terraform statefile, coexist with resources provisioned by Terraform by referencing a `.tfstate` file.

### Stage 1 (Terraform)
Deploy an EC2 instance hosting a web server via HTTP

### Stage 2 (Pulumi)
Deploy a security group rule to allow access to the internet.

### Questions
- Is Terraform aware of the out-of-band changes?
    - Can we push Pulumi changes to the `.tfstate` file?

## Other approaches
1. Coexist
    - Use Terraform Cloud backend
    - Use S3 bucket backend
2. Import
    - Import manual, CFN, TF created resources into Pulumi
3. Convert
    - Convert pre-existing Terraform code into Pulumi code

# Future Demos
- `pulumi-cfn-demo`
- `pulumi-tf-demo` - Phase 2 (Import)