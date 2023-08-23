# About

**1 - What is it?**

  My first API made in Golang following this [tutorial](https://www.youtube.com/watch?v=MD7b-iQMC24). Golang is a compiled language created by Google in 2007 to have a easy reading and fast perfomance.
  
**2 - How to run?**

  - Clone the repository
  - Create `config.toml` file in root directory repository and fill out with these informations:
  
```
[api]
port=":3001"

[db]
host="localhost"
port="5432"
user=""
pass=""
name=""
```
  - Type `go mod tidy`
  - Type `go run main.go` in root directory
  - Acess `http://localhost:3001`
  
**3 - Routes**

  - `GET` `http://localhost:3001/user`
  
    Return all users
    
  - `GET` `http://localhost:3001/user/${id}`
  
    Return user with `id` specified in url, like `http://localhost:3001/user/1`, it will return user com `id` equal to 1
    
  - `POST` `http://localhost:3000/user`
    
    Insert this informations in body request:
    
    ```
    {
      "name": "user",
      "password": "1234",
      "email": "user@gmail.com"
    }
    ```
    
    And it will create an object in database table `users`.
    
  - `PATCH` `http://localhost:3000/user/${id}`
  
    It will update a specific user with `id` that you passed in params and the informations in body, like:
    
    `http://localhost:3000/user/1`
    
    ```
    {
      "id": 1,
      "name": "user changed",
      "password": "1234",
      "email": "user@gmail.com"
    }
    ```
    
  - `DELETE` `http://localhost:3000/user/${id}`
  
    remove completetly a user with `id` that you specified in URL params.
