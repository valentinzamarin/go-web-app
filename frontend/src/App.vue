<script setup type="module">
import { ref, onMounted } from 'vue';
import Form from './components/Form.vue'


const message = ref();

onMounted(async () => {

});
</script>

<template>
  <div class="flex flex-col h-screen">
    <Form />
  </div>
</template>
