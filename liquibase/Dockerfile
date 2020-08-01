FROM openjdk:11-jre-slim-buster

RUN apt-get update \
    && apt-get -y install gnupg wget \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

RUN mkdir /opt/liquibase
WORKDIR /opt/liquibase

COPY . /opt/liquibase

ARG LIQUIBASE_VERSION=4.0.0
RUN set -x \
    && wget -O liquibase-${LIQUIBASE_VERSION}.tar.gz "https://github.com/liquibase/liquibase/releases/download/v${LIQUIBASE_VERSION}/liquibase-${LIQUIBASE_VERSION}.tar.gz" \
    && tar -xzf liquibase-${LIQUIBASE_VERSION}.tar.gz
RUN ln -s /opt/liquibase/liquibase /usr/local/bin/

RUN ls -la

ENTRYPOINT ["liquibase"]