<script setup>
import Userslayout from "@/layouts/Userslayout.vue"
import { ref } from "vue"
import { onMounted } from "vue"
import { useRouter, RouterLink } from "vue-router"
import { useAdminProductStore } from "@/stores/admin/product"
import { useAdminUserStore } from "@/stores/admin/user"

import Edit from "@/components/icons/Edit.vue"
import Trash from "@/components/icons/Trash.vue"
import Table from "@/components/Table.vue"
import IconCorrect from "@/components/icons/IconCorrect.vue"
import IconIncorrect from "@/components/icons/IconIncorrect.vue"

const adminProductStore = useAdminProductStore()
const adminUserStore = useAdminUserStore()
const router = useRouter()

onMounted(async () => {
  // adminProductStore.loadProducts()
  await adminProductStore.loadJoinedEvent()
  await adminProductStore.loadEventsUser()
})

const handleDelete = async (eventID) => {
  console.log(eventID)
    await adminProductStore.deleteEventUser(eventID)
    await adminProductStore.loadEventsUser()
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

const handleEdit = (event) => {
  router.push({ name: "user-myhangout-update", params: { id: event.ID } })
}

const handleFinish = async (event) => {
  try {
    // 1. อัปเดตสถานะ Event เป็น Finished
    const result = await adminProductStore.updateEventStatusToFinished(event.ID)
    console.log("Finish result:", result)

    // 2. ดึง point ของ event (string → int)
    const eventPoint = parseInt(event.point || event.Point || "0", 10)

    // 3. ดึง user ปัจจุบัน
    const userId = adminUserStore.userId || localStorage.getItem("userId")
    await adminUserStore.loadUser(userId)
    const currentPoint = adminUserStore.currentUser?.point || 0

    // 4. บวกแต้ม
    const newPoint = currentPoint + eventPoint

    // 5. อัปเดต point ใน backend
    await adminUserStore.updateUserPoint(userId, newPoint)

    await adminUserStore.loadUser(userId)
    

    // 6. โหลดข้อมูลใหม่
    await adminProductStore.deleteEventUser(event.ID)
    await adminProductStore.loadEventsUser()
  } catch (error) {
    console.error("Error finishing event:", error)
  }
}
</script>

<template>
  <Userslayout>
    <div class="shadow-md rounded-box p-4">
      <div class="flex justify-between items-center">
        <div class="text-4xl font-bold">My Hangout</div>
        <div>
          <RouterLink
            :to="{ name: 'user-myhangout-create' }"
            class="btn btn-lg btn-primary"
            >Add Event</RouterLink
          >
        </div>
      </div>
      <div class="divider"></div>
      <div class="mb-8">
        <div class="text-2xl font-bold mb-2">My Events</div>
        <Table
          :headers="[
            'Name',
            'Image',
            'Person',
            'Location',
            'Status',
            'DateAndTime',
            'Actions',
          ]"
        >
          <tr v-for="event in adminProductStore.eventsUser">
            <th>{{ event.name }}</th>
            <td>
              <div class="mask mask-squircle w-24 h-24">
                <img :src="event.image" class="w-full h-full object-cover" />
              </div>
            </td>
            <td>{{ event.person }}</td>
            <td>{{ event.location }}</td>
            <td>
              <div
                class="badge"
                :class="
                  event.status === 'Finished' ? 'badge-success' : 'badge-error'
                "
              >
                {{ event.status }}
              </div>
            </td>
            <td>{{ formatTimeWithoutTimezone(event.timestart) }}</td>
            <td>
              <button
                class="btn btn-xs btn-warning mr-2"
                @click="handleEdit(event)"
              >
                <Edit class="w-4 h-4" /> Edit
              </button>
              <button class="btn btn-xs btn-error mr-2" @click="handleDelete(event.ID)">
                <Trash class="w-4 h-4" /> Delete
              </button>
              <button class="btn btn-xs btn-success" @click="handleFinish(event)">
                <IconCorrect class="w-4 h-4" /> Finish
              </button>
            </td>
            
          </tr>
        </Table>
      </div>
      <div>
        <div class="text-2xl font-bold mb-2">Organizer Events</div>
        <Table
          :headers="[
            'Name',
            'Image',
            'Organizer',
            'Location',
            'Status',
            'DateAndTime',
          ]"
        >
          <tr
            v-for="event in adminProductStore.joinedEvent"
            :key="event.ID || event.id"
          >
            <th>{{ event.name }}</th>
            <td>
              <div class="mask mask-squircle w-24 h-24">
                <img :src="event.image" class="w-full h-full object-cover" />
              </div>
            </td>
            <td>{{ event.organizer || event.person }}</td>
            <td>{{ event.location }}</td>
            <td>
              <div
                class="badge"
                :class="
                  event.status === 'Open' ? 'badge-success' : 'badge-error'
                "
              >
                {{ event.status }}
              </div>
            </td>
            <td>{{ formatTimeWithoutTimezone(event.timestart) }}</td>
          </tr>
        </Table>
      </div>
    </div>
  </Userslayout>
</template>
