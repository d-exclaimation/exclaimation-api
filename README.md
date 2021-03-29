# exclaimation-api
GraphQL API for my personal website

This is either going to be one of the microservices, or the main server for my personal website depending on how much it grows

I am using GraphQL for the purpose of exploring more in term of backend, web services, apis and friends. I was just using gqlgen so I figured it would be nice to use.

I have plan to add some microservices here and there but I am not going to work on them until I have a concrete goal

### Tech Stack for the Backend
#### echo
I decided to use echo instead of gin, the slight performance gain is not worth the modularity with echo.
Might probably switch again to chi but that's unlikely
#### gqlgen
gqlgen as mentiond before is the one of the few ways I know to make GraphQL API 
and I prefer it over most other solution.
#### ent
ent is a code generation entity manager. I like this library since it has type safety for ORM and
using a query builder which is more clear
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