FROM mariadb:10.3

COPY ./migrations/* /docker-entrypoint-initdb.d/

CMD ["mysqld", "--character-set-server=utf8", "--collation-server=utf8_unicode_ci"]