export BOT_TOKEN=
export BOT_KEY=
export SECRET=
export DATA_SOURCE_NAME=

make dropdb && \
make createdb && \
make migrateset

clear && make build && ./bin/mvthbot