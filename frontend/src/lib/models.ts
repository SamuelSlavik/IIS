export interface Connection {
    ID: string;
    LineName: string;
    Type: string;
    ListStops: ConnectionStop[];
}

export interface ConnectionStop {
    DepartureTime: string;
    StopName: string;
}

export interface UserRegistration {
    firstName: string;
    lastName: string;
    birthDate: string;
    email: string;
    password: string;
    passwordRpt: string;
    role: string;
}

export interface User {
    ID: string;
    FirstName: string;
    LastName: string;
    BirthDate: string;
    Email: string;
    Role: string;
}
