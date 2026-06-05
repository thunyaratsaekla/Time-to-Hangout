<script setup>
import Adminlayout from "@/layouts/Adminlayout.vue"
import { RouterLink, useRouter } from "vue-router"
import { onMounted } from "vue"

import { useAdminUserStore } from "@/stores/admin/user"

import Edit from "@/components/icons/Edit.vue"
import Trash from "@/components/icons/Trash.vue"
import Table from "@/components/Table.vue"

const router = useRouter()
const adminUserStore = useAdminUserStore()

onMounted(async () => {
  try {
    const UserData = await adminUserStore.loadUsers()
    console.log("Loaded users:", UserData)
  } catch (error) {
    console.error("Error loading users:", error)
  }
})

const ChangeStatus = async (UserID) => {
  console.log("ChangeStatusID", UserID)
  const selectedUser = await adminUserStore.loadUser(UserID)
  console.log("selectedUserChangeStatus", selectedUser)
  selectedUser.status = selectedUser.status === "Active" ? "Inactive" : "Active"
  await adminUserStore.updateUserID(UserID, selectedUser)
  await adminUserStore.loadUsers()
}

const EditUser = async (UserID) => {
  router.push({ name: "admin-users-update", params: { id: UserID } })
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
</script>

<template>
  <Adminlayout>
    <div class="flex justify-between items-center">
      <div class="text-4xl font-bold">User</div>
    </div>
    <div class="divider"></div>
    <Table :headers="['Name', 'Role', 'Status', 'UpdateAt', '']">
      <tr v-for="(user, index) in adminUserStore.list" :key="user.ID">
        <th>{{ user.firstname }}</th>
        <td>{{ user.role }}</td>
        <td
          class="badge"
          :class="user.status === 'Active' ? 'badge-success' : 'badge-error'"
        >
          {{ user.status }}
        </td>
        <td>{{ formatTimeWithoutTimezone(user.UpdatedAt) }}</td>
        <td class="flex gap-5 justify-center">
          <button @click="EditUser(user.ID)" class="btn btn-info">Edit</button>
          <button @click="ChangeStatus(user.ID)" class="btn btn-info">
            {{ user.status === "active" ? "Disable" : "Enable" }}
          </button>
        </td>
      </tr>
    </Table>
  </Adminlayout>
</template>
