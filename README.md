# go-seeding-omdb/ayo-bibit-omdb

## What's going on here?
- [Setup](#01_Setup)
- [Request](#02_Request-Response)

## 01_Setup
- Create a config file named `config.yml` in root folder with same format with `config.sample.yml`

```yml
SERVICE:
  RESTPORT: 5000
  GRPCPORT: 6000
  APIKEY: yourapi
DATABASE: 
  HOST: localhost
  PORT: 3306
  USER: root
  PASSWORD: root
  NAME: db_name
```

- Create the Database with name align as in the configuration file

```bash
$ mysql -u root -p
msql> CREATE DATABASE omdb;
```

- Create `.pb.go` for gRPC models

```bash
$ cd go-seeding-omdb/models
$ protoc --go-grpc_out=. *.proto
```

## 02_Request-Response

- Available Endpoints

```txt
GET: {BASE_URL}/api/v1/movie?pagination={page}&searchword={query}
GET: {BASE_URL}/api/v1/movie/:imdbID
```

- Success Response

> List: {BASE_URL}/api/v1/movie?pagination={page}&searchword={query}
```json
{
    "data": [
        {
            "title": "Batman: The Killing Joke",
            "year": "2016",
            "imdbID": "tt4853102",
            "type": "movie",
            "poster": "https://example.com/images/image.jpg"
        },
        ...
    ],
    "message": "success fetch movies list",
    "status": 200,
    "meta": {
        "count": 10,
        "total": 463
    }
}
```
> Detail: {BASE_URL}/api/v1/movie/:imdbID
```json
{
    "data": {
        "title": "Naruto: The Lost Story - Mission: Protect the Waterfall Village",
        "year": "2003",
        "imdbID": "tt3634858",
        "type": "movie",
        "poster": "https://example.com/images/image.jpg",
        "released": "12 Jun 2007",
        "runtime": "40 min",
        "director": "Masahiko Murata, Hayato Date",
        "writer": "Masashi Kishimoto (original manga), Katsuyuki Sumizawa (screenplay)",
        "actors": "Junko Takeuchi, Noriaki Sugiyama, Chie Nakamura, Kazuhiko Inoue",
        "plot": "Naruto and his friends must get back a jug of stolen holy water from a band of higher class ninjas.",
        "language": "Japanese",
        "country": "Japan",
        "imdbRating": "6.6"
    },
    "message": "success fetch movie detail",
    "status": 200,
    "meta": {}
}
```

- Error Response

```json
{
    "data": {
        "reason": "searchword doesn't match with any of available data"
    },
    "message": "failed to fetch movies list",
    "status": 400,
    "errorCode": "MOVIE-001",
    "meta": {}
}
```
```json
{
    "data": {
        "reason": "imdb id not found"
    },
    "message": "failed to fetch movie detail",
    "status": 400,
    "errorCode": "MOVIE-002",
    "meta": {}
}
```
