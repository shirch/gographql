# gographql

## Steps to run the application:

1. docker run -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=dbpass -e MYSQL_DATABASE=dbname -d mysql:latest
2. go run server.go
3. Using the GraphQL server url http://localhost:8080/graphql, Apollo Server automatically serves the GUI. 
Pass the following mutations and run them to see the results:
```go
query links{
	links{
    title
    address,
  }
}

mutation createLink{
  createLink(input: {title: "something", address: "somewhere"}){
    title,
    address,
    id,
    user{
      name
    }
  }
}

mutation createUser{
  createUser(input: {username: "user", password: "1234"})
}

mutation login{
  login(input: {username: "user", password: "1234"})
}

mutation updateUser{
  updateUser(userId: "btosajevvhfjmbjje9b0" , input: {username: "shir", password: "pass"}){
    name,
    password,
    id,
  }
}

mutation updateLink{
  updateLink(linkId: "btosajevvhfjmbjje9b0" , input: {title: "something", address: "somewhere",userId: "btosajevvhfjmbjje9b0"}){
    title,
    address,
    id,
    user{
      name
    }
  }
}

mutation refreshToken{
  refreshToken(input: {token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDIxNjA0MzAsInVzZXJuYW1lIjoidXNlciJ9.p_SgOQF4jjvWJXi7FXtH9v-ZgteudFJ0UzEaYYN-rFg"})
}
```
