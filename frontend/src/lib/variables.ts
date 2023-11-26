export class Endpoints {
    private static readonly baseUrl = "http://localhost:8080/api"; // Replace with your base URL

    static get connections(): string {
        return `${Endpoints.baseUrl}/connections`;
    }
    static get connection(): string {
        return `${Endpoints.baseUrl}/connections/`;
    }


    static get login(): string {
        return `${Endpoints.baseUrl}/users/login`;
    }
    static get logout(): string {
        return `${Endpoints.baseUrl}/users/logout`;
    }
    static get signup(): string {
        return `${Endpoints.baseUrl}/users/signup`;
    }
    static deleteUser(id: string): string {
        return `${Endpoints.baseUrl}/users/delete/${id}`;
    }
    static updateUser(id: string): string {
        return `${Endpoints.baseUrl}/users/update/${id}`;
    }
    static get retrieveCurrentUser(): string {
        return `${Endpoints.baseUrl}/users/get`;
    }
    static retrieveUser(id: string): string {
        return `${Endpoints.baseUrl}/users/get/${id}`;
    }
    static listUsers(query?: string): string {
        if (query) {
            return `${Endpoints.baseUrl}/users/list?query=${query}`;
        }
        return `${Endpoints.baseUrl}/users/list`;
    }




    static get listVehicles(): string {
        return `${Endpoints.baseUrl}/vehicles/list`;
    }
    static get createVehicle(): string {
        return `${Endpoints.baseUrl}/vehicles/create`;
    }
    static retrieveVehicle(id: string): string {
        return `${Endpoints.baseUrl}/vehicles/get/${id}`;
    }
    static updateVehicle(id: string): string {
        return `${Endpoints.baseUrl}/vehicles/update/${id}`;
    }
    static deleteVehicle(id: string): string {
        return `${Endpoints.baseUrl}/vehicles/delete/${id}`;
    }




    static listStops(query?: string): string {
        if (query) {
            return `${Endpoints.baseUrl}/stops?query=${query}`;
        }
        return `${Endpoints.baseUrl}/stops`;
    }
    static listStopsByQuery(query:string): string {
        return `${Endpoints.baseUrl}/stops/${query}`;
    }
    static stopDetail(id: string): string {
        return `${Endpoints.baseUrl}/stops/get/${id}`;
    }
    static editStop(id: string): string {
        return `${Endpoints.baseUrl}/stops/edit/${id}`;
    }
    static deleteStop(id: string): string {
        return `${Endpoints.baseUrl}/stops/delete/${id}`;
    }
    static get createStop(): string {
        return `${Endpoints.baseUrl}/stops/create`;
    }



    static get listLines(): string {
        return `${Endpoints.baseUrl}/lines`;
    }



    static get reportMalfunction(): string {
        return `${Endpoints.baseUrl}/maintenance/malfunc/create`;
    }
    static retrieveMalfunction(id: string): string {
        return `${Endpoints.baseUrl}/maintenance/malfunc/get/${id}`;
    }
    static editMalfunction(id: string): string {
        return `${Endpoints.baseUrl}/maintenance/malfunc/update/${id}`;
    }
    static deleteMalfunction(id: string): string {
        return `${Endpoints.baseUrl}/maintenance/malfunc/delete/${id}`;
    }
    static listMalfunctions(status?: string): string {
        if (status) {
            return `${Endpoints.baseUrl}/maintenance/malfunc/list?status=${status}`;
        }
        return `${Endpoints.baseUrl}/maintenance/malfunc/list`;
    }
}