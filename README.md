# go-graphql-example

Example server using Gin and GraphQL API.

`go run *.go`

## Queries

`http://127.0.0.1:12345/graphql?query=query{allPlayers{name}}`

`http://127.0.0.1:12345/graphql?query=query{kClosest(k:5){id}}`

## Mutations

`http://127.0.0.1:12345/graphql?query=mutation{spawn(name:%22a%22,x:100,y:101){id}}`

`http://127.0.0.1:12345/graphql?query=mutation{move(id:%22sfjksdhfkldas%22,x:100,y:101){id}}`

`http://127.0.0.1:12345/graphql?query=mutation{leave(id:%22sfjksdhfkldas%22)}`
