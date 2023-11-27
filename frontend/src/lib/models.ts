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
    Registration: string
    Capacity: number
    Brand: string
    Type: string
    LastMaintenance: LastMaintenance
}
export interface NewVehicle {
    Registration: string
    Capacity: number | null
    Brand: string
    Type: string
}
export interface UpdateVehicle {
    Capacity: number | null
    Brand: string
    Type: string
    LineName: string
}

export interface VehicleInList {
    Registration: string
    Capacity: number
    Brand: string
    Type: string
    LastMaintenance: LastMaintenance
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


export interface Stop {
    ID: string
    Name: string
    Active: string
}

export interface NewStop {
    Name: string
}

export interface LineInList {
    Name: string
    InitialStop: string
    FinalStop: string
}
export interface NewLine {
    Name: string
    StopsSequence: SeqStop[]
}
export interface SeqStop {
    StopName: string
    Duration: number | null
}


export interface MalfunctionReport {
    Title: string
    Description: string
    VehicleRef: string
}
export interface Malfunction {
    ID: string
    Title: string
    Description: string
    Vehicle: VehicleRef
    CreatedAt: string
    CreatedBy: CreatedByRef
    Acknowledged: boolean
}
export interface VehicleRef {
    Registration: string
    Brand: string
    VehicleTypeName: string
}
export interface CreatedByRef {
    FirstName: string
    LastName: string
    Email: string
    Role: string
}
export interface LastMaintenance {
    Status: string
    Date: string
}


export interface NewRequest {
    Status: string
    Deadline: string
    MalfuncRepRef: string
    CreatedByRef: string
    ResolvedByRef: string
}

export interface ConnectionList {
    ConnectionID: string
    LineName: string
    InitialStop: string
    DepartureTime: string
    FinalStop: string
    ArrivalTime: string
    Direction: boolean
    VehicleReg: string | null
}
