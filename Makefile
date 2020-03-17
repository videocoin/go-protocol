.PHONY: abigen
abigen:
	./_assets/abigen.sh contracts abi

.PHONY: libs
libs:
	./_assets/libs.sh

.PHONY: codegen
codegen:
	./_assets/codegen.sh abi

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
