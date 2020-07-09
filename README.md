# Cache
An implementation of cache, postgres and rabbitmq

The project is meant to be deployed on docker container using docker-compose.

1. Clone the repository 
2. Go inside the Cache folder (run 'cd Cache')
3. Run 'docker-compose up' command.
4. Three docker containers will be deployed all together.
5. The endpoints can be accessed on localhost:9091 (unless docker toolbox is being used)
6. Hit the cache/create endpoint to insert data in cache.
7. Hit the cache/fetch endpoint to fecth data. Pagination is supported, 'limit' and 'offset' numerical values can be sent as query parameters.
8. Performace metrics is handled using pprof. Profiles can be viewed by running 'go tool pprof -top http://localhost:9091/debug/pprof/heap', 'go tool pprof -top http://localhost:9091/debug/pprof/goroutine' or 'go tool pprof -top http://localhost:9091/debug/pprof/allocs' command 
