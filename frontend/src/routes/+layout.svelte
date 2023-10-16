<script lang="ts">
	import '../app.postcss';
	import { AppShell, AppBar, AppRail, AppRailTile, AppRailAnchor } from '@skeletonlabs/skeleton';

	// Floating UI for Popups
	import { computePosition, autoUpdate, flip, shift, offset, arrow } from '@floating-ui/dom';
	import { storePopup } from '@skeletonlabs/skeleton';
	storePopup.set({ computePosition, autoUpdate, flip, shift, offset, arrow });

    let currentTile: number = 0;
    let visible: boolean = true;

    function toggleVisible() {
        visible = !visible;
    }
</script>

<!-- App Shell -->
<AppShell class="flex flex-row">
		<!-- App Bar -->
		<AppRail class="h-full w-[3%] fixed" background="bg-transparent" >
			<svelte:fragment slot="lead">
                <div class="text-center justify-center items-center" on:click={toggleVisible}>
                    <AppRailAnchor href="/"> <box-icon name='menu' size='sm', color='#CDF0F6'></box-icon></AppRailAnchor>
                </div>
			</svelte:fragment>
            {#if visible}
                <AppRailTile bind:group={currentTile} name="home" value={0} title="home" class="my-2 hover:scale-110 ease-in duration-300 hover:my-4">
		            <svelte:fragment slot="lead"><box-icon type='solid' name='home' size='sm', color='#CDF0F6'></svelte:fragment>
                    <span>Home</span>
	            </AppRailTile>
	            <AppRailTile bind:group={currentTile} name="profile" value={1} title="profile" class="my-2 hover:scale-110 ease-in duration-300 hover:my-4">
                    <svelte:fragment slot="lead"><box-icon name='user-account' type='solid' size='sm', color='#CDF0F6'></box-icon></svelte:fragment>
                    <span>Account</span>
	            </AppRailTile>
	            <AppRailTile bind:group={currentTile} name="transfer" value={2} title="transfer" class="my-2 hover:scale-110 ease-in duration-300 hover:my-4">
		            <svelte:fragment slot="lead"><box-icon name='transfer' size='sm', color='#CDF0F6'></box-icon></svelte:fragment>
		            <span>Transfer</span>
                </AppRailTile>
            {/if}
		</AppRail>
	<!-- Page Route Content -->
	<slot />
</AppShell>
