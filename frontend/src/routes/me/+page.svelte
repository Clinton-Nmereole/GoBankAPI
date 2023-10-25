<script lang="ts">
    import type { PageData } from './$types';

    export let data: PageData;
    console.log(data);
</script>
<div class="container mt-16 mx-auto">
    <h1 class="text-7xl font-bold h1 text-center">Bank</h1>
    <h3 class="text-3xl font-bold h3 text-center">Accounts</h3>
    <div class="flex justify-center">
        <div class="border-2 border-white rounded mt-8 card h-[42rem] w-[36rem] flex flex-col justify-center">
            <h2 class="text-2xl font-bold h2 text-center font-Satoshi py-7">First Name: {data.user.first_name}</h2>
            <h2 class="text-2xl font-bold h2 text-center font-Satoshi py-7">Last Name: {data.user.last_name}</h2>
            <h2 class="text-2xl font-bold h2 text-center font-Satoshi py-7">Account Number: {data.user.number}</h2>
            <h2 class="text-2xl font-bold h2 text-center font-Satoshi py-7">Balance: {data.user.balance}</h2>
            <h2 class="text-2xl font-bold h2 text-center font-Satoshi py-7">Creation Date: {data.user.created_at}</h2>
        </div>
    </div>
</div>
