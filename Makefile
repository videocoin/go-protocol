# solc has issues with non-root permissions
# e.g. fails to run with --user $(shell id -u):$(shell id -g)
ABIGEN_DOCKER=docker run \
	-v $(shell pwd)/contracts:/contracts \
	-v $(shell pwd)/abi:/abi \
	--rm -ti ethereum/solc:0.5.16
CODEGEN_DOCKER=docker run \
	--user $(shell id -u):$(shell id -g) \
	-v $(shell pwd)/abi:/abi \
	-v $(shell pwd)/staking:/staking \
	-v $(shell pwd)/streams:/streams \
	--rm -ti ethereum/client-go:alltools-latest abigen

.PHONY: abigen
abigen:
	./_assets/abigen.sh $(shell pwd)/contracts $(shell pwd)/abi

.PHONY: libs
libs:
	./_assets/libs.sh

.PHONY: dirs
dirs:
	mkdir -p staking
	mkdir -p streams

.PHONY: codegen
codegen: dirs
	./_assets/codegen.sh $(shell pwd)/abi $(shell pwd)/staking $(shell pwd)/streams

.PHONY: gomod
gomod:
	go mod tidy

.PHONY: generate
generate: libs abigen codegen gomod

.PHONY: clean
clean:
	rm -rf staking
	rm -rf streams
	rm -rf abi

.PHONY: dockerized
dockerized: libs dirs
	COMMAND="${ABIGEN_DOCKER}" ./_assets/abigen.sh /contracts /abi
	COMMAND="${CODEGEN_DOCKER}" ./_assets/codegen.sh /abi /staking /streams
	go mod tidy
