
DIR := ${CURDIR}

LIQUIBASE_CONFIGS := --defaultsFile=${DIR}/infra/liquibase/liquibase.properties
RUN_LIQUIBASE := echo ${LIQUIBASE_CONFIGS}

_migrate:
	${RUN_LIQUIBASE} update

tag:
	${RUN_LIQUIBASE} tag "$(shell date '+%Y-%m-%d %H:%M:%S')"

migrate: _migrate tag

rollback:
	${RUN_LIQUIBASE} rollback

test:
	echo

dangerously-exclude-all-migrations:
	${RUN_LIQUIBASE} rollback rollbackToDate 2000-01-01 18:56:57.668819