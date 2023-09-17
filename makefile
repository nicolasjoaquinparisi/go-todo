#!/bin/bash

tests:
	@docker compose -f ./docker-compose.test.yaml down --remove-orphans -v && \
	docker compose -f ./docker-compose.test.yaml up -d