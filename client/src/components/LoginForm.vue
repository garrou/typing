<script lang="ts">


export default {
    data: () => {
        return {
            username: '',
            password: '',
            error: ''
        }
    },
    methods: {
        async login(e: any): Promise<void> {
            e.preventDefault()

            const res = await fetch("http://localhost:8080/api/login", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    username: this.username,
                    password: this.password,
                }),
            })

            if (res.status === 200) {
                const { data } = await res.json()
                localStorage.setItem('jwt', data)
                this.$router.push('/dashboard')
            } else {
                const { message } = await res.json()
                this.error = message
            }
        }
    }
}
</script>

<template>
    <form @submit.prevent="login" class="auth-form">
        <p>Login</p>

        <div class="field">
            <input v-model="username" placeholder="Username" />
        </div>
        <div class="field">
            <input v-model="password" placeholder="Password" type="password" />
        </div>
        <button type="submit">Login</button>

        <div class="field">
            <RouterLink to="/register">Sign up</RouterLink>
        </div>

        <p class="error" v-if="error">{{ error }}</p>
    </form>
</template>