<script setup>
import { onMounted } from 'vue';
import FormInput from './form/FormInput.vue'



const createUser = async () => {

    let $this = event.target,
        username = $this.username,
        password = $this.password;
    if (!username.value || !password.value) {
        alert("Empty fields")
        return
    }

    const userData = {
        username: username.value,
        password: password.value,
        role: role.value
    }

    try {
        const response = await fetch('http://localhost:8080/api/create', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(userData)
        })

        if (response.ok) {
            alert('Пользователь создан')
            username.value = ''
            password.value = ''
            fetchUsers()
        } else {
            const errorData = await response.json()
            alert(`Ошибка при создании пользователя: ${errorData.message || response.statusText}`)
        }
    } catch (error) {
        console.error('Ошибка сети или сервера:', error)
        alert('Произошла ошибка при отправке запроса. Пожалуйста, попробуйте позже.')
    }
}

const fetchUsers = async () => {

    const url = "http://localhost:8080/api/users";
    try {
        const response = await fetch(url);
        if (!response.ok) {
            throw new Error(`Response status: ${response.status}`);
        }

        const json = await response.json();
        console.log(json);
    } catch (error) {
        console.error(error.message);
    }

};


onMounted(() => {
    fetchUsers();
});

</script>

<template>
    <form @submit.prevent="createUser" class="max-w-[480px] w-full mx-auto">
        <FormInput label="Username" id="username" name="username" placeholder="username" />
        <FormInput label="Password" id="password" name="password" placeholder="password" />
        <FormInput type="hidden" id="role" name="role" model-value="role" />
        <button type="submit"
            class="rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-xs hover:bg-indigo-500 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">
            Create
        </button>
    </form>
</template>