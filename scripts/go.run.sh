export BOT_TOKEN=<TOKEN>
export SECRET=<TOKEN>

make dropdb && \
make createdb && \
make migrateset

clear && make build && ./bin/mvthbot