# exclaimation-gql
GraphQL API for my personal website

This is either going to be one of the microservices, or the main server for my personal website depending on how much it grows

I am using GraphQL for the purpose of exploring more in term of backend, web services, apis and friends. I was just using gqlgen so I figured it would be nice to use.

I have plan to add some microservices here and there but I am not going to work on them until I have a concrete goal

### Tech Stack for the Backend
#### gin
I like gin, don't judge me. It's nice to work with, fast, and have good middleware support.
#### gqlgen
gqlgen as mentiond before is the one of the few ways I know to make GraphQL API 
and I prefer it over most other solution.
#### gorp
gorp is something I have been trying out in replacement for GORM, reform, ent or any orm. It provides a lower level API with lightweight packages. 
It still have the issues of empty interfaces but for something light it's a lot more tolerable
#### fx
fx is just good all around for doing structured architecture with dependencies injection that is clean and scalable

### Tech Stack for the entire thing (at least the thing I want to use)
> 1. React
> 2. Typescript
> 3. Go
> 4. Postgres
> 5. GraphQL

```go
func main() {
	fmt.Printf("Thanks for checking this out, %v\n", you)
}
```