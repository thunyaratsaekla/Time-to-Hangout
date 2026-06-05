<script setup>
import Userslayout from "@/layouts/Userslayout.vue"
</script>
<template>
  <Userslayout>
    <div
      class="h-screen flex items-center justify-center bg-cover bg-no-repeat bg-center"
      style="
        background-image: url('https://png.pngtree.com/thumb_back/fh260/background/20220609/pngtree-colorful-gradient-pink-yellow-pastel-image_1412650.jpg');
      "
    >
        <div class="flex flex-col items-center justify-center h-full">
          <div class="avatar">
            <div class="w-lg rounded-full">
              <img
                src="https://img.daisyui.com/images/profile/demo/yellingcat@192.webp"
              />
            </div>
          </div>
          <h1 class="text-4xl my-10">Time to Hangout</h1>
        </div>
    </div>
  </Userslayout>
</template>
