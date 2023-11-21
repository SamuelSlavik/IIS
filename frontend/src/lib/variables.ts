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
}