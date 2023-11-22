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

    static get retrieveCurrentUser(): string {
        return `${Endpoints.baseUrl}/users/get`;
    }

    static retrieveUser(id: string): string {
        return `${Endpoints.baseUrl}/users/get/${id}`;
    }

    static get listUsers(): string {
        return `${Endpoints.baseUrl}/users/list`;
    }

    static get listVehicles(): string {
        return `${Endpoints.baseUrl}/vehicles/list`;
    }

    static get listStops(): string {
        return `${Endpoints.baseUrl}/stops`;
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
}