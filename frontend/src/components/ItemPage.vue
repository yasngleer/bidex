<template>
    <v-row v-if="item" justify="center pa-10" fluid>
        <!-- item -->
        <v-col cols="12" sm="10" md="8" class="text-center">
            <v-card class="" :elevation="3">
                <v-row justify="center" no-gutters>
                    <v-col  cols="12" sm="10" md="6" >
                        <v-img   height="100%"  cover v-bind:src="item.image_url"></v-img>

                    </v-col>
                    <v-col  cols="12" sm="10" md="6">
                        <div style="display: table"
                            class="ma-5  pa-5 bg-indigo-lighten-1 rounded-pill text-h4 rounded-pill text-left">
                            {{ item.name }}
                        </div>
                        <div style="display: table"
                            class="ma-5  pa-5 bg-indigo-lighten-5 rounded-pill text-bosy-1 rounded-pill text-left">
                            {{ item.description }}
                        </div>
                    </v-col>
                </v-row>
            </v-card>
        </v-col>
        <!-- bids -->
        <v-col v-for="bid in item.bids" :key="bid.id" xs="12" sm="10" md="8" class="text-center">
            <v-card class="" :elevation="3">
                <v-row no-gutters>
                    <v-col>
                        <div style="display: table"
                            class="ma-5  pa-5 bg-indigo-lighten-1 rounded-pill text-h6 rounded-pill text-left">{{
                                bid.user.email }}
                        </div>

                    </v-col>
                    <v-row class="justify-end ma-5">
                        <div style="display: table"
                            class="pa-5 bg-indigo-lighten-1 rounded-pill text-h6 rounded-pill text-right">
                            {{ bid.price }} tl
                        </div>
                    </v-row>
                </v-row>
            </v-card>
        </v-col>
        <!-- no-bids -->

        <v-col v-if="item.bids.length==0" xs="12" sm="10" md="8" class="text-center">
            <v-card class="" :elevation="3">
                <v-col>
                    <div style="display: table" class="ma-5  pa-5 bg-red rounded-pill text-h6 rounded-pill text-left">There
                        is no bids</div>

                </v-col>
            </v-card>
        </v-col>

        <!-- fab -->
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
                                <v-text-field :rules="[
                                    v => !!v || 'Field is required'
                                ]" v-model="addprice" type="number" label="Bid Price" required></v-text-field>
                            </v-col>
                        </v-row>
                    </v-container>
                </v-card-text>
                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="blue-darken-1" variant="text" @click="dialog = false">
                        Close
                    </v-btn>
                    <v-btn color="blue-darken-1" variant="text" @click="addbid()">
                        Save
                    </v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>

    </v-row>
</template>

<script>
import axios from 'axios';

export default {
    
    data: () => {
        return {
            addprice: null,
            dialog: false,
            item: null,
        }
    },
    methods: {
        fetchitem: function () {
            let itemid = this.$route.params.id
            axios.get("http://localhost:8000/api/items/" + itemid, { withCredentials: true })
                .then((response) => {
                    console.log(response)
                    this.item = response.data
                })
                .catch((err) => { console.log(err) })
        },
        addbid: function () {
            let itemid = this.$route.params.id
            let data = {
                "price": parseFloat(this.addprice),
            }
            axios.post("http://localhost:8000/api/items/" + itemid + "/bid", data, { withCredentials: true })
                .then((response) => {
                    console.log(response)
                    this.dialog = false
                })
                .catch((err) => {
                    console.log("Failed to add")
                    console.log(err)
                })

        },
        startws: function () {
            let itemid = this.$route.params.id
            console.log("Starting connection to WebSocket Server")
            this.connection = new WebSocket("ws://localhost:8000/api/ws/" + itemid)
            console.log("item: "+this.item);

            this.connection.onmessage = (event) => {
                console.log("item: "+this.item);
                console.log("item new: "+event.data);

                this.item.bids.push(JSON.parse(event.data))
            }
        },
    },

    mounted() {
        this.fetchitem()
        this.startws()


    }
}
</script>

<style lang="scss" scoped></style>