<template>
    <div class="bg-white grid grid-cols-4 shadow-2xl">
        <!-- Image -->
        <div class="col-span-2 h-portImg md:block hidden overflow-hidden bg-darkBlue">
            <img class="" :src="image" alt="Large portfolio image" >
        </div>
        <!-- Right Side -->
        <div class="col-span-full md:col-span-2 bg-cover md:bg-none flex items-center">
            <div class="px-16">
                <h3 class="font-OpenSans font-bold text-orange text-4xl block mt-8 md:mt-16">{{title}}</h3>
                <p class="font-Nunito font-bold text-blue text-xl block mt-4">{{type}}</p>
                <a class="font-Nunito font-semibold text-darkBlue underline text-md my-8 md:mb-0 md:mt-16 block" :href="link">Read More</a>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
import Component from 'vue-class-component'

const LargePortProps = Vue.extend({

    props: {
        image: String,
        title: String,
        type: String,
        link: String,
    }

})

@Component
export default class LargePortfolio extends LargePortProps {

}
</script>

<style>

</style>
