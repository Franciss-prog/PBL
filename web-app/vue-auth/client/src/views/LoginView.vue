<script setup lang="ts">
import axios from "axios";
import { ref } from "vue";
import { errorMessage, successMessage } from "../utils/message";

const email = ref("");
const password = ref("");

const onLogin = async () => {
  // form validation
  if (!email.value || !password.value) {
    errorMessage({ message: "Fill out all Fields", icon: "warning" });

    return;
  }
  try {
    const response = await axios.post("http://localhost:3000/api/login", {
      email: email.value,
      password: password.value,
    });

    if (response.status === 200) {
      // alert the user
      successMessage({ message: response.data.message });
      window.location.href = "/dashboard";
    }
  } catch (error) {
    if (axios.isAxiosError(error)) {
      errorMessage({ message: error.response?.data.error, icon: "error" });
    }
  }
};
</script>
<template>
  <section class="flex justify-center items-center min-h-screen">
    <form action="" class="flex flex-col gap-5" @submit.prevent="onLogin">
      <span class="text-2xl">Login</span>
      <input
        type="email"
        class="outline-none border p-2"
        placeholder="Email"
        v-model="email"
      />
      <input
        type="password"
        class="outline-none border p-2"
        placeholder="Password"
        v-model="password"
      />
      <button class="bg-black w-full p-2 text-white cursor-pointer">
        Login
      </button>
      <!-- doesnt have an account? -->
      <span
        >Doesn't have an Account?
        <a href="/register" class="underline">Register Here</a></span
      >
    </form>
  </section>
</template>
