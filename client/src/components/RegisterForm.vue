<script lang="ts">
export default {
    data: () => {
        return {
            username: '',
            password: '',
            confirm: '',
        }
    },
    methods: {
        async register(e: any) {
            e.preventDefault()

            const res = await fetch("http://localhost:8080/api/register", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    username: this.username,
                    password: this.password,
                    confirm: this.confirm,
                }),
            })


            if (res.status === 201) {

            } else {

            }
        }
    }
}
</script>

<template>
    <form @submit.prevent="register" class="auth-form">
        <p>Register</p>

        <div class="field">
            <input v-model="username" placeholder="Username" />
        </div>
        <div class="field">
            <input v-model="password" placeholder="Password" type="password" />
        </div>
        <div class="field">
            <input v-model="confirm" placeholder="Confirm password" type="password" />
        </div>
        <button type="submit">Register</button>

        <div class="field">
            <RouterLink to="/login">Sign in</RouterLink>
        </div>
    </form>
</template>