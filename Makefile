# Makefile untuk mempersingkat pengetikan command sql migration
include .env
export

.PHONY: build with-swaginit

DSN=mysql://$(DB_USERNAME):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)

migrate-create:
	docker compose run --rm migrate create -ext sql -dir /migrations -seq $(name)
# example command : sudo make migrate-create name=create_users
# $(var) = menerima parameter dari luar

migrate-up:
	docker compose run --rm migrate -path=/migrations -database="$(DSN)" up

migrate-down:
	docker compose run --rm migrate -path=/migrations -database="$(DSN)" down 1

migrate-force:
	docker compose run --rm migrate -path=/migrations -database="$(DSN)" force $(v)

migrate-version:
	docker compose run --rm migrate -path=/migrations -database="$(DSN)" version

# go test path=...
test:
	docker compose run --rm test go test $(path) -v

# swag-init:
# 	swag init -g cmd/web/main.go --output cmd/web/docs

# Swag init included
# sudo make dev-build with-swaginit
dev-build:
ifneq (,$(findstring with-swaginit,$(MAKECMDGOALS)))
# pathnya hardcode, sesuaikan saja pake output which swag
	/home/iann-wsl/go/bin/swag init -g cmd/web/main.go --output cmd/web/docs
endif
	docker compose down
	docker compose build --no-cache
	docker compose up

with-swaginit: ;@: