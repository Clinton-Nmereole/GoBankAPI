import type { PageLoad } from './$types';

export const load: PageLoad = async (event: any) => {
    const userid = event.cookies.get('user-id')
    //const account = await fetch('http://0.0.0.0:8080/accounts', {
       // verbose: true,
     //   method: 'GET',
    //})
    //console.log(account.headers.get('x-jwt-token'))
    return {
        user: await fetch('http://0.0.0.0:8080/accounts/' + userid, {verbose: true}).then(data => data.json())
    };
 
};
