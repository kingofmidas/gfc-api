# GFC API (golang, docker, postgresql)
 Build the images and run the containers:

```sh
$ docker-compose up -d
```
Test it out at [http://localhost:8000](http://localhost:8080).

 Example request to create order:
````json
{"item_list":[{"item_id": 1, "count": 2}, {"item_id": 2, "count": 5}], "status":"not ready"}
````