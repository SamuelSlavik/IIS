# IIS
Project for VUT FIT IIS

https://github.com/cosmtrek/air/blob/master/README.md


## Endpoints
#### Legend
\* - user auth required

### GET
    USERAUTH: 
    /api/users/list         - retrieve all users *
    /api/users/get/:id      - retrieve single user by id *
    /api/users/get          - retrieve currently logged in user *
    VEHICLES:
    /api/vehicles/list      - retrieve all vehicles
    CONNECTIONS:
    /api/connections        - retrieve all conncections???
    /api/connections/:id    - retrieve connection by id
### POST
    USERAUTH: 
    /api/users/signup       - sign up user
    /api/users/login        - login user
    VEHICLES:
    /api/vehicles/create    - create vehicle
### PUT
### PATCH
### DELETE
