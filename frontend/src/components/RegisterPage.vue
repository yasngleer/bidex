<template>
    <div class="d-flex align-center justify-center" style="height: 90vh">
        <v-sheet width="400" class="mx-auto">
            <v-form fast-fail @submit.prevent="login">
                <v-text-field variant="outlined" v-model="username" label="User Name"></v-text-field>
                <v-text-field variant="outlined" v-model="password" type="password" label="password"></v-text-field>
                <v-btn type="submit" color="primary" block class="mt-2">Sign in</v-btn>
            </v-form>
            <div class="mt-2">
                <p class="text-body-2">You already have an account? <RouterLink to="/login">Login</RouterLink></p>
            </div>
        </v-sheet>
    </div>
</template>
<script>
import axios from "axios"    

import router from "../router"
export default {
    data() {
        return {
            "username": "",
            "password": "",
        };
    },
    methods: {
        login() {
            let data = {    
                        email: this.username,    
                        password: this.password    
                    }    
                    axios.post("http://localhost:8000/api/users", data,{ withCredentials: true })    
                        .then((response) => {
                            console.log(response)
                            router.push("/")    
                        })    
                        .catch((err) => {    
                            console.log("Failed to log in")    
                            console.log(err)
                        }) 
            // Your login logic here
        },
    },
    mounted(){
        axios.get("http://localhost:8000/api/user/me",{ withCredentials: true })    
                        .then(() => {
                            router.push("/")
                        })    
    }
    
}
</script>