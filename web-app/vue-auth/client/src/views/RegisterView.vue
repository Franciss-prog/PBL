<script setup lang="ts">
import { ref, withModifiers } from "vue";
import { errorMessage, successMessage } from "../utils/message";
import axios from "axios";

const username = ref("");
const email = ref("");
const password = ref("");
const confirmPassword = ref("");
const showPassword = ref(false);
const onRegister = async () => {
  // form validation
  if (
    !username.value ||
    !email.value ||
    !password.value ||
    !confirmPassword.value
  ) {
    return errorMessage({ message: "Fill out all Fields", icon: "warning" });
  }
  // compare password and confirmPasword
  if (password.value !== confirmPassword.value) {
    // clear the form value
    password.value = "";
    confirmPassword.value = "";
   errorMessage({
      message: "Incorrect Email and Password..",
      icon: "warning",
    });
  }

  // post the data to server
  try {
    const response = await axios.post("http://localhost:3000/api/register", {
      username: username.value,
      email: email.value,
      password: password.value,
    });

    if (response.status === 200 || response.status == 201) {
      successMessage({ message: response.data.message , icon : "success"});
      return (window.location.href = "/dashboard");
    }
  } catch (error) {
    if (axios.isAxiosError(error)) {
      errorMessage({ message: error.response?.data.error, icon: "error" });
    }
  }
};
const togglePass = () => {
  showPassword.value = !showPassword.value;
};
</script>
<template>
  <section class="flex justify-center items-center min-h-screen">
    <form action="" class="flex flex-col gap-5" @submit.prevent="onRegister">
      <span class="text-2xl">Register</span>
      <!-- username -->
      <input
        type="text"
        class="outline-none border p-2"
        placeholder="Username"
        v-model="username"
      />
      <input
        type="email"
        class="outline-none border p-2"
        placeholder="Email"
        v-model="email"
      />
      <input
        :type="showPassword ? 'text' : 'password'"
        class="outline-none border p-2"
        placeholder="Password"
        v-model="password"
      /><input
        :type="showPassword ? 'text' : 'password'"
        class="outline-none border p-2"
        placeholder="Confirm Password"
        v-model="confirmPassword"
      />
      <div class="flex items-center">
        <button class="bg-black w-full p-2 text-white cursor-pointer">
          Register
        </button>
        <button type="button"
          class="text-black w-full p-2 bg-white cursor-pointer"
          @click="togglePass"
        >
          {{ showPassword ?  "show": "hide" }} Passwords
        </button>
      </div>
      <!-- Already have an account? -->
      <span
        >Already have an Account?
        <a href="/" class="underline">Login Here</a></span
      >
    </form>
  </section>
</template>
