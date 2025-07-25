# These files are input OpenAPI specs from the Microsoft-supplied submodule.
SPECFILE_METERING=./commercial-marketplace-openapi/Microsoft.Marketplace.Metering/2018-08-31/meteringapi.v1
SPECFILE_FULFILLMENT=./commercial-marketplace-openapi/Microsoft.Marketplace.SaaS/2018-08-31/saasapi.v2

# For API users; modify these tests and run them from these targets
test-metering:
	source ~/ful-env.sh && cd metering && go test -v ./*.go -run=TestMeteringClient
	source ~/ful-env.sh && cd metering && go test -v ./*.go -run=TestMeteringPostEvent

test-fulfillment:
	source ~/ful-env.sh && cd fulfillment && go test -v ./*.go -run=TestFulfillment
	source ~/ful-env.sh && cd fulfillment && go test -v ./*.go -run=TestSubOps

# For package maintainers; modify these targets to update the generated code
# Use this to nuke your Autorest installation and install latest tool and language plugins
autorest-go-update-with-reset: 
	# You need to prefix sudo for the install globally command below.
	sudo npm install -g autorest
	autorest --reset
	autorest --go --help

codegen-metering: metering.md
	autorest metering.md --input-file=$(SPECFILE_METERING).json
	cd ./metering && go mod tidy && go get -u ./... && go mod tidy

codegen-fulfillment: fulfillment.md
	autorest fulfillment.md --input-file=$(SPECFILE_FULFILLMENT).json
	cd ./fulfillment && go mod tidy && go get -u ./... && go mod tidy
