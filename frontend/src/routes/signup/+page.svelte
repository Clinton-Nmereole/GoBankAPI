<script lang="ts">
</script>

<section class="py-24 lg:py-28 bg-surface-900 h-full w-full overflow-hidden">
	<div class="container px-4 mx-auto">
		<form class="mx-auto max-w-3xl" action="?/createAccount" method="POST">
			<h2 class="text-6xl tracking-tight font-Satoshi py-10">Create an Account</h2>

			<div class="grid grid-flow-row grid-cols-2 grid-rows-2 min-h-fit py-10">
				<div class="md:w-4/5 w-full py-5 px-1">
					<input
						type="text"
						placeholder="First Name"
						class="rounded-lg outline-none tracking-tight
                         text-gray-700 w-full scale-105"
						name="first_name"
					/>
				</div>

				<div class="md:w-4/5 w-full py-5 px-1">
					<input
						type="text"
						placeholder="Last Name"
						class="rounded-lg outline-none tracking-tight
                         text-gray-700 w-full scale-105"
						name="last_name"
					/>
				</div>

				<div class="md:w-4/5 w-full py-5 px-1">
					<input
						type="password"
						placeholder="Password"
						class="rounded-lg outline-none tracking-tight
                         text-gray-700 w-full scale-105"
						name="password"
					/>
				</div>
			</div>
			<div class="py-3 w-full">
				<button
					type="submit"
					class="w-[90%] bg-secondary-500 hover:bg-secondary-700 rounded-lg font-bold inline-block px-5
                    py-4 mb-7">Create Account</button
				>
				<p>
					Already have an account? <a href="/login" class="text-tertiary-400 hover:underline"
						>Sign In</a
					>
				</p>
			</div>
		</form>
	</div>
</section>
