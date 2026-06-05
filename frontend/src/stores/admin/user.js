import { defineStore } from "pinia"
import axios from "axios"

export const useAdminUserStore = defineStore("admin-user", {
  state: () => ({
    userId: localStorage.getItem("userId") || null,
    username: localStorage.getItem("username") || null,
    currentUser: null,
  }),
  actions: {
    async addEventUser(eventData) {
      const response = await axios.post(
        "http://localhost:8000/event-user",
        eventData
      )
    },
    async updateUserPoint(userId, newPoint) {
      try {
        await axios.put(`http://localhost:8000/user/${userId}`, {
          point: newPoint,
        })
      } catch (error) {
        console.log("error", error)
      }
    },
    async loadUsers() {
      try {
        const response = await axios.get("http://localhost:8000/users")
        this.list = response.data
        return this.list
      } catch (error) {
        console.log("error", error)
      }
    },
    async loadUser(id) {
      try {
        console.log("Fetching user with ID:", id)
        const response = await axios.get(`http://localhost:8000/user/${id}`)
        console.log("loadUserID", response.data)
        this.currentUser = response.data
        this.seletectUser = response.data
        return response.data
      } catch (error) {
        console.log("error", error)
      }
    },
    async updateUserID(id, userData) {
      try {
        const response = await axios.put(
          `http://localhost:8000/user/${id}`,
          userData
        )
        console.log("updateUser", response.data)
      } catch (error) {
        console.log("error", error)
      }
    },
    async registerUser(userData) {
      try {
        const bodyUser = {
          firstname: userData.firstname,
          lastname: userData.lastname,
          email: userData.email,
          password: userData.password,
          phone: userData.phone,
          gender: userData.gender,
          birthday: userData.birthday,
          occupation: userData.occupation,
          salary: userData.salary,
          role: "User",
          status: "Active",
          point: 0,
        }
        const response = await axios.post(
          "http://localhost:8000/register",
          bodyUser
        )
        console.log("User registered successfully:", response)
      } catch (error) {
        console.log("error", error)
      }
    },
    async loginUser(userData) {
      try {
        const bodyUser = {
          email: userData.email,
          password: userData.password,
        }
        const response = await axios.post(
          "http://localhost:8000/login",
          bodyUser
        )
        this.userId = response.data.user_id
        this.username = response.data.username
        localStorage.setItem("userId", this.userId)
        localStorage.setItem("username", this.username)
        return response.data
      } catch (error) {
        console.log("Error logging in", error.response?.data || error.message)
        throw error
      }
    },
    logout() {
      this.userId = null
      this.username = null
      localStorage.removeItem("userId")
      localStorage.removeItem("username")
    },
    getUser(index) {
      return this.list[index]
    },
    updateUser(index, userData) {
      // this.list[index].name = userData.name
      // this.list[index].role = userData.role
      // this.list[index].status = userData.state
      const fields = ["name", "role", "status"]
      for (let field of fields) {
        this.list[index][field] = userData[field]
      }
      this.list[index].update = new Date().toISOString()
    },
    deleteUser(index) {
      this.list.splice(index, 1)
    },
  },
})
