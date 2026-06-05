<script setup>
import Adminlayout from '@/layouts/Adminlayout.vue';
import { useAdminUserStore } from '@/stores/admin/user';
import { reactive, ref , onMounted } from 'vue';
import { useRoute , useRouter } from 'vue-router';


const adminUserStore = useAdminUserStore()

const route = useRoute()
const router = useRouter()

const userIndex = ref(-1)
const errorMessage = ref('')

const inputForm = [
    {
        name: 'First Name',
        field: "firstname",
        type: 'text'
    },
    {
        name: 'Role',
        field: 'role',
        type: 'select',
        dropdownlist: ['Admin', 'Moderator' , 'User']
    },
    {
        name: 'Status',
        field: 'status',
        type: 'select',
        dropdownlist: ['Active', 'Inactive']
    }
]

const userData = reactive({
    firstname: '',
    role: '',
    status: ''
})

onMounted(async() => {
    if (route.params.id) {
        userIndex.value = parseInt(route.params.id)
        const selectedUser = await adminUserStore.loadUser(userIndex.value)
        console.log("selectedUser" , selectedUser)
        userData.firstname = selectedUser.firstname
        userData.role = selectedUser.role
        userData.status = selectedUser.status
    }
    
})

const UpdateUser = async() => {
    errorMessage.value = ''
    if (!userData.firstname || !userData.role || !userData.status) {
        errorMessage.value = 'Please fill in all required fields.';
        return
    }
    try {
        await adminUserStore.updateUserID(userIndex.value , userData)
        router.push({ name: 'admin-users-list'})
    }catch {
        console.error('Error submitting update user:', error);
        errorMessage.value = 'An error occurred while submitting the update user.';
    }
    
}
</script>

<template>
    <Adminlayout>
        <div class="shadow-xl p-10">
            <div class="text-3xl font-bold"> Update</div>
            <div class="divider"></div>
            <div v-for="form in inputForm" class="py-4">
                <fieldset class="fieldset grid grid-cols-1 gap-4">
                    <div>
                        <legend class="fieldset-legend"> {{ form.name }} :</legend>
                        <input v-if="form.type === 'text'" v-model="userData[form.field]" type="text" class="input input-info w-full font-bold"
                            placeholder="Type here" />
                    </div>
                </fieldset>
                <fieldset class="fieldset">
                    <select v-if="form.type === 'select'" v-model="userData[form.field]" class="select input-info w-full font-bold">
                        <option v-for="item in form.dropdownlist">{{ item }}</option>
                    </select>
                </fieldset>
            </div>
            <div class="flex justify-end gap-4 items-center">
            <div v-if="errorMessage" class="alert alert-error mt-4">
                {{ errorMessage }}
            </div>
                <div>
                    <RouterLink :to="{name:'admin-users-list'}" class="btn btn-ghost">Back</RouterLink>
                </div>
                <div>
                    <button @click="UpdateUser()" class="btn btn-success">Update</button>
            </div>
            </div>
        </div>
    </Adminlayout>
</template>