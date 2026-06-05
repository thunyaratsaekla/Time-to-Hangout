<script setup>
import { RouterLink } from 'vue-router'
import { ref } from 'vue'
import { useRouter } from 'vue-router'

import { useAdminUserStore } from '@/stores/admin/user'

const router = useRouter()
const adminUserStore = useAdminUserStore()

const formData = ref({
    firstname: "",
    lastname: "",
    email: "",
    password: "",
    phone: "",
    gender: "",
    birthday: "",
    occupation: "",
    salary: "",
})

const errorMessage = ref('')

function isValidEmail(email) {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
    return emailRegex.test(email)
}

function isValidPhone(phone) {
    const phoneRegex = /^[0-9]{10}$/ // ตัวอย่าง: เบอร์โทรศัพท์ต้องมี 10 หลัก
    return phoneRegex.test(phone)
}

async function handleRegister() {
    errorMessage.value = ''
    if (
        !formData.value.firstname ||
        !formData.value.lastname ||
        !formData.value.email ||
        !formData.value.password ||
        !formData.value.phone ||
        !formData.value.gender ||
        !formData.value.birthday ||
        !formData.value.occupation ||
        !formData.value.salary) {
        errorMessage.value = 'Please fill in all fields.'
        return
    }

    if (!isValidEmail(formData.value.email)) {
        errorMessage.value = 'Please enter a valid email address.'
        return
    }

    if (!isValidPhone(formData.value.phone)) {
        errorMessage.value = 'Please enter a valid phone number (10 digits).'
        return
    }

    if (isChecked.value) {
        try {
            await adminUserStore.registerUser(formData.value)
            alert("User registered successfully!")
            router.push({ name: "admin-login" })
        } catch (error) {
            console.log(error)
        }
    }
}

const isChecked = ref(false)
</script>
<template>
    <div
      class="min-h-screen flex items-center bg-center bg-no-repeat bg-cover"
      style="
        background-image: url('https://png.pngtree.com/thumb_back/fh260/background/20220609/pngtree-colorful-gradient-pink-yellow-pastel-image_1412650.jpg');
      "
    >
      <div class="flex-1 max-w-3xl shadow-2xl mx-auto my-auto p-8">
        <div class="flex flex-col items-center">
          <h1 class="text-4xl">Register</h1>
          <div class="w-full mt-4 flex flex-col gap-4"></div>
          <fieldset class="fieldset w-full mt-8">
            <legend class="fieldset-legend">First Name :</legend>
            <input
              v-model="formData.firstname"
              type="text"
              class="input input-bordered w-full"
              placeholder="First Name"
            />
          </fieldset>
          <fieldset class="fieldset w-full mt-4">
            <legend class="fieldset-legend">Last Name :</legend>
            <input
              v-model="formData.lastname"
              type="text"
              class="input input-bordered w-full"
              placeholder="Last Name"
            />
          </fieldset>
          <fieldset class="fieldset w-full mt-4">
            <legend class="fieldset-legend">Email :</legend>
            <input
              v-model="formData.email"
              type="email"
              class="input input-bordered w-full"
              placeholder="Email"
            />
          </fieldset>
          <fieldset class="fieldset w-full mt-4">
            <legend class="fieldset-legend">Password :</legend>
            <input
              v-model="formData.password"
              type="password"
              class="input input-bordered w-full"
              placeholder="Password"
            />
          </fieldset>
          <fieldset class="fieldset w-full mt-4">
            <legend class="fieldset-legend">Phone Number :</legend>
            <input
              v-model="formData.phone"
              type="text"
              class="input input-bordered w-full"
              placeholder="Phone Number"
            />
          </fieldset>
          <fieldset class="fieldset w-full mt-4">
            <legend class="fieldset-legend">Gender :</legend>
            <input
              v-model="formData.gender"
              type="text"
              class="input input-bordered w-full"
              placeholder="Gender"
            />
          </fieldset>
          <fieldset class="fieldset w-full mt-4">
            <legend class="fieldset-legend">BirthDay :</legend>
            <input
              v-model="formData.birthday"
              type="text"
              class="input input-bordered w-full"
              placeholder="Birthday"
            />
          </fieldset>
          <fieldset class="fieldset w-full mt-4">
            <legend class="fieldset-legend">Occupation :</legend>
            <input
              v-model="formData.occupation"
              type="text"
              class="input input-bordered w-full"
              placeholder="Occupation"
            />
          </fieldset>
          <fieldset class="fieldset w-full mt-4">
            <legend class="fieldset-legend">Salary :</legend>
            <input
              v-model="formData.salary"
              type="text"
              class="input input-bordered w-full"
              placeholder="Salary"
            />
          </fieldset>
          <fieldset class="fieldset w-full mt-4">
            <input v-model="isChecked" type="checkbox" class="checkbox mx-1" />
            <span class="text-sm">
              I agree to the Terms of Service and Privacy Policy
            </span>
          </fieldset>
          <div v-if="errorMessage" class="alert alert-error mt-4">
            {{ errorMessage }}
          </div>
          <button
            @click="handleRegister"
            class="btn btn-primary w-full mt-4"
            :class="{ 'btn-disabled': !isChecked }"
            :disabled="!isChecked"
          >
            Create
          </button>
          <div class="flex items-center justify-center mt-4 space-x-1 text-sm">
            <p class="label">Already have an account?</p>
            <RouterLink :to="{ name: 'user-login' }" class="font-bold"
              >Sign in</RouterLink
            >
          </div>
        </div>
      </div>
    </div>
  </template>