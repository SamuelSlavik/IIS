import { defineStore } from 'pinia';
import type { User } from "@/lib/models";

export const useUserStore : any = defineStore(
    'user', {
        state: (): any => ({
            id: '',
            firstName: '',
            lastName: '',
            birthDate: '',
            email: '',
            role: '',
        }),
        actions: {
            logUserIn (token: string | undefined): void {
                this.id = token || '';
            },
            setUserData(user: User): void {
                this.id = user.ID;
                this.firstName = user.FirstName;
                this.lastName = user.LastName;
                this.birthDate = user.BirthDate;
                this.email = user.Email;
                this.role = user.Role;
            },
            logOut (): void {
                this.id = '';
                this.firstName = '';
                this.lastName = '';
                this.birthDate = '';
                this.email = '';
                this.role = '';
            },
            checkAuthentication (): boolean {
                return this.id ? true : false;
            }
        },
    });
