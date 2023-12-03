FROM postgres:15.5-bullseye

COPY initdb/init.sql /docker-entrypoint-initdb.d/

ENV POSTGRES_USER=todo
ENV POSTGRES_PASSWORD=todo123
ENV POSTGRES_DB=todo

CMD ["docker-entrypoint.sh", "postgres"]