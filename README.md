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
    /api/vehicles/get/:regnum      - retrieve specific vehicle
    CONNECTIONS:
    /api/connections/list/:linename        - retrieve all conncections on line (without driver and vehicle)
    /api/connections/list/:linename/:date      - retrieve all conncections on line at date 
    /api/connections/get/:id    - retrieve connection by id
    LINES:
    /api/lines/list         - list all lines
    /api/lines/get/:name    - get specific line
    DRIVERS:
    /api/drivers/list/:datetime - list drivers free at datetime
    STOPS:
    /api/stops/list       - get all stops
    MALFUNC REPORTS:
    /api/maintenance/malfunc/list     - list all malfunction reports
    /api/maintenance/malfunc/list/:status     - list all malfunction reports with status
    /api/maintenance/malfunc/get/:id     - get specific malfunction request
    MAINTENANCE REQUEST:
    /api/maintenreq/list    - list MAINTENANCE REQUESTs
    /api/maintenreq/list/:status     - list MAINTENANCE REQUEST with status
    /api/maintenreq/list/super/:userid     - list MAINTENANCE REQUEST of superuser with user id
    /api/maintenreq/list/tech/:status/:userid     - list MAINTENANCE REQUESTs with status of technician with user id + requests without technician with status
    /api/maintenreq/get/:id    - get specific MAINTENANCE REQUEST
### POST
    USERAUTH: 
    /api/users/signup       - sign up user
    /api/users/login        - login user
    VEHICLES:
    /api/vehicles/create    - create vehicle
    CONNECTIONS:
    /api/conncections/create - create connection (without vehicle and driver)
    LINES:
    /api/lines/create       - create line + its segments
    STOPS:
    /api/stops/create       - create stop
    MALFUNC REPORTS:
    /api/maintenance/malfunc/create    - create malfunction report
    MAINTENANCE REQUEST:
    /api/maintenreq/create     - create MAINTENANCE REQUEST
### PUT
    LINES:
    /api/lines/update/:name - update line and segments
    VEHICLES:
    /api/vehicles/update/:regnum      - update vehicle
    STOPS:
    /api/stops/update/:id       - update stop
### PATCH
    USER:
    /api/users/update/:id   - update user information (not role) *
    CONNECTIONS:
    /api/conncections/update/:id - update connection (without driver and vehicle)
    /api/conncections/assign/:id - assign driver + vehicle 
    MAINTENANCE REQUEST:
    /api/maintenreq/update/status/:id    - update status of maintenance request
### DELETE
    USER:
    /api/users/delete/:id   - delete user (if admin and only one admin exists do not delete)
    LINES:
    /api/lines/delete/:name - delete line and its segments
    VEHICLES:
    /api/vehicles/delete/:regnum      - delete vehicle
    STOPS:
    /api/stops/delete/:id       - create stop
    CONNECTIONS:
    /api/conncections/delete/:id - delete conncetion
