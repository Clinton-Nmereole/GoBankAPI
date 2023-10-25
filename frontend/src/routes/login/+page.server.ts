import { redirect } from "@sveltejs/kit";
import type { PageLoad } from "./$types";
import type { Actions } from "./$types";




export const actions: Actions = {
    loginAccount: async (event) => {
        const formData = await event.request.formData()
        const account_number = formData.get('account_number') as string
        const password = formData.get('password') as string
        //const account_number = Number(account_str)


    
        const response = await fetch('http://0.0.0.0:8080/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                "account_number": Number(account_number),
                "password": password
            })
        })

        console.log(account_number)
    
        if (response.ok) {
            //const x_jwt_token = response.headers.get('x-jwt-token') as string
            const user = response.headers.get('user-id') as string
            //event.cookies.set('x-jwt-token', x_jwt_token, {
             //   path: '/',           
            //})
            event.cookies.set('user-id', user, {
                path: '/',
            })
            console.log('Login successful')
            console.log(user)
            throw redirect(303, '/me')
        }
        else {
            console.log('Login Failed because of', response.status)
            return {
                status: 400
            }
        }
    
    }
};
