go build &&
dapr run --app-id dapr-demo \
         --app-protocol http \
         --app-port 9003 \
         --dapr-http-port 3501 \
         --log-level debug \
         ./dapr-demo
