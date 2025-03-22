<script lang="ts">
	let email: string = '';
	let password: string = '';
	import { login } from '$lib/actions/auth';
	// submit function

	const onLogin = async () => {
		// form validation
		if (!email || !password) {
			alert('loko kulang!');
			email = '';
			password = '';
			return;
		}
		// await
		await login(email, password);
	};
</script>

<main class="flex min-h-screen items-center justify-center bg-gray-100">
	<section class="w-full max-w-md rounded-lg bg-white p-8 shadow-lg">
		<h1 class="mb-6 text-center text-2xl font-bold">Login</h1>
		<form class="space-y-4" on:submit={onLogin}>
			<label class="block">
				<span class="text-gray-700">Email</span>
				<input
					type="email"
					bind:value={email}
					placeholder="Enter your email"
					class="w-full rounded-lg border border-gray-300 p-3 focus:ring-2 focus:ring-blue-500 focus:outline-none"
				/>
			</label>

			<label class="block">
				<span class="text-gray-700">Password</span>
				<input
					type="password"
					bind:value={password}
					placeholder="Enter your password"
					class="w-full rounded-lg border border-gray-300 p-3 focus:ring-2 focus:ring-blue-500 focus:outline-none"
				/>
			</label>

			<button
				class="w-full cursor-pointer rounded-lg bg-blue-500 py-3 text-white transition hover:bg-blue-600"
			>
				Login
			</button>
			<span>Doesn't have an Account? <a href="/register" class="underline">Register Here</a></span>
		</form>
	</section>
</main>
