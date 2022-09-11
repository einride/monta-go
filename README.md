Monta Go SDK
============

A Go SDK for the [Monta](https://monta.app) Partner API.

Usage
-----

### Client

```go
client := monta.NewClient(monta.WithClientIDAndSecret("<ID>", "<SECRET>")
```

#### `GET /v1/auth/me`

```go
me, err := client.GetMe(ctx)
if err != nil {
	panic(err)
}
fmt.Println(me)
```

#### `GET /v1/sites`

```go
response, err := client.ListSites(ctx, &monta.ListSitesRequest{
	Page:    1,
	PerPage: 10,
})
if err != nil {
	panic(err)
}
fmt.Println(response)
```

#### `GET /v1/charge-points`

```go
response, err := client.ListChargePoints(ctx, &monta.ListChargePointsRequest{
	Page:    1,
	PerPage: 10,
})
if err != nil {
	panic(err)
}
fmt.Println(response)
```

### CLI

#### Login

```
$ monta login --client-id <ID> --client-secret <SECRET>
```

#### `GET /v1/auth/me`

```
$ monta me
```

#### `GET /v1/sites`

```
$ monta sites
```

#### `GET /v1/charge-points`

```
$ monta charge-points
```
