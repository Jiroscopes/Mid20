<template>
    <nav class="bg-white flex-wrap justify-between items-center px-6 py-3 w-full fixed z-30">
        <div class="block lg:hidden">
            <button @click="toggleNav" class="flex items-center px-3 py-2 border rounded text-blue border-darkBlue hover:text-darkBlue hover:border-darkBlue">
                <svg class="fill-current h-3 w-3" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><title>Menu</title><path d="M0 3h20v2H0V3zm0 6h20v2H0V9zm0 6h20v2H0v-2z"/></svg>
            </button>
        </div>
        <div class="w-full block flex-grow lg:flex lg:items-center lg:w-auto">
            <div class="text-base lg:flex-grow">
                <a href="#home" v-bind:class="{hidden: isHidden}" class="transition duration-500 font-Nunito font-bold text-xl block mt-4 lg:inline-block lg:mt-0 text-orange hover:text-darkBlue mr-4">
                    Mid-Twenty
                </a>
                <a href="#home" v-bind:class="{hidden: isHidden}" class="font-Nunito font-bold text block mt-4 lg:inline-block lg:mt-0 text-blue hover:text-darkBlue mr-4">
                    Home
                </a>
                <a href="#recentWork" v-bind:class="{hidden: isHidden}" class="font-Nunito font-bold text block mt-4 lg:inline-block lg:mt-0 text-blue hover:text-darkBlue mr-4">
                    Our Work
                </a>
                <a href="#contact" v-bind:class="{hidden: isHidden}" class="font-Nunito font-bold text block mt-4 lg:inline-block lg:mt-0 text-blue hover:text-darkBlue mr-4">
                    Contact Us
                </a>

            </div>
        </div>
    </nav>
</template>
<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'

@Component
export default class Navbar extends Vue {
   
    isHidden: boolean = true;

    toggleNav(): void {
        this.isHidden = !this.isHidden;
    }


}
</script>
<style>
</style>
