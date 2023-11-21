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
    Role: Roles | undefined;
}

export enum Roles {
    Admin = "admin",
    Superuser = "superuser",
    Technician = "technician",
    Dispatcher = "dispatcher",
    Driver = "driver",
}

export interface Vehicle {
    ID: string
    Registration: string
    Capacity: number
    Brand: string
    ImageData: string
    VehicleType: VehicleType
    VehicleTypeName: string
    LineName: string
    Connections: Connection[]
}

export interface VehicleType {
    ID: string
    Type: VehicleTypeEnum
}

export enum VehicleTypeEnum {
    Bus = "bus",
    Tram = "tram",
    ObrnenaDodavka = "obrnena_dodavka",
}

