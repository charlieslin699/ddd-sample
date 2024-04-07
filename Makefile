# mockery版本
MOCKERY_VERSION = v2.42.1

generate-command:
	sh generate_command.sh $(domain) $(command) $(filename)

generate-query:
	sh generate_query.sh $(domain) $(query) $(filename)

mocks:
	@if [ "$$(mockery --version | awk '{print $$1}')" = "${MOCKERY_VERSION}" ]; then \
		mockery; \
	else \
		go install github.com/vektra/mockery/v2@${MOCKERY_VERSION}; \
		mockery; \
	fi

mocks-install:
	go install github.com/vektra/mockery/v2@${MOCKERY_VERSION}

swag-init:
	swag init -g cmd/apiserver/swagger/main.go -o cmd/apiserver/swagger/docs

help:
	@echo "make generate-command      -- 自動產生command"
	@echo "make generate-query        -- 自動產生query"
	@echo "make mocks                 -- 產生mocks檔"
