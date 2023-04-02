<template>
       <v-col  sm=6 md=4 lg=3>
          <v-card :to='"/item/"+item.id'
    class="mx-auto"
  >
    <v-img
    v-bind:src="item.image_url"
    aspect-ratio="16/9"
    cover
    ></v-img>

    <v-card-title>
     {{item.name}}
    </v-card-title>

    <v-card-subtitle>
{{item.description}}    </v-card-subtitle>

    <v-card-actions>
      <v-btn
        color="orange-lighten-2"
        variant="text"
      >
        Explore
      </v-btn>

      <v-spacer></v-spacer>

      <v-btn
      ></v-btn>
    </v-card-actions>

  </v-card>         </v-col>
</template>

<script>

export default {
    name:"ItemCard",
    props:["item"],
    setup () {
        

        return {}
    }
}
</script>

<style lang="scss" scoped>

</style>