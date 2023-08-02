<script lang="ts">


export default {
    data: () => {
        return {
            username: '',
            password: '',
        }
    },
    methods: {
        async login(e: any) {
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
    </form>
</template>

<style>
.auth-form {
    width: 50%;
    margin: auto;
    margin-top: 5%;
    text-align: center;
    border: 1px solid #000;
    border-radius: 10px;
    padding: 1em;
}

.auth-form .field {
    margin-top: 0.2em;
    padding: 1em;
}

.auth-form input {
    width: 75%;
    padding: 0.5em;
    margin: 8px 0;
    font-size: 1.2em;
    border-radius: 10px;
    outline: none !important;
    border: 0.1em solid #000;
}

.auth-form input:focus {
    border: 0.1em solid #05be4c;
}

.auth-form p {
    font-size: 2em;
}

.auth-form button {
    font-size: 1.5em;
}

.auth-form a {
    font-size: 1.5em;
}
</style>