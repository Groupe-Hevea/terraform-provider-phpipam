init:
	terraform init

plan:
	terraform plan

clean:
	rm -rf .terraform/
	rm -rf *.state*
	rm -rf .terraform.lock.hcl

all: clean init plan
