# solc has issues with non-root permissions
# e.g. fails to run with --user $(shell id -u):$(shell id -g)
ABIGEN_DOCKER=docker run \
	-v $(shell pwd)/contracts:/contracts \
	-v $(shell pwd)/abi:/abi \
	--rm -ti ethereum/solc:0.5.16
BINGEN_DOCKER=docker run \
	-v $(shell pwd)/contracts:/contracts \
	-v $(shell pwd)/build:/build \
	--rm -ti ethereum/solc:0.5.16
CODEGEN_DOCKER=docker run \
	--user $(shell id -u):$(shell id -g) \
	-v $(shell pwd)/abi:/abi \
	-v $(shell pwd)/build:/build \
	-v $(shell pwd)/staking:/staking \
	-v $(shell pwd)/streams:/streams \
	--rm -ti ethereum/client-go:alltools-latest abigen

.PHONY: abigen
abigen:
	./_assets/abigen.sh $(shell pwd)/contracts $(shell pwd)/abi

.PHONY: bingen
bingen:
	mkdir -p build
	./_assets/bingen.sh $(shell pwd)/contracts $(shell pwd)/build

.PHONY: libs
libs:
	./_assets/libs.sh

.PHONY: dirs
dirs:
	mkdir -p staking
	mkdir -p streams
	mkdir -p abi
	mkdir -p build

.PHONY: codegen
codegen: dirs
	./_assets/codegen.sh $(shell pwd)/abi $(shell pwd)/build $(shell pwd)/staking $(shell pwd)/streams

.PHONY: gomod
gomod:
	go mod tidy

.PHONY: generate
generate: abigen bingen codegen gomod

.PHONY: clean
clean:
	rm -rf staking
	rm -rf streams
	rm -rf build
	rm -rf abi

.PHONY: dockerized
dockerized: libs dirs
	COMMAND="${ABIGEN_DOCKER}" ./_assets/abigen.sh /contracts /abi
	COMMAND="${BINGEN_DOCKER}" ./_assets/bingen.sh /contracts /build
	COMMAND="${CODEGEN_DOCKER}" ./_assets/codegen.sh /abi /build /staking /streams
	go mod tidy

.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor
	modvendor -copy="**/*.c **/*.h" -v