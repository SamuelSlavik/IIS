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

    static get signup(): string {
        return `${Endpoints.baseUrl}/users/signup`;
    }

    static get retrieveUser(): string {
        return `${Endpoints.baseUrl}/users/get`;
    }
}