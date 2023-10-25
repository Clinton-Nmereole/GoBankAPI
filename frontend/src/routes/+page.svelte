<script lang="ts">
    import type { PageData } from './$types';

    export let data: PageData;
</script>
<div class="container mt-16 mx-auto">
    <h1 class="text-8xl font-bold h1 text-center font-Satoshi">Bank</h1>
    <h3 class="text-4xl font-bold h3 text-center font-Satoshi">Accounts</h3>
    <div class="grid grid-flow-row grid-cols-4 grid-rows-5 min-h-fit">
    {#each data.bank as account}
        <div class="card border-2 border-white rounded mt-8 h-96 w-80 mx-4 text-center items-center flex flex-col justify-center">
            <h2 class="text-xl font-bold h2 text-center font-Satoshi">First Name: {account.first_name}</h2>
            <h2 class="text-xl font-bold h2 text-center font-Satoshi">Last Name: {account.last_name}</h2>
            <h2 class="text-xl font-bold h2 text-center font-Satoshi">Account Number: {account.number}</h2>
        </div>
    {/each}
    </div>
</div>
