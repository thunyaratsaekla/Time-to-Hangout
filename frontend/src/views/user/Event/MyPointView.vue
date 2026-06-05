<script setup>
import { onMounted, ref } from "vue"
import { RouterLink, useRouter } from "vue-router"
import { useAdminUserStore } from "@/stores/admin/user"
import Userslayout from "@/layouts/Userslayout.vue"

const adminUserStore = useAdminUserStore()
const userPoint = ref(0)

onMounted(async () => {
    await adminUserStore.loadUser(adminUserStore.userId)
    userPoint.value = adminUserStore.currentUser?.point || 0
})

</script>
<template>
  <Userslayout>
    <div
      class="h-screen flex items-center bg-cover bg-no-repeat bg-center"
      style="
        background-image: url('https://png.pngtree.com/thumb_back/fh260/background/20220609/pngtree-colorful-gradient-pink-yellow-pastel-image_1412650.jpg');
      "
    >
      <div class="flex-1 max-w-3xl shadow-2xl mx-auto p-8">
        <div class="flex flex-col items-center">
          <h1 class="text-4xl">My point</h1>
                    <div class="text-2xl mt-8">Point: <span class="font-bold">{{ userPoint }}</span></div>
        </div>
      </div>
    </div>
  </Userslayout>
</template>
