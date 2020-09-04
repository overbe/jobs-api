# Home exercise from:

https://gist.github.com/renandincer/29802d6c84f92ad79bb94ef54fb92444

# Development

* Run on local environment:

```sh
docker-compose -f deploy/local-docker-compose.yml --compatibility up -d --build
```

* Stop and remove on local environment:

```sh
docker-compose -f deploy/local-docker-compose.yml down
```

* On a file changes run the above command

* Rebuild and restart the specific service

```sh
docker-compose -f deploy/local-docker-compose.yml rm -f -s -v jobs
docker-compose -f deploy/local-docker-compose.yml --compatibility up -d --build jobs
```

* Run jobs server via cmd for debugging:

```sh
docker-compose -f deploy/local-docker-compose.yml --compatibility stop jobs
JOBS_ENV=local go run --race cmd/jobs.go
```

* Server starts on: http://localhost:8080 

* Example test queries at:

```
/deploy/restclient.http
```
