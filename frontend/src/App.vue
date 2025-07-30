<script setup type="module">
import { ref, onMounted } from 'vue';

const message = ref('Загрузка...');

onMounted(async () => {
  try {
    const response = await fetch('http://localhost/api/hello');

    if (!response.ok) {
      throw new Error(`Ошибка HTTP: ${response.status}`);
    }

    const data = await response.json();

    message.value = data.message;
    console.log(message.value);

  } catch (error) {
    console.error('Ошибка при получении данных:', error);
    message.value = 'Не удалось загрузить сообщение.';
  }
});
</script>

<template>
  <div>
    {{ message }}
  </div>
</template>
