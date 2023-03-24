import { defineStore } from 'pinia'
import {API} from "../helpers/api";

// Store used to keep track of the admin token and other admin only related data
export const useAdminStore = defineStore('admin', {
    state: () => ({
        token: ''
    }),
    actions: {
        // login tries the login the user with the given credentials
        async login(user: string, password: string): Promise<boolean> {
            const form = new URLSearchParams();
            form.append('user', user);
            form.append('password', password);
            const {response, ok} = await API.Req('POST', '/api/login', {body: form});
            if (ok) this.token = response;
            return ok
        }
    }
})