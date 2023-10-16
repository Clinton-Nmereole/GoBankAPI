import { redirect } from "@sveltejs/kit";
import type { Actions } from "./$types";

export const actions: Actions = {
    createAccount: async ({ request }) => {
            
        const { first_name, last_name, password } = Object.fromEntries(await request.formData()) as { first_name: string, last_name: string, password: string }
    
        const response = await fetch('http://0.0.0.0:8080/accounts', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                first_name,
                last_name,
                password
            })
        })
    
        if (response.ok) {
            console.log('Account created')
            throw redirect(303, '/')
        }
        else {
            console.log('Account not created because of', response.status)
            return {
                status: 400
            }
        }
    
    }
};
