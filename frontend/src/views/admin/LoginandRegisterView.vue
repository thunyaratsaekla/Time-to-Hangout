<script setup>
import { RouterLink } from 'vue-router'
</script>
<template>
    <div class="h-screen flex items-center bg-cover bg-no-repeat bg-center" style="background-image: url('https://png.pngtree.com/thumb_back/fh260/background/20220609/pngtree-colorful-gradient-pink-yellow-pastel-image_1412650.jpg')">
        <div class="flex-1 max-w-3xl shadow-2xl mx-auto p-8">
            <div class="flex flex-col items-center my-10">
                <div class="avatar">
                <div class="w-lg rounded-full">
                <img src="https://img.daisyui.com/images/profile/demo/yellingcat@192.webp" />
                </div>
                </div>
                <h1 class="text-4xl my-10">Time to Hangout</h1>
                <RouterLink :to="{ name : 'admin-login'}" class="btn btn-primary w-full mt-4">Login</RouterLink>
                <RouterLink :to="{ name : 'admin-register'}" class="btn btn-primary w-full mt-4">Register</RouterLink>
            </div>
        </div>
    </div>
    
</template>