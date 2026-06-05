<script setup>
import { ref } from 'vue';
import { RouterLink , useRouter } from 'vue-router'
import { useAdminUserStore } from '@/stores/admin/user';

const adminUserStore = useAdminUserStore()
const router = useRouter()

const errorMessage = ref('')

const formData = ref({
    email : '',
    password: '',
})

function isValidEmail(email) {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
    return emailRegex.test(email)
}


const handleLogin = async () => {
    errorMessage.value = ''
    if (!formData.value.email || !formData.value.password) {
        errorMessage.value = 'Please fill in all fields'
        return
    }

    if (!isValidEmail(formData.value.email)) {
        errorMessage.value = 'Please enter a valid email address'
        return
    }

    try {
        await adminUserStore.loginUser(formData.value)
        alert("Login successful")
        router.push({ name: "admin-product-list" })
    } catch (error) {
        if (error.response?.status === 401) {
            errorMessage.value = 'Email or Password is incorrect.'
        }else {
            errorMessage.value = error.response?.data?.message || 'Invalid email or password.'
        }
        
    }
    
}
</script>
<template>
    <div class="h-screen flex items-center bg-cover bg-no-repeat bg-center" style="background-image: url('https://png.pngtree.com/thumb_back/fh260/background/20220609/pngtree-colorful-gradient-pink-yellow-pastel-image_1412650.jpg')">
        <div class="flex-1 max-w-3xl shadow-2xl mx-auto p-8" >
            <div class="flex flex-col items-center" >
                <h1 class="text-4xl">Login</h1>
                <fieldset class="fieldset w-full mt-8">
                <legend class="fieldset-legend">Email</legend>
                <input v-model="formData.email" type="email" class="input input-bordered w-full" placeholder="Email" />
                </fieldset>
                <fieldset class="fieldset w-full mt-4">
                <legend class="fieldset-legend">Password</legend>
                <input v-model="formData.password" type="password" class="input input-bordered w-full" placeholder="Password" />
                </fieldset>
                <div v-if="errorMessage" class="alert alert-error mt-4">
                    {{ errorMessage }}
                </div>
                <button @click="handleLogin" class="btn btn-primary w-full mt-4">Login</button>
                <div class="flex items-center justify-center mt-4 space-x-1 text-sm">
                <p class="label">Dont have an account?</p>
                <RouterLink :to="{ name : 'admin-register'}" class="font-bold">Register</RouterLink>
                </div>
                
            </div>
        </div>
    </div>
    
</template>