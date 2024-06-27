APP_NAME = api
GOOS = linux
GOARCH = amd64
VERSION=$(shell git describe --tags --always)

#以下是生成需要的表名-空格间隔
DATABASE = log
TABLES = sys_menu

.PHONY: sql2file
# init sql2file
sql2file:
	@for item in $(TABLES); do \
  		rm -f ./scripts/sql/$$item.sql ; \
		mysqldump --compact --skip-add-drop-table -d -h 54.255.223.36 -P 53306 -uback -pBack_789456123 $$DATABASE $$item > ./scripts/sql/$$item.sql; \
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

.PHONY: build
# go build
build:
	rm -rf bin
	mkdir -p ./bin
	go build -buildvcs=false -o ./bin/ ./


.PHONY: package
# go package
package:
	cp -R configs ./bin/
	cp start.sh ./bin/
	cp stop.sh ./bin/
	sed -i "s/\r//" ./bin/start.sh
	sed -i "s/\r//" ./bin/stop.sh


# show HELP
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
