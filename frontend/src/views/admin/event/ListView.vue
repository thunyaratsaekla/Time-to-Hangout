<script setup>
import Adminlayout from "@/layouts/Adminlayout.vue"
import { ref } from "vue"
import { onMounted } from "vue"
import { useRoute, RouterLink } from "vue-router"
import { useAdminProductStore } from "@/stores/admin/product"

import Edit from "@/components/icons/Edit.vue"
import Trash from "@/components/icons/Trash.vue"
import Table from "@/components/Table.vue"
import IconCorrect from "@/components/icons/IconCorrect.vue"
import IconIncorrect from "@/components/icons/IconIncorrect.vue"

const adminProductStore = useAdminProductStore()

onMounted(async () => {
  // adminProductStore.loadProducts()
  await adminProductStore.loadEvents()
  console.log("List", adminProductStore.list)
})

const DeleteEvent = async (eventID) => {
  await adminProductStore.deleteEvent(eventID)
  await adminProductStore.loadEvents()
}

function formatTimeWithoutTimezone(dateString) {
  const date = new Date(dateString)
  const year = date.getFullYear()
  const month = String(date.getMonth() + 1).padStart(2, "0")
  const day = String(date.getDate()).padStart(2, "0")
  const hours = String(date.getHours()).padStart(2, "0")
  const minutes = String(date.getMinutes()).padStart(2, "0")
  const seconds = String(date.getSeconds()).padStart(2, "0")
  return `${year}-${month}-${day} Time ${hours}:${minutes}:${seconds}`
}

const GivePoint = async (eventID) => {
  await adminProductStore.givePoint(eventID)
  await adminProductStore.deleteEvent(eventID)
  await adminProductStore.loadEvents()
}
</script>

<template>
  <Adminlayout>
    <div class="shadow-md rounded-box p-4">
      <div class="flex justify-between items-center">
        <div class="text-4xl font-bold">Event</div>
        <div>
          <RouterLink
            :to="{ name: 'admin-product-create' }"
            class="btn btn-lg btn-primary"
            >Add Event</RouterLink
          >
        </div>
      </div>
      <div class="divider"></div>
      <Table
        :headers="['Name', 'Image', 'Organizer', 'Location', 'Status']"
      >
        <!-- row 1 -->
        <tr v-for="(event, index) in adminProductStore.list">
          <th>{{ event.name }}</th>
          <td>
            <div class="mask mask-squircle w-24 h-24">
              <img :src="event.image" class="w-full h-full object-cover" />
            </div>
          </td>
          <td>{{ event.organizer }}</td>
          <td>{{ event.location }}</td>
          <td>
            <div
              class="badge"
              :class="event.status === 'Open' ? 'badge-success' : 'badge-error'"
            >
              {{ event.status }}
            </div>
          </td>
          <td>
            <div class="flex w-32 gap-4 justify-center">
              <RouterLink
                :to="{ name: 'admin-product-update', params: { id: event.ID } }"
                class="flex w-12 btn btn-info"
              >
                <Edit></Edit>
              </RouterLink>
              <div
                @click="DeleteEvent(event.ID)"
                class="flex w-12 btn btn-error"
              >
                <Trash></Trash>
              </div>
              <div
                @click="GivePoint(event.ID)"
                class="flex w-12 btn btn-success"
              >
                <IconCorrect></IconCorrect>
              </div>
            </div>
          </td>
        </tr>
      </Table>
    </div>
  </Adminlayout>
</template>
