<script lang="ts">
	import { AppRail } from '@skeletonlabs/skeleton';
	import type { PageData } from './$types';
	import { goto, invalidateAll } from '$app/navigation';

	export let data: PageData;
	let inputDemo = '';
	let points = 5701;
	console.log(data);

	async function logout() {
		await fetch(`http://0.0.0.0:8080/logout`, { mode: 'no-cors' });
		console.log('logout successful');
		invalidateAll();
	}
</script>

<div class="h-screen w-full flex">
	<AppRail class="h-screen w-[13%]">
		<h1 class="text-3xl text-center mr-10 mt-6 mb-12">Dashboard</h1>
		<button
			class="flex my-10 items-center hover:underline hover:scale-110 hover:cursor-pointer hover:text-primary-500"
		>
			<box-icon name="home" size="md" color="white" class="ml-6" />
			<h1 class="text-2xl font-light text-center mx-4">Overview</h1>
		</button>
		<button
			class="flex my-10 items-center hover:underline hover:scale-110 hover:cursor-pointer hover:text-primary-500"
		>
			<box-icon name="pie-chart" size="md" color="white" class="ml-6" />
			<h1 class="text-2xl font-light text-center mx-4">Payments</h1>
		</button>
		<button
			class="flex my-10 items-center hover:underline hover:scale-110 hover:cursor-pointer hover:text-primary-500"
		>
			<box-icon name="money-withdraw" size="md" color="white" class="ml-6" />
			<h1 class="text-2xl font-light text-center mx-4">Cashback</h1>
		</button>
		<button
			class="flex my-10 items-center mb-48 hover:underline hover:scale-110 hover:cursor-pointer hover:text-primary-500"
		>
			<box-icon name="briefcase" size="md" color="white" class="ml-6" />
			<h1 class="text-2xl font-light text-center mx-4">Exchange</h1>
		</button>
		<button
			class="flex my-10 items-center hover:underline hover:scale-110 hover:cursor-pointer hover:text-primary-500"
		>
			<box-icon name="cog" size="md" color="white" class="ml-6" />
			<h1 class="text-2xl font-light text-center mx-4">Settings</h1>
		</button>
		<button
			class="flex my-10 items-center hover:underline hover:scale-110 hover:cursor-pointer hover:text-primary-500"
			type="button"
			on:click={logout}
			><a href="/login" class="flex">
				<box-icon name="log-out" size="md" color="white" class="ml-6" />
				<h1 class="text-2xl font-light text-center mx-4">Logout</h1>
			</a>
		</button>
	</AppRail>

	<div class="w-full flex flex-col ml-24 mt-5">
		<div>
			<input
				class="input w-[35%] h-10"
				type="search"
				name="demo"
				bind:value={inputDemo}
				placeholder="Search..."
			/>
			<div class="flex">
				<h1 class="mt-16 mb-4 text-3xl font-bold font-Satoshi mr-44">Balance</h1>
				<h1 class="mt-16 mb-4 text-3xl font-bold font-Satoshi ml-44">Points</h1>
			</div>
			<div class="flex">
				<div class="card h-48 w-96 items-center flex justify-center mr-10">
					<h1 class="text-5xl font-bold font-Satoshi">${data.user.balance}.00</h1>
				</div>

				<div class="card h-48 w-96 items-center flex justify-center ml-10">
					<h1 class="text-5xl font-bold font-Satoshi">{points}</h1>
				</div>
			</div>
			<div class="flex">
				<h1 class="mt-16 mb-4 text-3xl font-bold font-Satoshi mr-44">Cards</h1>
			</div>
		</div>
	</div>
</div>
