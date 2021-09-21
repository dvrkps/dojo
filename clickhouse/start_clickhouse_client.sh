docker container run -it --rm --net host \
	clickhouse/clickhouse-client:latest \
	--host 127.0.0.1 \
	--password dojopassword
