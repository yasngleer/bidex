<template>
  <v-container fluid>
    <v-row class="text-center">
      <ItemCard v-for="item in items" :key="item.id" :item="item" />

    </v-row>

  </v-container>
  <VLayoutItem model-value position="bottom" class="text-end" size="88">
    <div class="ma-4">

      <VBtn @click="dialog = true" icon="mdi-plus" size="large" color="primary" elevation="8" />
    </div>
  </VLayoutItem>



  <!-- add dialog -->
  <v-dialog v-model="dialog" width="1024">

    <v-card>
      <v-card-title>
        <span class="text-h5">Add Item</span>
      </v-card-title>
      <v-card-text>
        <v-container>
          <v-row>

            <v-col cols="12">
              <v-text-field v-model="addname" label="Item Name" required></v-text-field>
            </v-col>
            <v-col cols="12">
              <v-text-field v-model="adddescription" label="Item Description" required></v-text-field>
            </v-col>
            <v-col cols="12">
              <v-text-field v-model="addimageurl" label="Image URL" required></v-text-field>
            </v-col>
          </v-row>
        </v-container>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="blue-darken-1" variant="text" @click="dialog = false">
          Close
        </v-btn>
        <v-btn color="blue-darken-1" variant="text" @click="additems()">
          Save
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
import axios from "axios"
import ItemCard from "./Card.vue"
import router from "@/router"


export default {
  name: 'HelloWorld',
  components: { ItemCard },
  data: () => {
    return {
      addname: "",
      adddescription: "",
      addimageurl:"https://cdn.vuetifyjs.com/images/cards/sunshine.jpg",
      items: [],
      dialog: false,
    }
  },
  methods: {
    fetchitems: function () {
      axios.get("http://localhost:8000/api/items", { withCredentials: true })
        .then((response) => {
          console.log(response)
          this.items = response.data
        })
        .catch((err) => {
          console.log("Failed to log in")
          console.log(err)
          router.push("/login")
        })

    },
    additems: function () {
      let data = {
        "name": this.addname,
        "description": this.adddescription,
        "image_url":this.addimageurl,
      }
      axios.post("http://localhost:8000/api/items", data, { withCredentials: true })
        .then((response) => {
          console.log(response)
          this.dialog = false
          this.fetchitems()

        })
        .catch((err) => {
          console.log("Failed to add")
          console.log(err)
        })

    },
  },
  mounted() {
    this.fetchitems()
  }

}
</script>
