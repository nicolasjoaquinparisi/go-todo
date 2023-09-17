#!/bin/bash

tests:
	@docker compose -f ./_test/docker-compose.test.yaml down --remove-orphans -v && \
	docker compose -f ./_test/docker-compose.test.yaml up -d