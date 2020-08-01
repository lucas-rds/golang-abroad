
DIR := ${CURDIR}

LIQUIBASE_CONFIGS := --defaultsFile=/infra/liquibase/liquibase.properties
RUN_LIQUIBASE := docker build -t liquibase ./liquibase && docker run --network host liquibase

_migrate:
	${RUN_LIQUIBASE} update

tag:
	${RUN_LIQUIBASE} tag "$(shell date '+%Y-%m-%d %H:%M:%S')"

migrate: _migrate tag

# example: make rollback tag="2000-01-01 18:56:57.668819"
rollback:
	${RUN_LIQUIBASE} rollback $(tag)

dangerouslyRollback:
	${RUN_LIQUIBASE} rollback rollbackToDate 2000-01-01 18:56:57.668819