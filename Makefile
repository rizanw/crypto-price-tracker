.PHONY: build
build:
	go build -v -o bin/crypto-tracker cmd/*.go

.PHONY: build
run:
	@echo " >> build crypto-tracker"
	@make build
	@echo " >> crypto-tracker built."
	@echo " >> executing crypto-tracker"
	@./bin/crypto-tracker
	@echo " >> crypto-tracker is running"

.PHONY: testf
testf:
	@make test | grep "FAIL" || echo "ALL tests passed"