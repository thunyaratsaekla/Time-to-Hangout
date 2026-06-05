import { defineStore } from "pinia"
import axios from "axios"
import { useAdminUserStore } from "@/stores/admin/user"

export const useAdminProductStore = defineStore("admin-product", {
  state: () => ({
    list: [],
    loaded: false,
    selectedEvent: {},
    status: ["Open", "lose"],
    joinedEvent: JSON.parse(localStorage.getItem("joinedEvent") || "[]"),
    eventsUser: [],
    selectedEventUser: {},
  }),
  actions: {
   async loadEventsUser() {
  const userStore = useAdminUserStore()
  const userId = userStore.userId || localStorage.getItem("userId")
  try {
    const response = await axios.get(`http://localhost:8000/events-user/${userId}`)
    this.eventsUser = response.data
  } catch (error) {
    console.log("error", error)
  }
},
    async loadEventUser(id) {
      try {
        console.log("Fetching event with ID:", id)
        const response = await axios.get(
          `http://localhost:8000/event-user/${id}`
        )
        console.log("Event data:", response.data)
        this.selectedEventUser = response.data
        return response.data
      } catch (error) {
        console.log("error", error)
      }
    },
    async addEventUser(eventData) {
      try {
        const bodyEvent = {
          name: eventData.name,
          person: eventData.person,
          location: eventData.location,
          about: eventData.about,
          point: "1",
          user_id: eventData.user_id,
          status: "Unfinished",
          timestart: eventData.timestart,
          timeend: eventData.timeend,
          image: eventData.image,
        }
        const response = await axios.post(
          "http://localhost:8000/event-user",
          bodyEvent
        )
        console.log("addEvent complete", response)
      } catch (error) {
        console.log("error", error)
      }
    },
    async updateEventUser(id, eventData) {
      const eventId = eventData.ID
      try {
        const response = await axios.put(
          `http://localhost:8000/event-user/${id}`,
          eventData
        )
        console.log("editEvent complete", response)
      } catch (error) {
        console.log("error", error)
      }
    },
    async deleteEventUser(id) {
      try {
        const response = await axios.delete(
          `http://localhost:8000/event-user/${id}`
        )
        console.log("deleteEvent", response)
      } catch (error) {
        console.log("error", error)
      }
    },
    async addJoinedEvent(event) {
      const userStore = useAdminUserStore()
      const userId = userStore.userId || localStorage.getItem("userId")
      if (!userId) return
      try {
        console.log("userId:", userId, "eventId:", event.ID)
        await axios.post("http://localhost:8000/join-event", {
          user_id: Number(userId),
          event_id: Number(event.ID),
        })
        this.joinedEvent.push(event)
      } catch (error) {
        console.log("error", error)
      }
    },
    async loadJoinedEvent() {
      const userStore = useAdminUserStore()
      const userId = userStore.userId || localStorage.getItem("userId")
      if (!userId) return
      try {
        const response = await axios.get(
          `http://localhost:8000/joined-events/${userId}`
        )
        this.joinedEvent = response.data
      } catch (error) {
        console.log("error", error)
      }
    },
    async loadEvents() {
      try {
        const response = await axios.get("http://localhost:8000/events")
        this.list = response.data
      } catch (error) {
        console.log("error", error)
      }
    },
    async loadEvent(id) {
      try {
        console.log("Fetching event with ID:", id)
        const response = await axios.get(`http://localhost:8000/event/${id}`)
        console.log("Event data:", response.data)
        this.selectedEvent = response.data
        return response.data
      } catch (error) {
        console.log("error", error)
      }
    },
    async addEvent(eventData) {
      try {
        const bodyEvent = {
          name: eventData.name,
          organizer: eventData.organizer,
          location: eventData.location,
          about: eventData.about,
          point: eventData.point,
          status: eventData.status,
          timestart: eventData.timestart,
          timeend: eventData.timeend,
          image: eventData.image,
        }
        const response = await axios.post(
          "http://localhost:8000/event",
          bodyEvent
        )
        console.log("addEvent complete", response)
      } catch (error) {
        console.log("error", error)
      }
    },
    async updateEvent(id, eventData) {
      const eventId = eventData.ID
      try {
        const response = await axios.put(
          `http://localhost:8000/event/${id}`,
          eventData
        )
        console.log("editEvent complete", response)
      } catch (error) {
        console.log("error", error)
      }
    },
    async deleteEvent(id) {
      try {
        const response = await axios.delete(`http://localhost:8000/event/${id}`)
        console.log("deleteEvent", response)
      } catch (error) {
        console.log("error", error)
      }
    },
    async givePoint(eventID) {
      try {
        const response = await axios.post(
          `http://localhost:8000/give-point/${eventID}`
        )
        console.log(response.data)
        return response.data
      } catch (error) {
        console.log("error", error)
      }
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
    loadProducts() {
      const productsData = localStorage.getItem("admin-products")
      if (productsData) {
        this.list = JSON.parse(productsData)
        this.loaded = true
      }
    },
    getProduct(index) {
      if (!this.loaded) {
        this.loadProducts()
      }
      return this.list[index]
    },
    addProduct(productData) {
      productData.remainquantity = productData.quantity
      productData.update = new Date().toISOString()
      this.list.push(productData)
      localStorage.setItem("admin-products", JSON.stringify(this.list))
    },
    updateProduct(index, product) {
      this.list[index].name = product.name
      this.list[index].image = product.image
      this.list[index].price = product.price
      this.list[index].organizer = product.organizer
      this.list[index].location = product.location
      this.list[index].timestart = product.timestart
      this.list[index].timeend = product.timeend
      this.list[index].quantity = product.quantity
      this.list[index].remainquantity = product.quantity
      this.list[index].status = product.status
      this.list[index].event = product.event
      this.list[index].update = new Date().toISOString()
      localStorage.setItem("admin-products", JSON.stringify(this.list))
    },
    removeProduct(index) {
      this.list.splice(index, 1)
      localStorage.setItem("admin-products", JSON.stringify(this.list))
    },
    async updateEventStatusToFinished(id) {
      try {
        const response = await axios.put(`http://localhost:8000/event-user/finish/${id}`)
        console.log("Event status updated to Finished:", response.data)
        return response.data
      } catch (error) {
        console.log("Error updating event status:", error)
      }
    }
  },
})
