import type { PageLoad } from './$types';
import { redirect } from "@sveltejs/kit";

export const load: PageLoad = async (event: any) => {
    const login = event.cookies.get('jwt-x-token')
    if (login) {
        throw redirect(303, '/me')
    }
    throw redirect(303, '/login')
};
