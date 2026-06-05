import { createRouter, createWebHistory } from "vue-router"

// User
import LoginandRegisterUserView from "@/views/user/LoginandRegisterView.vue"
import LoginUserView from "@/views/user/LoginView.vue"
import RegisterUserView from "@/views/user/RegisterView.vue"
import HomeView from "@/views/user/HomeView.vue"
import UserOrganizerView from "@/views/user/Event/UserOrganizerView.vue"
import MyHangoutView from "@/views/user/Event/MyHangoutView.vue"
import AddEventView from "@/views/user/Event/AddEventView.vue"
import UpdateEventView from "@/views/user/Event/UpdateEventView.vue"
import MyPointView from "@/views/user/Event/MyPointView.vue"

// Admin 
import HomeAdminView from "@/views/admin/HomeAdmin.vue"
import LoginandRegisterAdminView from "@/views/admin/LoginandRegisterView.vue"
import LoginAdminView from "@/views/admin/LoginView.vue"
import RegisterAdminView from "@/views/admin/RegisterView.vue"
import AdminEventListView from "@/views/admin/event/ListView.vue"
import AdminEventUpdateView from "@/views/admin/event/UpdateView.vue"

import AdminUserListView from "@/views/admin/user/ListView.vue"
import AdminUserUpdateView from "@/views/admin/user/UpdateView.vue"


const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      redirect: "/lr",
    },
    // User
     {
      path: "/lr",
      name: "user-lr",
      component: LoginandRegisterUserView,
    },
    {
      path: "/login", 
      name: "user-login",
      component: LoginUserView,
    },
    {
      path: "/register",
      name: "user-register",
      component: RegisterUserView,
    },
    {
      path: "/home",
      name: "user-home",
      component: HomeView,
    },
    {
      path: "/eventorganizer",
      name: "user-organizer",
      component: UserOrganizerView,
    },
    {
      path: "/myhangout",
      name: "user-myhangout",
      component: MyHangoutView,
    },
        {
      path: "/myhangout/create",
      name: "user-myhangout-create",
      component: AddEventView,
    },
    {
      path: "/myhangout/update/:id",
      name: "user-myhangout-update",
      component: UpdateEventView,
    },
    {
      path: "/mypoint",
      name: "user-mypoint",
      component: MyPointView,
    },
    //Admin
    {
      path: "/admin/lr",
      name: "admin-lr",
      component: LoginandRegisterAdminView,
    },
    {
      path: "/admin/login",
      name: "admin-login",
      component: LoginAdminView,
    },
    {
      path: "/admin/register",
      name: "admin-register",
      component: RegisterAdminView,
    },
        {
      path: "/admin/home",
      name: "admin-home",
      component: HomeAdminView,
    },
    {
      path: "/admin/products",
      name: "admin-product-list",
      component: AdminEventListView,
    },
    {
      path: "/admin/products/create",
      name: "admin-product-create",
      component: AdminEventUpdateView,
    },
    {
      path: "/admin/products/update/:id",
      name: "admin-product-update",
      component: AdminEventUpdateView,
    },
    {
      path: "/admin/users",
      name: "admin-users-list",
      component: AdminUserListView,
    },
    {
      path: "/admin/users/update/:id",
      name: "admin-users-update",
      component: AdminUserUpdateView,
    },
    
  ],
})

export default router
