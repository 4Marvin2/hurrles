FROM tarantool/tarantool:latest

# COPY ../db/tarantool/hurrles.lua /opt/tarantool

COPY ./db/tarantool/hurrles.lua /opt/tarantool

CMD ["tarantool", "/opt/tarantool/hurrles.lua"]