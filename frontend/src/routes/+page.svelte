<script lang="ts">
    import type { PageData } from './$types';

    export let data: PageData;
</script>
<div class="container mt-16 mx-auto">
    <h1 class="text-7xl font-bold h1 text-center">Bank</h1>
    <h3 class="text-3xl font-bold h3 text-center">Accounts</h3>
    {#each data.bank as account}
        <div class="border-2 border-white rounded mt-8">
            <h2 class="text-xl font-bold h2 text-center">First Name: {account.first_name}</h2>
            <h2 class="text-xl font-bold h2 text-center">Last Name: {account.last_name}</h2>
            <h2 class="text-xl font-bold h2 text-center">Account Number: {account.number}</h2>
        </div>
    {/each}
</div>
