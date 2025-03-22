<script lang="ts" setup >

import { ref } from "vue";
const email = ref("");
const password = ref("");
const onLogin = () => {
  //  form validation
  if (!email.value || !password.value) {
    return alert("put some data!");
  }
  console.log(email.value, password.value);
};
</script>


<template>
    <!-- login section -->
    <section class="flex justify-center items-center min-h-screen">
      <!-- form -->
      <form class="flex flex-col gap-5 items-center" @submit.prevent="onLogin">
        <span class="text-2xl">LOGIN</span>
        <input type="email" placeholder="Email" v-model="email" class="p-2 outline-none border"/>
        <input type="password" placeholder="Password" v-model="password" class="p-2 outline-none border"/>
        <button class="bg-black text-white w-full py-2 cursor-pointer">Submit</button>
      </form>
    </section>
  </template>
  
  
