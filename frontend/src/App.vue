<template>
  <v-app>

    <v-main>

      <v-app-bar v-if="$router.currentRoute.value.path !== '/login' && $router.currentRoute.value.path !== '/register'" :elevation="3">
        <v-btn @click="logout">Logout</v-btn>
      </v-app-bar>

      <router-view />
    </v-main>

  </v-app>
</template>
<script>
import axios from "axios"
import router from "./router"
export default {

  name: 'App',
  mounted() { this.checkauth(); },
  components: {
  },
  methods: {
    checkauth: function () {

      axios.get("http://localhost:8000/api/user/me", { withCredentials: true })
        .catch(() => {
          router.push("/login")
        })

    },
    logout: function () {
      axios.get("http://localhost:8000/api/user/logout", { withCredentials: true })
        .then(() => {
          router.push("/login")
        })
        .catch((err) => {
          console.log("Failed to logout")
          router.push("/login")
          console.log(err)
        })

    },
  },
  data: () => ({
    //
  }),
}
</script>
