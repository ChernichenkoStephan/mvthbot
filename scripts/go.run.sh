export BOT_TOKEN=<TOKEN>
export SECRET=<TOKEN>
export DB_USER=admin
export DB_NAME=mvthdb

make dropdb && \
make createdb && \
make migrateset

clear && make build && ./bin/mvthbot