 <script setup>
import { defineProps } from 'vue';

defineProps({
    headers: Array
})

</script>
 <template>
    <table class="table table-zebra shadow">
        <!-- head -->
        <thead>
            <tr>
                <th v-for="header in headers" class="text-2xl"> {{ header }}</th>
            </tr>
        </thead>
        <tbody>
            <slot></slot>
        </tbody>
    </table>
</template>