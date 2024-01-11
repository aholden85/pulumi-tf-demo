.DEFAULT_GOAL := build

# --- Terraform vars
terraform_folder := tf

# --- Pulumi vars
pulumi_folder := pulumi
pulumi_stack_name := dev

all: build destroy destroy-full build-terraform destroy-terraform build-pulumi destroy-pulumi

.PHONY: all

# - INIT

# - BUILD
build: build-terraform build-pulumi

# --- Terraform
build-terraform:
	cd ${terraform_folder}; \
	terraform apply -auto-approve; \
	cd ..

# --- Pulumi
build-pulumi:
	cd ${pulumi_folder}; \
	pulumi up --stack ${pulumi_stack_name} --yes; \
	cd ..

# - TEST
test:
	

# - DESTROY
destroy: destroy-pulumi destroy-terraform
destroy-full: destroy
	cd ${pulumi_folder}; \
	pulumi stack rm ${pulumi_stack_name}; \
	cd ..

# --- Terraform
destroy-terraform:
	cd ${terraform_folder}; \
	terraform apply -destroy -auto-approve; \
	cd ..

# --- Pulumi
destroy-pulumi:
	cd ${pulumi_folder}; \
	pulumi destroy --stack ${pulumi_stack_name} --yes; \
	cd ..
