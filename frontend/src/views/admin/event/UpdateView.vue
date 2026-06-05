<script setup>
import Adminlayout from "@/layouts/Adminlayout.vue"
import { reactive, ref } from "vue"
import { useRouter, useRoute, RouterLink } from "vue-router"
import { onMounted } from "vue"

import { useAdminProductStore } from "@/stores/admin/product"

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
    name: "Organizer",
    field: "organizer",
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
  {
    name: "Point",
    field: "point",
  },
]

const adminProductStore = useAdminProductStore()

const router = useRouter()

const route = useRoute()

const productIndex = ref(-1)
const errorMessage = ref("")

const mode = ref("Add Event")

const productData = reactive({
  name: "",
  image: "https://img.daisyui.com/images/profile/demo/yellingcat@192.webp",
  organizer: "",
  location: "",
  timestart: "",
  timeend: "",
  about: "",
  status: "",
  point: "",
})

function formatDateForBackend(dateString) {
  const date = new Date(dateString)
  console.log(date)
  return date.toISOString() // แปลงเป็น ISO 8601
}

function formatDateForInput(dateString) {
  const date = new Date(dateString)
  const year = date.getFullYear()
  const month = String(date.getMonth() + 1).padStart(2, "0")
  const day = String(date.getDate()).padStart(2, "0")
  const hours = String(date.getHours()).padStart(2, "0")
  const minutes = String(date.getMinutes()).padStart(2, "0")
  return `${year}-${month}-${day}T${hours}:${minutes}`
}

const SubmitProduct = async () => {
  errorMessage.value = ""

  if (
    !productData.name ||
    !productData.organizer ||
    !productData.location ||
    !productData.timestart ||
    !productData.timeend ||
    !productData.about ||
    !productData.status ||
    !productData.point
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
  try {
    if (mode.value === "Edit Event") {
      await adminProductStore.updateEvent(productIndex.value, productData)
    } else {
      await adminProductStore.addEvent(productData)
    }

    router.push({ name: "admin-product-list" }).then(() => {
      adminProductStore.loadEvents()
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

    const SelectedProduct = await adminProductStore.loadEvent(
      productIndex.value
    )
    console.log("SelectedProduct", SelectedProduct)
    productData.name = SelectedProduct.name
    productData.image = SelectedProduct.image
    productData.organizer = SelectedProduct.organizer
    productData.location = SelectedProduct.location
    productData.timestart = formatDateForInput(SelectedProduct.timestart)
    productData.timeend = formatDateForInput(SelectedProduct.timeend)
    productData.about = SelectedProduct.about
    productData.status = SelectedProduct.status
    productData.id = SelectedProduct.id
     productData.point = SelectedProduct.point
  }
})
</script>

<template>
  <Adminlayout>
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
      <fieldset class="fieldset">
        <legend class="fieldset-legend">Status :</legend>
        <select v-model="productData.status" class="select input-info">
          <option disabled selected>Choose Status</option>
          <option value="Open">Open</option>
          <option value="Close">Close</option>
        </select>
      </fieldset>
      <div v-if="errorMessage" class="alert alert-error mt-4">
        {{ errorMessage }}
      </div>
      <div class="flex justify-end gap-4 mt-4">
        <div>
          <RouterLink :to="{ name: 'admin-product-list' }" class="btn btn-ghost"
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
  </Adminlayout>
</template>
