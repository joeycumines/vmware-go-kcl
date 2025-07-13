LOCALSTACK_ENDPOINT := http://localhost:4566
STREAM_NAME := kcl-test
TABLE_NAME := appName
AWS_REGION := us-west-2

##@ Integration Targets

.PHONY: start-localstack
start-localstack: ## Start LocalStack using Docker Compose.
	@echo "Starting LocalStack..."
	docker compose up -d localstack
	@echo "Waiting for LocalStack to be ready..."
	@for i in {1..30}; do \
		if curl -s $(LOCALSTACK_ENDPOINT)/_localstack/health | grep -q '"kinesis": "available"'; then \
			echo "LocalStack is ready!"; \
			break; \
		fi; \
		echo "Waiting for LocalStack... ($$i/30)"; \
		sleep 2; \
	done

.PHONY: stop-localstack
stop-localstack: ## Stop LocalStack.
	@echo "Stopping LocalStack..."
	docker compose down

.PHONY: clean-localstack
clean-localstack: ## Clean up LocalStack resources, containers, and state.
	$(info $(shell docker compose kill))
	docker compose rm -f -s -v
	rm -rf .localstack

.PHONY: create-test-resources
create-test-resources: ## Set up test resources in LocalStack.
	@echo "Setting up test resources..."
	@echo "Creating Kinesis stream: $(STREAM_NAME)"
	@AWS_ACCESS_KEY_ID=test AWS_SECRET_ACCESS_KEY=test \
		aws --endpoint-url=$(LOCALSTACK_ENDPOINT) --region=$(AWS_REGION) \
		kinesis create-stream --stream-name $(STREAM_NAME) --shard-count 2 || \
		echo "Stream may already exist"
	@echo "Waiting for stream to become active..."
	@AWS_ACCESS_KEY_ID=test AWS_SECRET_ACCESS_KEY=test \
		aws --endpoint-url=$(LOCALSTACK_ENDPOINT) --region=$(AWS_REGION) \
		kinesis wait stream-exists --stream-name $(STREAM_NAME)
	@echo "Test resources ready!"

.PHONY: delete-test-resources
delete-test-resources: ## Clean up test resources after running tests.
	@echo "Cleaning up test resources..."
	AWS_ACCESS_KEY_ID=test AWS_SECRET_ACCESS_KEY=test aws --endpoint-url=$(LOCALSTACK_ENDPOINT) --region=$(AWS_REGION) kinesis delete-stream --stream-name $(STREAM_NAME) || echo "Stream may not exist"
	AWS_ACCESS_KEY_ID=test AWS_SECRET_ACCESS_KEY=test aws --endpoint-url=$(LOCALSTACK_ENDPOINT) --region=$(AWS_REGION) dynamodb delete-table --table-name $(TABLE_NAME) || echo "Table may not exist"

.PHONY: reset-integration-tests
reset-integration-tests: ## Prepare for integration tests, e.g. start+setup LocalStack.
	@$(MAKE) --no-print-directory clean-localstack
	@$(MAKE) --no-print-directory start-localstack
	@$(MAKE) --no-print-directory create-test-resources

.PHONY: run-integration-tests
run-integration-tests: reset-integration-tests ## Run integration tests against LocalStack.
	@echo "Running integration tests..."
	AWS_ACCESS_KEY_ID=test \
	AWS_SECRET_ACCESS_KEY=test \
	AWS_REGION=$(AWS_REGION) \
	KINESIS_ENDPOINT=$(LOCALSTACK_ENDPOINT) \
	DYNAMODB_ENDPOINT=$(LOCALSTACK_ENDPOINT) \
	go test -v -timeout=3h ./test

