import type { PageLoad } from './$types';

export const load: PageLoad = async (event: any) => {
    const user_id = event.cookies.get('user-id') as string;
    console.log(user_id);
    return {
        user: await fetch('http://0.0.0.0:8080/accounts/4').then(data => data.json())
    };
};
