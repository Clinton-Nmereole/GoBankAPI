import type { PageLoad } from './$types';

export const load: PageLoad = async () => {
    return {
        bank: await fetch('http://0.0.0.0:8080/accounts').then(data => data.json())
    };
};
