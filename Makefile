#GOPATH:=$(shell go env GOPATH)
#VERSION=$(shell git describe --tags --always)
TABLES = merchant_channel_bind

.PHONY: sql2file
# init sql2file
sql2file:
	@for item in $(TABLES); do \
  		rm -f ./scripts/sql/$$item.sql ; \
		mysqldump --compact --skip-add-drop-table -d -h 192.168.1.22 -P 3306 -uback -pBack_789456123 merchant $$item > ./scripts/sql/$$item.sql; \
		sed '1d' ./scripts/sql/$$item.sql > ./scripts/sql/$$item.sql.new; \
		sed '1d' ./scripts/sql/$$item.sql.new > ./scripts/sql/$$item.sql; \
		sed '$$d' ./scripts/sql/$$item.sql > ./scripts/sql/$$item.sql.new; \
		cp ./scripts/sql/$$item.sql.new ./scripts/sql/$$item.sql ; \
		rm -f ./scripts/sql/$$item.sql.new ; \
	done

.PHONY: sql2dto
# init sql2dto
sql2dto:
	@for item in $(TABLES); do \
  		rm -f ./internal/logic/dto/$$item.go ; \
		sql2struct -sf=./scripts/sql/$$item.sql -gorm=false -db=false -form=false -with-tablename-func=false -package="dto" > ./internal/logic/dto/$$item.go ; \
	done

.PHONY: sql2po
# init sql2po
sql2po:
	@for item in $(TABLES); do \
		rm -f ./internal/data/po/$$item.go ; \
        sql2struct -sf=./scripts/sql/$$item.sql -db=false -form=false -with-tablename-func=true -package="po" >> ./internal/data/po/$$item.go ; \
	done


# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
