go build 
dapr run --app-id go-web-port \
         --app-protocol http \
         --dapr-http-port 3500 \
         --log-level debug \
        ./hostfunc
