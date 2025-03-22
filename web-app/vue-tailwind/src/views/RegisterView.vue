<script lang="ts" setup >
import { ref } from "vue";
const username = ref("");
const email = ref("");
const password = ref("");
const confirmPass = ref("");

const onLogin = () => {
  //  form validation
  if (!email.value || !password.value) {
    return alert("put some data!");
  }
};
</script>


<template>
    <!-- Register section -->
    <section class="flex justify-center items-center min-h-screen">
      <!-- form -->
      <form class="flex flex-col gap-5 items-center" @submit.prevent="onLogin">
        <span class="text-2xl">REGISTER</span>
        <input type="email" placeholder="Username" v-model="username" class="p-2 outline-none border"/>
        <input type="email" placeholder="Email" v-model="email" class="p-2 outline-none border"/>
        <input type="password" placeholder="Password" v-model="password" class="p-2 outline-none border"/>
        <input type="password" placeholder="Confirm Password" v-model="confirmPass" class="p-2 outline-none border"/>
        <button class="bg-black text-white w-full py-2 cursor-pointer">Submit</button>
        <span>Already Have an Account? <a href="/"></a></span>
      </form>
    </section>
  </template>
  
  
