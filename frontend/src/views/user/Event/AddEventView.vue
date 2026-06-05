<script setup>
import Userslayout from "@/layouts/Userslayout.vue"
import { reactive, ref } from "vue"
import { useRouter, useRoute, RouterLink } from "vue-router"
import { onMounted } from "vue"

import { useAdminProductStore } from "@/stores/admin/product"
import { useAdminUserStore } from "@/stores/admin/user"

import IconCorrect from "@/components/icons/IconCorrect.vue"


const inputForm = [
  {
    name: "Name Event",
    field: "name",
  },
  {
    name: "Image",
    field: "image",
  },
  {
    name: "Person",
    field: "person",
  },
  {
    name: "Location",
    field: "location",
  },
  {
    name: "Date and time Start",
    field: "timestart",
    type:"datetime-local",
  },
  {
    name: "Date and time End",
    field: "timeend",
    type:"datetime-local",
  },
  {
    name: "Description",
    field: "about",
  },
]

const adminProductStore = useAdminProductStore()
const adminUserStore = useAdminUserStore()

const router = useRouter()

const route = useRoute()

const productIndex = ref(-1)
const errorMessage = ref("")

const mode = ref("Add Event")

const productData = reactive({
  name: "",
  image: "https://img.daisyui.com/images/profile/demo/yellingcat@192.webp",
  person: "",
  location: "",
  timestart: "",
  timeend: "",
  about: "",
  status: "Unfinished",
  point: "1",
})

console.log("productData" , productData)

function formatDateForBackend(dateString) {
  const date = new Date(dateString)
  console.log(date)
  return date.toISOString() // แปลงเป็น ISO 8601
}

const SubmitProduct = async () => {
  errorMessage.value = ""

  if (
    !productData.name ||
    !productData.person ||
    !productData.location ||
    !productData.timestart ||
    !productData.timeend ||
    !productData.about 
  ) {
    errorMessage.value = "Please fill in all required fields."
    return
  }

  if (productData.timestart) {
    productData.timestart = formatDateForBackend(productData.timestart)
    console.log("Formatted Date Start:", productData.timestart)
  }
  if (productData.timeend) {
    productData.timeend = formatDateForBackend(productData.timeend)
    console.log("Formatted Date End:", productData.timeend)
  }

  const userId = adminUserStore.userId || localStorage.getItem("userId")
  const bodyEvent = {
    ...productData,
    user_id: Number(userId),
  }

  try {
    if (mode.value === "Edit Event") {
      await adminProductStore.updateEventUser(productIndex.value, productData)
    } else {
      console.log(bodyEvent)
      await adminProductStore.addEventUser(bodyEvent)
    }

    router.push({ name: "user-myhangout" }).then(() => {
      adminProductStore.loadEventsUser()
    })
    console.log("productData", productData)
  } catch (error) {
    console.error("Error submitting product:", error)
    errorMessage.value = "An error occurred while submitting the product."
  }
}

onMounted(async () => {
  console.log("route.params.id", route.params.id)
  if (route.params.id) {
    productIndex.value = parseInt(route.params.id)
    if (isNaN(productIndex.value)) {
      console.log("Invalid productIndex", route.params.id)
    }
    mode.value = "Edit Event"

    const SelectedProduct = await adminProductStore.loadEventsUser(
      productIndex.value
    )
    console.log("SelectedProduct", SelectedProduct)
    productData.name = SelectedProduct.name
    productData.image = SelectedProduct.image
    productData.person = SelectedProduct.person
    productData.location = SelectedProduct.location
    productData.timestart = SelectedProduct.timestart
    productData.timeend = SelectedProduct.timeend
    productData.about = SelectedProduct.about
    productData.id = SelectedProduct.id
  }
})
</script>

<template>
  <Userslayout>
    <div class="shadow-xl p-10">
      <div class="text-3xl font-bold">{{ mode }}</div>
      <div class="divider"></div>
      <fieldset class="fieldset py-4 grid grid-cols-2 gap-4">
        <div v-for="form in inputForm">
          <legend class="fieldset-legend">{{ form.name }} :</legend>
          <input
            v-if="form.type === 'datetime-local'"
            v-model="productData[form.field]"
            type="datetime-local"
            class="input input-info w-full"
          />
          <input
            v-else
            v-model="productData[form.field]"
            type="text"
            class="input input-info w-full"
            placeholder="Enter text"
          />
        </div>
      </fieldset>
      <div class="divider"></div>
      <div v-if="errorMessage" class="alert alert-error mt-4">
        {{ errorMessage }}
      </div>
      <div class="flex justify-end gap-4 mt-4">
        <div>
          <RouterLink :to="{ name: 'user-myhangout' }" class="btn btn-ghost"
            >Back</RouterLink
          >
        </div>
        <div>
          <button @click="SubmitProduct()" class="btn btn-success">
            {{ mode }}
          </button>
        </div>
      </div>
    </div>
  </Userslayout>
</template>
